package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-provider-mattermost/internal/testutils"
)

func TestAccResourceUser(t *testing.T) {
	t.Run("simple", func(t *testing.T) {
		suffix := acctest.RandString(16)

		resource.ParallelTest(t, resource.TestCase{
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config_TestAccResourceUser_simple(suffix),
					Check: resource.ComposeAggregateTestCheckFunc(
						testutils.TestCheckResourceAttrf("mattermost_user.sut", "email", "sut-%s@example.com", suffix),
						testutils.TestCheckResourceAttrf("mattermost_user.sut", "username", "sut-%s-username", suffix),
						resource.TestCheckResourceAttr("mattermost_user.sut", "first_name", ""),
						resource.TestCheckResourceAttr("mattermost_user.sut", "last_name", ""),
						resource.TestCheckResourceAttr("mattermost_user.sut", "nickname", ""),
						resource.TestCheckNoResourceAttr("mattermost_user.sut", "auth_data"),
						resource.TestCheckResourceAttr("mattermost_user.sut", "auth_service", ""),
						resource.TestCheckResourceAttr("mattermost_user.sut", "password", "example password"),
						resource.TestCheckResourceAttrSet("mattermost_user.sut", "locale"),
					),
				},
			},
		})
	})

	t.Run("full", func(t *testing.T) {
		suffix := acctest.RandString(16)

		resource.ParallelTest(t, resource.TestCase{
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config_TestAccResourceUser_full(suffix),
					Check: resource.ComposeAggregateTestCheckFunc(
						testutils.TestCheckResourceAttrf("mattermost_user.sut", "email", "sut-%s@example.com", suffix),
						testutils.TestCheckResourceAttrf("mattermost_user.sut", "username", "sut-%s-username", suffix),
						resource.TestCheckResourceAttr("mattermost_user.sut", "first_name", "Example first name"),
						resource.TestCheckResourceAttr("mattermost_user.sut", "last_name", "Example last name"),
						resource.TestCheckResourceAttr("mattermost_user.sut", "nickname", "Example nickname"),
						testutils.TestCheckResourceAttrf("mattermost_user.sut", "auth_data", "sut-%s@googledomain.com", suffix),
						resource.TestCheckResourceAttr("mattermost_user.sut", "auth_service", "google"),
						resource.TestCheckNoResourceAttr("mattermost_user.sut", "password"),
						resource.TestCheckResourceAttr("mattermost_user.sut", "locale", "fr"),
					),
				},
			},
		})
	})
}

func config_TestAccResourceUser_simple(suffix string) string {
	return fmt.Sprintf(`
resource "mattermost_user" "sut" {
  email    = "sut-%[1]s@example.com"
  username = "sut-%[1]s-username"
  password = "example password"
}
`, suffix)
}

func config_TestAccResourceUser_full(suffix string) string {
	return fmt.Sprintf(`
resource "mattermost_user" "sut" {
  email        = "sut-%[1]s@example.com"
  username     = "sut-%[1]s-username"
  first_name   = "Example first name"
  last_name    = "Example last name"
  nickname     = "Example nickname"
  auth_data    = "sut-%[1]s@googledomain.com"
  auth_service = "google"
  locale       = "fr"
  props = {
    foo = "bar"
  }
}
`, suffix)
}
