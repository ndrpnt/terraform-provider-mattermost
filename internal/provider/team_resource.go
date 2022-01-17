package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/mattermost/mattermost-server/v6/model"
)

// FIXME quid du nomage de cette resource ?
type teamResourceType struct{}

func (t teamResourceType) GetSchema(ctx context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "Sample resource in the Terraform provider Team.",

		Attributes: map[string]tfsdk.Attribute{
			"name": {
				MarkdownDescription: "Name of the team.",
				Type:                types.StringType,
				Required:            true, // FIXME Optional: false c'est pareil ?
			},
			"id": { // FIXME est ce que c'est l'ID généré de la resource ?
				Computed:            true,
				MarkdownDescription: "Example identifier",
				PlanModifiers: tfsdk.AttributePlanModifiers{
					tfsdk.UseStateForUnknown(),
				},
				Type: types.StringType,
			},
			"description": {
				MarkdownDescription: "Description of the team.",
				Type:                types.StringType,
				Optional:            true,
			},
			"display_name": {
				MarkdownDescription: "Display name of the team.",
				Type:                types.StringType,
				Required:            true,
			},
			"type": {
				MarkdownDescription: "Type of the team.",
				Type:                types.StringType,
				Optional:            true,
				// Default:      model.TeamOpen, FIXME ça existe encore ?
				Validators: validation.StringInSlice([]string{model.TeamOpen, model.TeamInvite}, false),
			},
			"email": {
				MarkdownDescription: "Administrator Email (anyone with this email is automatically a team admin).",
				Type:                types.StringType,
				Computed:            true,
				// TODO add email format validator ?
			},
		},
	}, nil
}
