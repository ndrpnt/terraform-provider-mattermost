package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceScaffolding(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceScaffolding,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.mattermost_team.test", "description", "pas bon"),
					resource.TestCheckResourceAttr("data.mattermost_team.test", "id", "549x3sak1jd49qyde8xihso3fe"),
				),
			},
		},
	})
}

const testAccDataSourceScaffolding = `
data "mattermost_team" "test" {
  name = "test"
}
`
