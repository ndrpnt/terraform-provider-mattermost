package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccResourceChannel(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccResourceChannel,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("mattermost_channel.foo", "name", "bar"),
					resource.TestCheckResourceAttr("mattermost_channel.foo", "display_name", "bar display"),
					resource.TestCheckResourceAttr("mattermost_channel.foo", "description", "bar description"),
				),
			},
		},
	})
}

const testAccResourceChannel = `
resource "mattermost_team" "test" {
	name = "sheh"
	display_name = "An sheh bis"
}

resource "mattermost_channel" "foo" {
	description = "bar description"
	display_name = "bar display"
  	name = "bar"
	team_id = mattermost_team.test.id
}
`
