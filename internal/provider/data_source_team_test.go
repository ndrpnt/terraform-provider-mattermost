package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceTeam(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceTeam,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.mattermost_team.test", "description", "pas bon"),
					resource.TestCheckResourceAttr("data.mattermost_team.test", "id", "549x3sak1jd49qyde8xihso3fe"),
				),
			},
		},
	})
}

const testAccDataSourceTeam = `
data "mattermost_team" "test" {
  name = "test"
}
`
