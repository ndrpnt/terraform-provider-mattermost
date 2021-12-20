package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/mattermost/mattermost-server/v6/model"
)

func resourcePost() *schema.Resource {
	return &schema.Resource{
		Description: "Manage a Mattermost post.",

		CreateContext: resourcePostCreate,
		ReadContext:   resourcePostRead,
		UpdateContext: resourcePostUpdate,
		DeleteContext: resourcePostDelete,
		Importer:      &schema.ResourceImporter{StateContext: schema.ImportStatePassthroughContext},

		Schema: map[string]*schema.Schema{
			"message": {
				Description: "Content of the message.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"channel_id": {
				Description: "Id of the channel to send the message in.",
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
			},
		},
	}
}

func resourcePostCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*model.Client4)

	post, resp, err := c.CreatePost(&model.Post{
		ChannelId: d.Get("channel_id").(string),
		Message:   d.Get("message").(string),
	})
	if err != nil {
		return diag.Errorf("cannot create post: %v", err)
	}

	if resp.StatusCode != 201 {
		return diag.Errorf("invalid status returned %d", resp.StatusCode)
	}

	d.SetId(post.Id)

	return resourcePostRead(ctx, d, meta)
}

func resourcePostRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*model.Client4)

	post, resp, err := c.GetPost(d.Id(), "")
	if err != nil {
		return diag.Errorf("cannot get post by ID: %v", err)
	}

	if resp.StatusCode == 404 {
		d.SetId("")
		return nil
	}

	if resp.StatusCode != 200 {
		return diag.Errorf("invalid status returned %d", resp.StatusCode)
	}

	d.Set("message", post.Message)
	d.Set("channel_id", post.ChannelId)

	return nil
}

func resourcePostUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*model.Client4)

	_, resp, err := c.UpdatePost(d.Id(), &model.Post{
		Id:        d.Id(),
		Message:   d.Get("message").(string),
		ChannelId: d.Get("channel_id").(string),
	})
	if err != nil {
		return diag.Errorf("cannot update post: %v", err)
	}

	if resp.StatusCode != 200 {
		return diag.Errorf("invalid status returned %d", resp.StatusCode)
	}

	return resourcePostRead(ctx, d, meta)
}

func resourcePostDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*model.Client4)

	resp, err := c.DeletePost(d.Id())
	if err != nil {
		return diag.Errorf("cannot delete post: %v", err)
	}

	if resp.StatusCode != 200 {
		return diag.Errorf("invalid status returned %d", resp.StatusCode)
	}

	return nil
}
