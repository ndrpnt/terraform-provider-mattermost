package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccResourceSlashCommand(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccResourceSlashCommand,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("mattermost_slash_command.test", "trigger", "test"),
					resource.TestCheckResourceAttr("mattermost_slash_command.test", "display_name", "Test Command"),
					resource.TestCheckResourceAttr("mattermost_slash_command.test", "description", "A test slash command"),
					resource.TestCheckResourceAttr("mattermost_slash_command.test", "url", "http://example.com/command"),
					resource.TestCheckResourceAttr("mattermost_slash_command.test", "method", "P"),
					resource.TestCheckResourceAttr("mattermost_slash_command.test", "username", "testbot"),
					resource.TestCheckResourceAttr("mattermost_slash_command.test", "auto_complete", "true"),
					resource.TestCheckResourceAttr("mattermost_slash_command.test", "auto_complete_desc", "Test command description"),
					resource.TestCheckResourceAttr("mattermost_slash_command.test", "auto_complete_hint", "[text]"),
				),
			},
			{
				Config: testAccResourceSlashCommandUpdate,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("mattermost_slash_command.test", "trigger", "test"),
					resource.TestCheckResourceAttr("mattermost_slash_command.test", "display_name", "Updated Test Command"),
					resource.TestCheckResourceAttr("mattermost_slash_command.test", "description", "An updated test slash command"),
					resource.TestCheckResourceAttr("mattermost_slash_command.test", "url", "http://example.com/updated-command"),
					resource.TestCheckResourceAttr("mattermost_slash_command.test", "method", "G"),
					resource.TestCheckResourceAttr("mattermost_slash_command.test", "username", "updatedbot"),
					resource.TestCheckResourceAttr("mattermost_slash_command.test", "auto_complete", "false"),
				),
			},
		},
	})
}

const testAccResourceSlashCommand = `
resource "mattermost_team" "test" {
  name         = "testteam"
  display_name = "Test Team"
}

resource "mattermost_slash_command" "test" {
  trigger            = "test"
  display_name       = "Test Command"
  description        = "A test slash command"
  url                = "http://example.com/command"
  method             = "P"
  username           = "testbot"
  team_id            = mattermost_team.test.id
  auto_complete      = true
  auto_complete_desc = "Test command description"
  auto_complete_hint = "[text]"
}
`

const testAccResourceSlashCommandUpdate = `
resource "mattermost_team" "test" {
  name         = "testteam"
  display_name = "Test Team"
}

resource "mattermost_slash_command" "test" {
  trigger       = "test"
  display_name  = "Updated Test Command"
  description   = "An updated test slash command"
  url           = "http://example.com/updated-command"
  method        = "G"
  username      = "updatedbot"
  team_id       = mattermost_team.test.id
  auto_complete = false
}
`
