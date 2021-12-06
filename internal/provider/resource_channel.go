package provider

import (
	"context"

	"github.com/mattermost/mattermost-server/v6/model"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceChannel() *schema.Resource {
	return &schema.Resource{
		// This description is used by the documentation generator and the language server.
		Description: "Sample resource in the Terraform provider Channel.",

		CreateContext: resourceChannelCreate,
		ReadContext:   resourceChannelRead,
		UpdateContext: resourceChannelUpdate,
		DeleteContext: resourceChannelDelete,

		Schema: map[string]*schema.Schema{
			"name": {
				Description: "Name of the channel.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"team_id": {
				Description: "Id of the team of the channel.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"description": {
				Description: "Description of the channel.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"display_name": {
				Description: "Display name of the channel.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"header": {
				Description: "Header of the channel.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"type": {
				Description:  "Type of the channel.",
				Type:         schema.TypeString,
				Optional:     true,
				Default:      model.ChannelTypeOpen,
				ValidateFunc: validation.StringInSlice([]string{string(model.ChannelTypeDirect), string(model.ChannelTypeGroup), string(model.ChannelTypeOpen), string(model.ChannelTypePrivate)}, false),
			},
		},
	}
}

func resourceChannelCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*model.Client4)

	channel, resp, err := c.CreateChannel(&model.Channel{
		DisplayName: d.Get("display_name").(string),
		Name:        d.Get("name").(string),
		Type:        model.ChannelType(d.Get("type").(string)),
		Header:      d.Get("header").(string),
		TeamId:      d.Get("team_id").(string),
	})
	if err != nil {
		return diag.Errorf("cannot create channel: %v", err)
	}

	if resp.StatusCode != 201 {
		return diag.Errorf("invalid status returned %d", resp.StatusCode)
	}

	d.SetId(channel.Id)

	return resourceChannelRead(ctx, d, meta)
}

func resourceChannelRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*model.Client4)
	id := d.Id()

	channel, _, err := c.GetChannel(id, "")
	if err != nil {
		return diag.Errorf("cannot get channel by name: %v", err)
	}

	if channel == nil {
		return diag.Errorf("channel with Id: %q not found", id)
	}

	d.SetId(channel.Id)
	d.Set("name", channel.Name)
	d.Set("display_name", channel.DisplayName)
	d.Set("type", channel.Type)
	d.Set("header", channel.Header)
	d.Set("team_id", channel.TeamId)

	return nil
}

func resourceChannelUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*model.Client4)

	_, resp, err := c.UpdateChannel(&model.Channel{
		Id:          d.Id(),
		DisplayName: d.Get("display_name").(string),
		Name:        d.Get("name").(string),
		Type:        model.ChannelType(d.Get("type").(string)),
		Header:      d.Get("header").(string),
		TeamId:      d.Get("team_id").(string),
	})
	if err != nil {
		return diag.Errorf("cannot update channel: %v", err)
	}

	if resp.StatusCode != 200 {
		return diag.Errorf("invalid status returned %d", resp.StatusCode)
	}

	return resourceChannelRead(ctx, d, meta)
}

func resourceChannelDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*model.Client4)

	resp, err := c.PermanentDeleteChannel(d.Id())
	if err != nil {
		return diag.Errorf("cannot delete channel: %v", err)
	}

	if resp.StatusCode != 200 {
		return diag.Errorf("invalid status returned %d", resp.StatusCode)
	}

	return nil
}
