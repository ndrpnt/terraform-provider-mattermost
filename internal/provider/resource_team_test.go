package provider

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/mattermost/mattermost/server/public/model"
)

func TestAccResourceTeam(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccResourceTeam,
				Check: resource.ComposeAggregateTestCheckFunc(
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
  description  = "bar description"
  display_name = "bar display"
  name         = "bar"
}
`

func TestAccResourceTeam_disappears(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config:             testAccResourceTeam_disappears,
				Check:              testAccResourceTeamDisappears("mattermost_team.drift"),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func testAccResourceTeamDisappears(name string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		c := testProvider.Meta().(*model.Client4)
		rs, ok := s.RootModule().Resources[name]
		if !ok {
			return fmt.Errorf("not found: %s", name)
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf("no ID set on %s", name)
		}
		if _, err := c.PermanentDeleteTeam(context.Background(), rs.Primary.ID); err != nil {
			return fmt.Errorf("cannot delete team out of band: %w", err)
		}
		return nil
	}
}

const testAccResourceTeam_disappears = `
resource "mattermost_team" "drift" {
  description  = "drift description"
  display_name = "drift display"
  name         = "drift"
}
`
