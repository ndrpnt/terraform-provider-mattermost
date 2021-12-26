---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "mattermost_user Resource - terraform-provider-mattermost"
subcategory: ""
description: |-
  Manage a user.
---

# mattermost_user (Resource)

Manage a user.



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- **email** (String)
- **username** (String)

### Optional

- **auth_data** (String) Service-specific authentication data, such as email address.
- **auth_service** (String) The authentication service, one of "email", "gitlab", "ldap", "saml", "office365", "google", and "".
- **first_name** (String)
- **id** (String) The ID of this resource.
- **last_name** (String)
- **locale** (String)
- **nickname** (String)
- **notify_props** (Block List, Max: 1) (see [below for nested schema](#nestedblock--notify_props))
- **password** (String, Sensitive) The password used for email authentication.
- **props** (Map of String)

<a id="nestedblock--notify_props"></a>
### Nested Schema for `notify_props`

Optional:

- **channel** (Boolean) Set to "true" to enable channel-wide notifications (@channel, @all, etc.), "false" to disable. Defaults to "true".
- **desktop** (String) Set to "all" to receive desktop notifications for all activity, "mention" for mentions and direct messages only, and "none" to disable. Defaults to "all".
- **desktop_sound** (Boolean) Set to "true" to enable sound on desktop notifications, "false" to disable. Defaults to "true".
- **email** (Boolean) Set to "true" to enable email notifications, "false" to disable. Defaults to "true".
- **first_name** (Boolean) Set to "true" to enable mentions for first name. Defaults to "true" if a first name is set, "false" otherwise.
- **mention_keys** (String) A comma-separated list of words to count as mentions. Defaults to username and @username.
- **push** (String) Set to "all" to receive push notifications for all activity, "mention" for mentions and direct messages only, and "none" to disable. Defaults to "mention".

