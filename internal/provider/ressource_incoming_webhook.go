package provider

import (
	"context"

	"github.com/mattermost/mattermost-server/v6/model"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ressourceIncomingWebhook() *schema.Resource {
	return &schema.Resource{
		// This description is used by the documentation generator and the language server.
		Description: "Sample resource in the Terraform provider IncomingWebhook.",

		CreateContext: ressourceIncomingWebhookCreate,
		ReadContext:   ressourceIncomingWebhookRead,
		UpdateContext: ressourceIncomingWebhookUpdate,
		DeleteContext: ressourceIncomingWebhookDelete,
		Importer:      &schema.ResourceImporter{StateContext: schema.ImportStatePassthroughContext},

		Schema: map[string]*schema.Schema{
			"description": {
				Description: "Description of the webhook.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"name": {
				Description: "Name of the webhook",
				Type:        schema.TypeString,
				Required:    true,
			},
			"channel_id": {
				Description: "Id of the channel that receives the webhook payloads",
				Type:        schema.TypeString,
				Required:    true,
			},
			"username": {
				Description: "The display name for this incoming webhook",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"icon_url": {
				Description: "The profile picture this incoming webhook will use when posting",
				Type:        schema.TypeString,
				Optional:    true,
			},
		},
	}
}

func ressourceIncomingWebhookCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*model.Client4)

	webhook, resp, err := c.CreateIncomingWebhook(&model.IncomingWebhook{
		ChannelId:   d.Get("channel_id").(string),
		DisplayName: d.Get("name").(string),
		Description: d.Get("description").(string),
		Username:    d.Get("username").(string),
		IconURL:     d.Get("icon_url").(string),
	})
	if err != nil {
		return diag.Errorf("cannot create webhook: %v", err)
	}

	if resp.StatusCode != 201 {
		return diag.Errorf("invalid status returned %d", resp.StatusCode)
	}

	d.SetId(webhook.Id)

	return ressourceIncomingWebhookRead(ctx, d, meta)
}

func ressourceIncomingWebhookRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*model.Client4)
	id := d.Id()

	webhook, resp, err := c.GetIncomingWebhook(id, "")
	if err != nil {
		return diag.Errorf("cannot get webhook by id: %v", err)
	}

	if resp.StatusCode == 404 {
		d.SetId("")
		return nil
	}

	if resp.StatusCode != 200 {
		return diag.Errorf("invalid status returned %d", resp.StatusCode)
	}

	d.SetId(webhook.Id)
	d.Set("name", webhook.DisplayName)
	d.Set("channel_id", webhook.ChannelId)
	d.Set("description", webhook.Description)
	d.Set("username", webhook.Username)
	d.Set("icon_url", webhook.IconURL)

	return nil
}

func ressourceIncomingWebhookUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*model.Client4)

	_, resp, err := c.UpdateIncomingWebhook(&model.IncomingWebhook{
		Id:          d.Id(),
		DisplayName: d.Get("name").(string),
		ChannelId:   d.Get("channel_id").(string),
		Description: d.Get("description").(string),
		Username:    d.Get("username").(string),
		IconURL:     d.Get("icon_url").(string),
	})
	if err != nil {
		return diag.Errorf("cannot update webhook: %v", err)
	}

	if resp.StatusCode != 200 {
		return diag.Errorf("invalid status returned %d", resp.StatusCode)
	}

	return ressourceIncomingWebhookRead(ctx, d, meta)
}

func ressourceIncomingWebhookDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*model.Client4)

	resp, err := c.DeleteIncomingWebhook(d.Id())
	if err != nil {
		return diag.Errorf("cannot delete webhook: %v", err)
	}

	if resp.StatusCode != 200 {
		return diag.Errorf("invalid status returned %d", resp.StatusCode)
	}

	return nil
}
