package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccResourceIncommingWebhook(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccResourceIncomingWebhook,
				Check:  resource.ComposeAggregateTestCheckFunc(),
			},
		},
	})
}

const testAccResourceIncomingWebhook = `
resource "mattermost_team" "test" {
  name         = "sheh"
  display_name = "An sheh bis"
}
resource "mattermost_channel" "foo" {
  description  = "foo description"
  display_name = "foo display"
  name         = "foo"
  team_id      = mattermost_team.test.id
}
resource "mattermost_incoming_webhook" "test" {
  name        = "test_webhook"
  description = "A test incoming webhook"
  channel_id  = mattermost_channel.foo.id
}
`
