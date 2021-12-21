package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceChannel(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceChannel,
				Check:  resource.ComposeTestCheckFunc(),
			},
		},
	})
}

const testAccDataSourceChannel = `
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

data "mattermost_channel" "test" {
  name = mattermost_channel.foo.name
  team_id = mattermost_team.test.id
}
`
