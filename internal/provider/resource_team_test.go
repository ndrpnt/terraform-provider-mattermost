package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccResourceTeam(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccResourceTeam,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("mattermost_team.foo", "name", "bar"),
					resource.TestCheckResourceAttr("mattermost_team.foo", "display_name", "bar display"),
					resource.TestCheckResourceAttr("mattermost_team.foo", "description", "bar description"),
				),
			},
		},
	})
}

const testAccResourceTeam = `
resource "mattermost_team" "foo" {
	description = "bar description"
	display_name = "bar display"
  	name = "bar"
}
`
