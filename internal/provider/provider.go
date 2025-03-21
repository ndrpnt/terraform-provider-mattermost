package provider

import (
	"context"
	"fmt"
	"runtime"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/mattermost/mattermost-server/v6/model"
)

func init() {
	// Set descriptions to support markdown syntax, this will be used in document generation
	// and the language server.
	schema.DescriptionKind = schema.StringMarkdown

	// Customize the content of descriptions when output. For example you can add defaults on
	// to the exported descriptions if present.
	// schema.SchemaDescriptionBuilder = func(s *schema.Schema) string {
	// 	desc := s.Description
	// 	if s.Default != nil {
	// 		desc += fmt.Sprintf(" Defaults to `%v`.", s.Default)
	// 	}
	// 	return strings.TrimSpace(desc)
	// }
}

func New(version string) func() *schema.Provider {
	return func() *schema.Provider {
		p := &schema.Provider{
			Schema: map[string]*schema.Schema{
				"url": {
					Type:        schema.TypeString,
					Required:    true,
					DefaultFunc: schema.EnvDefaultFunc("MM_URL", nil),
					Description: "Can also be provided via the MM_URL environment variable",
				},
				"token": {
					Type:         schema.TypeString,
					Optional:     true,
					DefaultFunc:  schema.EnvDefaultFunc("MM_TOKEN", nil),
					ExactlyOneOf: []string{"token", "login_id"},
					Description:  "Can also be provided via the MM_TOKEN environment variable",
				},
				"login_id": {
					Type:         schema.TypeString,
					Optional:     true,
					DefaultFunc:  schema.EnvDefaultFunc("MM_LOGIN_ID", nil),
					ExactlyOneOf: []string{"token", "login_id"},
					Description:  "Can also be provided via the MM_LOGIN_ID environment variable",
				},
				"password": {
					Type:         schema.TypeString,
					Optional:     true,
					DefaultFunc:  schema.EnvDefaultFunc("MM_PASSWORD", nil),
					RequiredWith: []string{"login_id"},
					Description:  "Can also be provided via the MM_PASSWORD environment variable",
				},
			},
			DataSourcesMap: map[string]*schema.Resource{
				"mattermost_team":    dataSourceTeam(),
				"mattermost_channel": dataSourceChannel(),
				"mattermost_user":    dataSourceUser(),
			},
			ResourcesMap: map[string]*schema.Resource{
				"mattermost_team":             resourceTeam(),
				"mattermost_channel":          resourceChannel(),
				"mattermost_channel_member":   resourceChannelMember(),
				"mattermost_team_member":      resourceTeamMember(),
				"mattermost_post":             resourcePost(),
				"mattermost_user":             resourceUser(),
				"mattermost_incoming_webhook": ressourceIncomingWebhook(),
			},
		}

		p.ConfigureContextFunc = configure(version, p)

		return p
	}
}

func configure(version string, p *schema.Provider) func(context.Context, *schema.ResourceData) (interface{}, diag.Diagnostics) {
	return func(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
		url := d.Get("url").(string)

		c := model.NewAPIv4Client(url)
		userAgent := fmt.Sprintf("terraform-provider-mattermost/%s (%s)", version, runtime.GOOS)
		c.HTTPHeader = map[string]string{"User-Agent": userAgent}

		token, ok := d.GetOk("token")
		if ok {
			c.SetOAuthToken(token.(string))
		} else {
			loginId := d.Get("login_id").(string)
			password := d.Get("password").(string)
			_, _, err := c.Login(loginId, password)
			if err != nil {
				return nil, diag.Errorf("cannot login with given login_id and password: %v", err)
			}
		}

		return c, nil
	}
}
