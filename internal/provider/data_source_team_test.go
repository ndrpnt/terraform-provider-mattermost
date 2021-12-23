package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceTeam(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceTeam,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.mattermost_team.sut", "display_name", "Example display name"),
					resource.TestCheckResourceAttr("data.mattermost_team.sut", "description", "Example description"),
				),
			},
		},
	})
}

const testAccDataSourceTeam = `
resource "mattermost_team" "example_team" {
	name = "examplename"
	display_name = "Example display name"
	description = "Example description"
}

data "mattermost_team" "sut" {
  name = mattermost_team.example_team.name
}
`
