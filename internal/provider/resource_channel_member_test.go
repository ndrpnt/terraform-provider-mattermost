package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccResourceChannelMember(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccResourceChannelMember,
				Check:  resource.ComposeTestCheckFunc(),
			},
		},
	})
}

const testAccResourceChannelMember = `
resource "mattermost_team" "test" {
	name = "sheh"
	display_name = "An sheh bis"
}

resource "mattermost_team_member" "foo" {
	team_id = mattermost_team.test.id
	user_id = data.mattermost_user.test.id
}

resource "mattermost_channel" "foo" {
	description = "bar description"
	display_name = "bar display"
  	name = "bar"
	team_id = mattermost_team.test.id
}

data "mattermost_user" "test" {
  username = "mattermostadmin"
}

resource "mattermost_channel_member" "foo" {
	channel_id = mattermost_channel.foo.id
	user_id = data.mattermost_user.test.id
}
`
