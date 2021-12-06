package provider

import (
	"context"
	"github.com/mattermost/mattermost-server/v6/model"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceTeam() *schema.Resource {
	return &schema.Resource{
		// This description is used by the documentation generator and the language server.
		Description: "Sample resource in the Terraform provider Team.",

		CreateContext: resourceTeamCreate,
		ReadContext:   resourceTeamRead,
		UpdateContext: resourceTeamUpdate,
		DeleteContext: resourceTeamDelete,

		Schema: map[string]*schema.Schema{
			"name": {
				Description: "Name of the team.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"description": {
				Description: "Description of the team.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"display_name": {
				Description: "Display name of the team.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"email": {
				Description: "Administrator Email (anyone with this email is automatically a team admin).",
				Type:        schema.TypeString,
				Optional:    true,
			},
		},
	}
}

func resourceTeamCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*model.Client4)
	id := d.Id()

	team, _, err := c.CreateTeam(&model.Team{
		DisplayName: d.Get("display_name").(string),
		Name:        d.Get("name").(string),
		Description: d.Get("description").(string),
		Email:       d.Get("email").(string),
	})
	if err != nil {
		return diag.Errorf("cannot create team: %v", err)
	}

	if team == nil {
		return diag.Errorf("team with Id: %q not found", id)
	}

	d.SetId(team.Id)
	d.Set("name", team.Name)
	d.Set("display_name", team.DisplayName)
	d.Set("description", team.Description)
	d.Set("email", team.Email)

	return resourceTeamRead(ctx, d, meta)
}

func resourceTeamRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*model.Client4)
	id := d.Id()

	team, _, err := c.GetTeam(id, "")
	if err != nil {
		return diag.Errorf("cannot get team by name: %v", err)
	}

	if team == nil {
		return diag.Errorf("team with Id: %q not found", id)
	}

	d.SetId(team.Id)
	d.Set("name", team.Name)
	d.Set("display_name", team.DisplayName)
	d.Set("description", team.Description)
	d.Set("email", team.Email)

	return nil
}

func resourceTeamUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// use the meta value to retrieve your client from the provider configure method
	// client := meta.(*apiClient)

	return diag.Errorf("not implemented")
}

func resourceTeamDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// use the meta value to retrieve your client from the provider configure method
	// client := meta.(*apiClient)

	return diag.Errorf("not implemented")
}
