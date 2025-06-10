package provider

import (
	"context"

	"github.com/mattermost/mattermost-server/v6/model"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSlashCommand() *schema.Resource {
	return &schema.Resource{
		Description: "Manages a Mattermost slash command.",

		CreateContext: resourceSlashCommandCreate,
		ReadContext:   resourceSlashCommandRead,
		UpdateContext: resourceSlashCommandUpdate,
		DeleteContext: resourceSlashCommandDelete,
		Importer:      &schema.ResourceImporter{StateContext: schema.ImportStatePassthroughContext},

		Schema: map[string]*schema.Schema{
			"trigger": {
				Description: "The trigger word for the slash command (without the leading slash).",
				Type:        schema.TypeString,
				Required:    true,
			},
			"display_name": {
				Description: "Display name for the slash command.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"description": {
				Description: "Description of the slash command.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"url": {
				Description: "The URL that the command will make a request to.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"method": {
				Description:  "The HTTP method for the request (P for POST, G for GET).",
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "P",
				ValidateFunc: validation.StringInSlice([]string{"P", "G"}, false),
			},
			"username": {
				Description: "The username that responses will post as.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"icon_url": {
				Description: "The profile picture that responses will use.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"team_id": {
				Description: "The team ID where this command will be available.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"auto_complete": {
				Description: "Whether the command shows up in autocomplete.",
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
			},
			"auto_complete_desc": {
				Description: "Description for autocomplete.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"auto_complete_hint": {
				Description: "Hint for autocomplete arguments.",
				Type:        schema.TypeString,
				Optional:    true,
			},
		},
	}
}

func resourceSlashCommandCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*model.Client4)

	command, resp, err := c.CreateCommand(&model.Command{
		Trigger:          d.Get("trigger").(string),
		DisplayName:      d.Get("display_name").(string),
		Description:      d.Get("description").(string),
		URL:              d.Get("url").(string),
		Method:           d.Get("method").(string),
		Username:         d.Get("username").(string),
		IconURL:          d.Get("icon_url").(string),
		TeamId:           d.Get("team_id").(string),
		AutoComplete:     d.Get("auto_complete").(bool),
		AutoCompleteDesc: d.Get("auto_complete_desc").(string),
		AutoCompleteHint: d.Get("auto_complete_hint").(string),
	})
	if err != nil {
		return diag.Errorf("cannot create slash command: %v", err)
	}

	if resp.StatusCode != 201 {
		return diag.Errorf("invalid status returned %d", resp.StatusCode)
	}

	d.SetId(command.Id)

	return resourceSlashCommandRead(ctx, d, meta)
}

func resourceSlashCommandRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*model.Client4)
	id := d.Id()

	command, resp, err := c.GetCommandById(id)
	if err != nil {
		return diag.Errorf("cannot get slash command by id: %v", err)
	}

	if resp.StatusCode == 404 {
		d.SetId("")
		return nil
	}

	if resp.StatusCode != 200 {
		return diag.Errorf("invalid status returned %d", resp.StatusCode)
	}

	d.SetId(command.Id)
	d.Set("trigger", command.Trigger)
	d.Set("display_name", command.DisplayName)
	d.Set("description", command.Description)
	d.Set("url", command.URL)
	d.Set("method", command.Method)
	d.Set("username", command.Username)
	d.Set("icon_url", command.IconURL)
	d.Set("team_id", command.TeamId)
	d.Set("auto_complete", command.AutoComplete)
	d.Set("auto_complete_desc", command.AutoCompleteDesc)
	d.Set("auto_complete_hint", command.AutoCompleteHint)

	return nil
}

func resourceSlashCommandUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*model.Client4)

	_, resp, err := c.UpdateCommand(&model.Command{
		Id:               d.Id(),
		Trigger:          d.Get("trigger").(string),
		DisplayName:      d.Get("display_name").(string),
		Description:      d.Get("description").(string),
		URL:              d.Get("url").(string),
		Method:           d.Get("method").(string),
		Username:         d.Get("username").(string),
		IconURL:          d.Get("icon_url").(string),
		TeamId:           d.Get("team_id").(string),
		AutoComplete:     d.Get("auto_complete").(bool),
		AutoCompleteDesc: d.Get("auto_complete_desc").(string),
		AutoCompleteHint: d.Get("auto_complete_hint").(string),
	})
	if err != nil {
		return diag.Errorf("cannot update slash command: %v", err)
	}

	if resp.StatusCode != 200 {
		return diag.Errorf("invalid status returned %d", resp.StatusCode)
	}

	return resourceSlashCommandRead(ctx, d, meta)
}

func resourceSlashCommandDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*model.Client4)

	resp, err := c.DeleteCommand(d.Id())
	if err != nil {
		return diag.Errorf("cannot delete slash command: %v", err)
	}

	if resp.StatusCode != 200 {
		return diag.Errorf("invalid status returned %d", resp.StatusCode)
	}

	return nil
}
