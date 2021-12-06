package provider

import (
	"context"

	"github.com/mattermost/mattermost-server/v6/model"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceTeam() *schema.Resource {
	return &schema.Resource{
		// This description is used by the documentation generator and the language server.
		Description: "Sample resource in the Terraform provider Team.",

		CreateContext: resourceTeamCreate,
		ReadContext:   resourceTeamRead,
		UpdateContext: resourceTeamUpdate,
		DeleteContext: resourceTeamDelete,
		Importer:      &schema.ResourceImporter{StateContext: schema.ImportStatePassthroughContext},

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
			"type": {
				Description:  "Type of the team.",
				Type:         schema.TypeString,
				Optional:     true,
				Default:      model.TeamOpen,
				ValidateFunc: validation.StringInSlice([]string{model.TeamOpen, model.TeamInvite}, false),
			},
			"email": {
				Description: "Administrator Email (anyone with this email is automatically a team admin).",
				Type:        schema.TypeString,
				Computed:    true,
			},
		},
	}
}

func resourceTeamCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*model.Client4)

	team, resp, err := c.CreateTeam(&model.Team{
		DisplayName: d.Get("display_name").(string),
		Name:        d.Get("name").(string),
		Description: d.Get("description").(string),
		Type:        d.Get("type").(string),
	})
	if err != nil {
		return diag.Errorf("cannot create team: %v", err)
	}

	if resp.StatusCode != 201 {
		return diag.Errorf("invalid status returned %d", resp.StatusCode)
	}

	d.SetId(team.Id)

	return resourceTeamRead(ctx, d, meta)
}

func resourceTeamRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*model.Client4)
	id := d.Id()

	team, resp, err := c.GetTeam(id, "")
	if err != nil {
		return diag.Errorf("cannot get team by name: %v", err)
	}

	if resp.StatusCode == 404 {
		d.SetId("")
		return nil
	}

	if resp.StatusCode != 200 {
		return diag.Errorf("invalid status returned %d", resp.StatusCode)
	}

	d.SetId(team.Id)
	d.Set("name", team.Name)
	d.Set("display_name", team.DisplayName)
	d.Set("description", team.Description)
	d.Set("email", team.Email)
	d.Set("type", team.Type)

	return nil
}

func resourceTeamUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*model.Client4)

	_, resp, err := c.UpdateTeam(&model.Team{
		Id:          d.Id(),
		DisplayName: d.Get("display_name").(string),
		Name:        d.Get("name").(string),
		Description: d.Get("description").(string),
		Type:        d.Get("type").(string),
	})
	if err != nil {
		return diag.Errorf("cannot update team: %v", err)
	}

	if resp.StatusCode != 200 {
		return diag.Errorf("invalid status returned %d", resp.StatusCode)
	}

	return resourceTeamRead(ctx, d, meta)
}

func resourceTeamDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*model.Client4)

	resp, err := c.PermanentDeleteTeam(d.Id())
	if err != nil {
		return diag.Errorf("cannot delete team: %v", err)
	}

	if resp.StatusCode != 200 {
		return diag.Errorf("invalid status returned %d", resp.StatusCode)
	}

	return nil
}
