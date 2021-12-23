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
			"team_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"header": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"purpose": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"creator_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceChannelRead(_ context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*model.Client4)
	name := d.Get("name").(string)
	teamId := d.Get("team_id").(string)

	channel, resp, err := c.GetChannelByName(name, teamId, "")
	if err != nil {
		return diag.Errorf("cannot get channel by name: %v", fmtErr(resp, err))
	}

	d.SetId(channel.Id)
	d.Set("type", channel.Type)
	d.Set("display_name", channel.DisplayName)
	d.Set("header", channel.Header)
	d.Set("purpose", channel.Purpose)
	d.Set("creator_id", channel.CreatorId)

	return nil
}
