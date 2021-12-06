package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/mattermost/mattermost-server/v6/model"
)

func dataSourceTeam() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceTeamRead,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"description": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceTeamRead(_ context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*model.Client4)
	name := d.Get("name").(string)

	team, _, err := c.GetTeamByName(name, "")
	if err != nil {
		return diag.Errorf("cannot get team by name: %v", err)
	}

	if team == nil {
		return diag.Errorf("team %s not found", name)
	}

	d.SetId(team.Id)
	d.Set("description", team.Description)

	return nil
}
