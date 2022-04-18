package provider

import (
	"context"
	"fmt"
	"runtime"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/mattermost/mattermost-server/v6/model"
)

type provider struct {
	client     *model.Client4
	configured bool
	version    string
}

type providerData struct {
	Url      types.String `tfsdk:"url"`
	Token    types.String `tfsdk:"token"`
	LoginId  types.String `tfsdk:"login_id"`
	Password types.String `tfsdk:"password"`
}

func (p *provider) ValidateConfig(ctx context.Context, req tfsdk.ValidateProviderConfigRequest, resp *tfsdk.ValidateProviderConfigResponse) {
	var data providerData
	diags := req.Config.Get(ctx, &data)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	badCredentials := "Either token or login_id/password must be specified"

	if data.LoginId.Null && isSet(data.Password) {
		resp.Diagnostics.AddError(badCredentials, "Password is set but login is null")
	}

	if data.Password.Null && isSet(data.LoginId) {
		resp.Diagnostics.AddError(badCredentials, "Login is set but password is null")
	}

	if data.Token.Null && data.LoginId.Null {
		resp.Diagnostics.AddError(badCredentials, "Both token and login are null")
	}

	if data.Token.Null && data.Password.Null {
		resp.Diagnostics.AddError(badCredentials, "Both token and password are null")
	}

	if isSet(data.Token) && isSet(data.LoginId) {
		resp.Diagnostics.AddError(badCredentials, "Both token and login are set")
	}

	if isSet(data.Token) && isSet(data.Password) {
		resp.Diagnostics.AddError(badCredentials, "Both token and password are set")
	}
}

func (p *provider) Configure(ctx context.Context, req tfsdk.ConfigureProviderRequest, resp *tfsdk.ConfigureProviderResponse) {
	var data providerData
	diags := req.Config.Get(ctx, &data)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Check for unknown and env vars …
	// See https://learn.hashicorp.com/tutorials/terraform/plugin-framework-create
	c := model.NewAPIv4Client(data.Url.Value)
	userAgent := fmt.Sprintf("terraform-provider-mattermost/%s (%s)", p.version, runtime.GOOS)
	c.HTTPHeader = map[string]string{"User-Agent": userAgent}

	if !data.Token.Null {
		c.SetOAuthToken(data.Token.Value)
	} else {
		_, _, err := c.Login(data.LoginId.Value, data.Password.Value)
		if err != nil {
			resp.Diagnostics.AddError("Cannot login with given login_id and password", err.Error())
			return
		}
	}

	p.configured = true
}

func (p *provider) GetResources(context.Context) (map[string]tfsdk.ResourceType, diag.Diagnostics) {
	return map[string]tfsdk.ResourceType{
		// "mattermost_team":           teamResource{},
		// "mattermost_channel":        channelResource{},
		// "mattermost_channel_member": channelMemberResource{},
		// "mattermost_team_member":    teamMemberResource{},
		// "mattermost_post":           postResource{},
		// "mattermost_user":           userResource{},
	}, nil
}

func (p *provider) GetDataSources(context.Context) (map[string]tfsdk.DataSourceType, diag.Diagnostics) {
	return map[string]tfsdk.DataSourceType{
		// "mattermost_team":    teamDataSource{},
		// "mattermost_channel": channelDataSource{},
		// "mattermost_user":    userDataSource{},
	}, nil
}

func (p *provider) GetSchema(context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Attributes: map[string]tfsdk.Attribute{
			"url": {
				MarkdownDescription: "Can also be provided via the MM_URL environment variable",
				Optional:            true,
				Type:                types.StringType,
			},
			"token": {
				MarkdownDescription: "Can also be provided via the MM_TOKEN environment variable",
				Optional:            true,
				Type:                types.StringType,
			},
			"login_id": {
				MarkdownDescription: "Can also be provided via the MM_LOGIN_ID environment variable",
				Optional:            true,
				Type:                types.StringType,
			},
			"password": {
				MarkdownDescription: "Can also be provided via the MM_PASSWORD environment variable",
				Optional:            true,
				Type:                types.StringType,
			},
		},
	}, nil
}

func New(version string) func() tfsdk.Provider {
	return func() tfsdk.Provider {
		return &provider{
			version: version,
		}
	}
}

func convertProviderType(in tfsdk.Provider) (provider, diag.Diagnostics) {
	var diags diag.Diagnostics

	p, ok := in.(*provider)
	if !ok {
		diags.AddError(
			"Unexpected Provider Instance Type",
			fmt.Sprintf("While creating the data source or resource, an unexpected provider type (%T) was received. This is always a bug in the provider code and should be reported to the provider developers.", p),
		)
		return provider{}, diags
	}

	if p == nil {
		diags.AddError(
			"Unexpected Provider Instance Type",
			"While creating the data source or resource, an unexpected empty provider instance was received. This is always a bug in the provider code and should be reported to the provider developers.",
		)
		return provider{}, diags
	}

	return *p, diags
}
