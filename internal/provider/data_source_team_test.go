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
					resource.TestCheckResourceAttr("data.mattermost_team.test", "description", "bar description"),
				),
			},
		},
	})
}

const testAccDataSourceTeam = `
resource "mattermost_team" "foo" {
	description = "bar description"
	display_name = "bar display"
	name = "bar"
}

data "mattermost_team" "test" {
  name = mattermost_team.foo.name
}
`
