package provider

import (
	"context"
	"fmt"
	"runtime"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/mattermost/mattermost-server/v6/model"
)

var version = "unknown"

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
					Optional:    true,
					DefaultFunc: schema.EnvDefaultFunc("MM_URL", ""),
				},
				"token": {
					Type:        schema.TypeString,
					Optional:    true,
					DefaultFunc: schema.EnvDefaultFunc("MM_TOKEN", ""),
				},
			},
			DataSourcesMap: map[string]*schema.Resource{
				"mattermost_team":    dataSourceTeam(),
				"mattermost_channel": dataSourceChannel(),
				"mattermost_user":    dataSourceUser(),
			},
			ResourcesMap: map[string]*schema.Resource{
				"mattermost_team":           resourceTeam(),
				"mattermost_channel":        resourceChannel(),
				"mattermost_channel_member": resourceChannelMember(),
				"mattermost_team_member":    resourceTeamMember(),
			},
		}

		p.ConfigureContextFunc = configure(version, p)

		return p
	}
}

func configure(version string, p *schema.Provider) func(context.Context, *schema.ResourceData) (interface{}, diag.Diagnostics) {
	return func(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
		url := d.Get("url").(string)
		token := d.Get("token").(string)
		c := model.NewAPIv4Client(url)
		c.SetOAuthToken(token)
		userAgent := fmt.Sprintf("terraform-provider-mattermost/%s (%s)", version, runtime.GOOS)
		c.HTTPHeader = map[string]string{"User-Agent": userAgent}

		return c, nil
	}
}
