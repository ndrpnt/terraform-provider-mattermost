package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceChannel(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceChannelConfig(acctest.RandString(16)),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.mattermost_channel.sut", "type", "O"),
					resource.TestCheckResourceAttr("data.mattermost_channel.sut", "display_name", "Example display name"),
					resource.TestCheckResourceAttr("data.mattermost_channel.sut", "header", "Example header"),
					resource.TestCheckResourceAttr("data.mattermost_channel.sut", "purpose", ""),
					resource.TestCheckResourceAttrSet("data.mattermost_channel.sut", "creator_id"),
				),
			},
		},
	})
}

func testAccDataSourceChannelConfig(suffix string) string {
	return fmt.Sprintf(`
resource "mattermost_team" "example_team" {
  name         = "examplename%[1]s"
  display_name = "Example display name"
  description  = "Example description"
}

resource "mattermost_channel" "example_channel" {
  name         = "examplename%[1]s"
  display_name = "Example display name"
  header       = "Example header"
  team_id      = mattermost_team.example_team.id
}

data "mattermost_channel" "sut" {
  name    = mattermost_channel.example_channel.name
  team_id = mattermost_team.example_team.id
}
`, suffix)
}
