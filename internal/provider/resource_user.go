package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/mattermost/mattermost-server/v6/model"
)

// Note that:
// 1) `auth_data` and `password` cannot be read. Thus, Terraform's refresh and
//    drift detection are not working as expected for these fields.
// 2) `notify_props` can only be read for the authenticated user. This provider
//    doesn't manage them at all, as there is little benefit to it anyway.
func resourceUser() *schema.Resource {
	return &schema.Resource{
		Description: "Manage a user.",

		CreateContext: resourceUserCreate,
		ReadContext:   resourceUserRead,
		UpdateContext: resourceUserUpdate,
		DeleteContext: resourceUserDelete,
		Importer:      &schema.ResourceImporter{StateContext: schema.ImportStatePassthroughContext},

		Schema: map[string]*schema.Schema{
			"email": {
				Type:     schema.TypeString,
				Required: true,
			},
			"username": {
				Type:     schema.TypeString,
				Required: true,
			},
			"first_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"last_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"nickname": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"auth_data": {
				Type:         schema.TypeString,
				Optional:     true,
				Description:  "Service-specific authentication data, such as email address.",
				ForceNew:     true,
				RequiredWith: []string{"auth_service"},
			},
			"auth_service": {
				Type:             schema.TypeString,
				Optional:         true,
				Description:      `The authentication service, one of "email", "gitlab", "ldap", "saml", "office365", and "google".`,
				ForceNew:         true,
				ValidateDiagFunc: validation.ToDiagFunc(validation.StringInSlice([]string{"email", "gitlab", "ldap", "saml", "office365", "google"}, false)),
				ExactlyOneOf:     []string{"auth_service", "password"},
				RequiredWith:     []string{"auth_data"},
			},
			"password": {
				Type:         schema.TypeString,
				Optional:     true,
				Sensitive:    true,
				Description:  "The password used for email authentication.",
				ExactlyOneOf: []string{"auth_service", "password"},
			},
			"locale": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"props": {
				Type:     schema.TypeMap,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func resourceUserCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*model.Client4)

	user := &model.User{
		Username:    d.Get("username").(string),
		Password:    d.Get("password").(string),
		AuthService: d.Get("auth_service").(string),
		Email:       d.Get("email").(string),
		Nickname:    d.Get("nickname").(string),
		FirstName:   d.Get("first_name").(string),
		LastName:    d.Get("last_name").(string),
		Locale:      d.Get("locale").(string),
	}
	if authData, ok := d.GetOk("auth_data"); ok {
		ad := authData.(string)
		user.AuthData = &ad
	}
	if props, ok := d.GetOk("props"); ok {
		user.Props = expandStringMap(props.(map[string]interface{}))
	}

	user, resp, err := c.CreateUser(user)
	if err != nil {
		return diag.Errorf("cannot create user: %v", fmtErr(resp, err))
	}

	d.SetId(user.Id)

	return resourceUserRead(ctx, d, meta)
}

func resourceUserRead(_ context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*model.Client4)

	user, resp, err := c.GetUser(d.Id(), "")
	if resp.StatusCode == 404 {
		d.SetId("")
		return nil
	}
	if err != nil {
		return diag.Errorf("cannot get user: %v", fmtErr(resp, err))
	}

	d.Set("username", user.Username)
	d.Set("auth_service", user.AuthService)
	d.Set("email", user.Email)
	d.Set("nickname", user.Nickname)
	d.Set("first_name", user.FirstName)
	d.Set("last_name", user.LastName)
	d.Set("locale", user.Locale)
	d.Set("props", user.Props)

	return nil
}

func resourceUserUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*model.Client4)

	if d.HasChange("password") {
		oldPassword, newPassword := d.GetChange("password")
		resp, err := c.UpdatePassword(d.Id(), oldPassword.(string), newPassword.(string))
		if err != nil {
			return diag.Errorf("cannot update password: %v", fmtErr(resp, err))
		}
	}

	user := &model.User{
		CreateAt:  -1, // Mattermost doesn't persist this value but returns an HTTP 400 error if it is 0.
		Id:        d.Id(),
		Username:  d.Get("username").(string),
		Email:     d.Get("email").(string),
		Nickname:  d.Get("nickname").(string),
		FirstName: d.Get("first_name").(string),
		LastName:  d.Get("last_name").(string),
		Locale:    d.Get("locale").(string),
	}
	if props, ok := d.GetOk("props"); ok {
		user.Props = expandStringMap(props.(map[string]interface{}))
	}

	user, resp, err := c.UpdateUser(user)
	if err != nil {
		return diag.Errorf("cannot update user: %v", fmtErr(resp, err))
	}

	return resourceUserRead(ctx, d, meta)
}

func resourceUserDelete(_ context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*model.Client4)

	resp, err := c.PermanentDeleteUser(d.Id())
	if err != nil {
		return diag.Errorf("cannot delete user: %v", fmtErr(resp, err))
	}

	return nil
}
