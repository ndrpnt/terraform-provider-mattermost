package provider

import (
	"context"
	"strings"

	"github.com/mattermost/mattermost-server/v6/model"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceTeamMember() *schema.Resource {
	return &schema.Resource{
		// This description is used by the documentation generator and the language server.
		Description: "Sample resource in the Terraform provider TeamMember.",

		CreateContext: resourceTeamMemberCreate,
		ReadContext:   resourceTeamMemberRead,
		DeleteContext: resourceTeamMemberDelete,
		Importer:      &schema.ResourceImporter{StateContext: schema.ImportStatePassthroughContext},

		Schema: map[string]*schema.Schema{
			"team_id": {
				Description: "The id of the team",
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
			},
			"user_id": {
				Description: "The id of the user to add",
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
			},
		},
	}
}

func resourceTeamMemberCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*model.Client4)
	teamId := d.Get("team_id").(string)
	userId := d.Get("user_id").(string)

	_, resp, err := c.AddTeamMember(teamId, userId)
	if err != nil {
		return diag.Errorf("cannot create team_member: %v", err)
	}

	if resp.StatusCode != 201 {
		return diag.Errorf("invalid status returned %d", resp.StatusCode)
	}

	d.SetId(teamId + "/" + userId)

	return resourceTeamMemberRead(ctx, d, meta)
}

func resourceTeamMemberRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*model.Client4)

	parts := strings.Split(d.Id(), "/")
	teamId := parts[0]
	userId := parts[1]
	_, resp, err := c.GetTeamMember(teamId, userId, "")
	if err != nil {
		return diag.Errorf("cannot get team_member: %v", err)
	}

	if resp.StatusCode == 404 {
		d.SetId("")
		return nil
	}

	if resp.StatusCode != 200 {
		return diag.Errorf("invalid status returned %d", resp.StatusCode)
	}

	d.Set("team_id", teamId)
	d.Set("user_id", userId)

	return nil
}

func resourceTeamMemberDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*model.Client4)
	teamId := d.Get("team_id").(string)
	userId := d.Get("user_id").(string)

	resp, err := c.RemoveTeamMember(teamId, userId)
	if err != nil {
		return diag.Errorf("cannot delete team_member: %v", err)
	}

	if resp.StatusCode != 200 {
		return diag.Errorf("invalid status returned %d", resp.StatusCode)
	}

	return nil
}
