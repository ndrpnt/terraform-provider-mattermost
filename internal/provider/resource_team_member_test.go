package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccResourceTeamMember(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccResourceTeamMember,
				Check:  resource.ComposeAggregateTestCheckFunc(),
			},
		},
	})
}

const testAccResourceTeamMember = `
resource "mattermost_team" "test" {
	name = "sheh"
	display_name = "An sheh bis"
}

data "mattermost_user" "test" {
  username = "admin"
}

resource "mattermost_team_member" "foo" {
	team_id = mattermost_team.test.id
	user_id = data.mattermost_user.test.id
}
`
