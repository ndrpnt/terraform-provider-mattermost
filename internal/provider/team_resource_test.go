package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccTeamResource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read testing
			{
				Config: testAccTeamResourceConfig("one"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("team_resource.test", "name", "bogops"),
					resource.TestCheckResourceAttr("team_resource.test", "id", "uuid"),
					resource.TestCheckResourceAttr("team_resource.test", "description", "bogops"),
					resource.TestCheckResourceAttr("team_resource.test", "display_name", "Ze bogops team"),
					resource.TestCheckResourceAttr("team_resource.test", "type", "O"),
					resource.TestCheckResourceAttr("team_resource.test", "email", "hello@octo.com"),
				),
			},
			// ImportState testing
			{
				ResourceName:      "scaffolding_example.test",
				ImportState:       true,
				ImportStateVerify: true,
				// This is not normally necessary, but is here because this
				// example code does not have an actual upstream service.
				// Once the Read method is able to refresh information from
				// the upstream service, this can be removed.
				ImportStateVerifyIgnore: []string{"configurable_attribute"},
			},
			// Update and Read testing
			{
				Config: testAccTeamResourceConfig("two"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("scaffolding_example.test", "configurable_attribute", "two"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccTeamResourceConfig(configurableAttribute string) string {
	return fmt.Sprintf(`
resource "scaffolding_example" "test" {
  configurable_attribute = %[1]q
}
`, configurableAttribute)
}
