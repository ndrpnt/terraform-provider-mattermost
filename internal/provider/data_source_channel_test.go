package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceChannel(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceChannel,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.mattermost_channel.sut", "display_name", "Example display name"),
					resource.TestCheckResourceAttr("data.mattermost_channel.sut", "header", "Example header"),
				),
			},
		},
	})
}

const testAccDataSourceChannel = `
resource "mattermost_team" "example_team" {
	name = "examplename"
	display_name = "Example display name"
	description = "Example description"
}

resource "mattermost_channel" "example_channel" {
	name = "examplename"
	display_name = "Example display name"
	header = "Example header"
	team_id = mattermost_team.example_team.id
}

data "mattermost_channel" "sut" {
  name = mattermost_channel.example_channel.name
  team_id = mattermost_team.example_team.id
}
`
