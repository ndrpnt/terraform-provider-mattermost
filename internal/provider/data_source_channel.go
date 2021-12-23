package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/mattermost/mattermost-server/v6/model"
)

func dataSourceChannel() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceChannelRead,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"header": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"team_id": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func dataSourceChannelRead(_ context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*model.Client4)
	name := d.Get("name").(string)
	teamId := d.Get("team_id").(string)

	channel, _, err := c.GetChannelByName(name, teamId, "")
	if err != nil {
		return diag.Errorf("cannot get channel: %v", err)
	}

	if channel == nil {
		return diag.Errorf("channel %s not found", name)
	}

	d.SetId(channel.Id)
	d.Set("display_name", channel.DisplayName)
	d.Set("header", channel.Header)

	return nil
}
