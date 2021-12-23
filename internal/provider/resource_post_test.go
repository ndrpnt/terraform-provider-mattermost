package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccResourcePost(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: postConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair("mattermost_post.example_post", "channel_id", "mattermost_channel.example_channel", "id"),
					resource.TestCheckResourceAttr("mattermost_post.example_post", "message", "Example post message"),
				),
			},
		},
	})
}

const postConfig = `
data "mattermost_user" "example_user" {
  username = "test"
}

resource "mattermost_team" "example_team" {
	name = "myexampleteam"
	display_name = "My example team"
}

resource "mattermost_team_member" "example_team_member" {
	team_id = mattermost_team.example_team.id
	user_id = data.mattermost_user.example_user.id
}

resource "mattermost_channel" "example_channel" {
	display_name = "My example channel"
	name = "myexamplechannel"
	team_id = mattermost_team.example_team.id
}

resource "mattermost_post" "example_post" {
	channel_id = mattermost_channel.example_channel.id
	message = "Example post message"
}
`
