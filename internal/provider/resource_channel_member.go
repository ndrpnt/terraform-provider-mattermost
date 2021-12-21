package provider

import (
	"context"
	"strings"

	"github.com/mattermost/mattermost-server/v6/model"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceChannelMember() *schema.Resource {
	return &schema.Resource{
		// This description is used by the documentation generator and the language server.
		Description: "Sample resource in the Terraform provider ChannelMember.",

		CreateContext: resourceChannelMemberCreate,
		ReadContext:   resourceChannelMemberRead,
		DeleteContext: resourceChannelMemberDelete,
		Importer:      &schema.ResourceImporter{StateContext: schema.ImportStatePassthroughContext},

		Schema: map[string]*schema.Schema{
			"channel_id": {
				Description: "The id of the channel",
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

func resourceChannelMemberCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*model.Client4)
	channelId := d.Get("channel_id").(string)
	userId := d.Get("user_id").(string)

	_, resp, err := c.AddChannelMember(channelId, userId)
	if err != nil {
		return diag.Errorf("cannot create channel_member: %v", err)
	}

	if resp.StatusCode != 201 {
		return diag.Errorf("invalid status returned %d", resp.StatusCode)
	}

	d.SetId(channelId + "/" + userId)

	return resourceChannelMemberRead(ctx, d, meta)
}

func resourceChannelMemberRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*model.Client4)
	parts := strings.Split(d.Id(), "/")
	channelId := parts[0]
	userId := parts[1]

	_, resp, err := c.GetChannelMember(channelId, userId, "")
	if err != nil {
		return diag.Errorf("cannot get channel_member: %v", err)
	}

	if resp.StatusCode == 404 {
		d.SetId("")
		return nil
	}

	if resp.StatusCode != 200 {
		return diag.Errorf("invalid status returned %d", resp.StatusCode)
	}

	d.Set("channel_id", channelId)
	d.Set("user_id", userId)

	return nil
}

func resourceChannelMemberDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*model.Client4)
	channelId := d.Get("channel_id").(string)
	userId := d.Get("user_id").(string)

	resp, err := c.RemoveUserFromChannel(channelId, userId)
	if err != nil {
		return diag.Errorf("cannot delete channel_member: %v", err)
	}

	if resp.StatusCode != 200 {
		return diag.Errorf("invalid status returned %d", resp.StatusCode)
	}

	return nil
}
