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
					resource.TestCheckResourceAttr(
						"mattermost_team.foo", "name", "bar"),
				),
			},
		},
	})
}

const testAccResourceTeam = `
resource "mattermost_team" "foo" {
	description = "bar description"
	display_name = "bar display"
	email = "foo@bar.xyz"
  	name = "bar"
}
`
