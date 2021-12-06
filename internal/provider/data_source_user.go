package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/mattermost/mattermost-server/v6/model"
)

func dataSourceUser() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceUserRead,

		Schema: map[string]*schema.Schema{
			"username": {
				Type:     schema.TypeString,
				Required: true,
			},
			"email": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"first_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"last_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceUserRead(_ context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*model.Client4)
	name := d.Get("username").(string)

	user, _, err := c.GetUserByUsername(name, "")
	if err != nil {
		return diag.Errorf("cannot get user by name: %v", err)
	}

	if user == nil {
		return diag.Errorf("user %s not found", name)
	}

	d.SetId(user.Id)
	d.Set("email", user.Email)
	d.Set("first_name", user.FirstName)
	d.Set("last_name", user.LastName)

	return nil
}
