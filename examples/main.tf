terraform {
  required_providers {
    mattermost = {
      source = "registry.terraform.io/ndrpnt/mattermost"
    }
  }
}

provider "mattermost" {
  url      = "http://localhost:8065"
  login_id = "admin"
  password = "admin"
}

resource "mattermost_user" "simple" {
  email    = "simple@example.com"
  username = "simple"
  password = "secure123"
}

resource "mattermost_user" "complex" {
  email        = "jdoe@example.com"
  username     = "john-doe"
  first_name   = "John"
  last_name    = "Doe"
  nickname     = "JD"
  auth_service = "google"
  auth_data    = "john.doe@gmail.com"
  locale       = "fr"
  props = {
    customStatus = jsonencode({
      duration   = "this_week"
      emoji      = "palm_tree"
      expires_at = "2022-01-01T22:59:59.999Z"
      text       = "On a vacation"
    })
  }
}