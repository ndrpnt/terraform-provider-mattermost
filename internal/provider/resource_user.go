package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/mattermost/mattermost-server/v6/model"
)

func resourceUser() *schema.Resource {
	authMethods := []string{"auth_data", "password"}

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
				ExactlyOneOf: authMethods,
				ForceNew:     true,
			},
			"auth_service": {
				Type:         schema.TypeString,
				Optional:     true,
				Description:  `The authentication service, one of "email", "gitlab", "ldap", "saml", "office365", "google", and "".`,
				RequiredWith: []string{"auth_data"},
				ForceNew:     true,
			},
			"password": {
				Type:         schema.TypeString,
				Optional:     true,
				Sensitive:    true,
				Description:  "The password used for email authentication.",
				ExactlyOneOf: authMethods,
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
			"notify_props": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"email": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: `Set to "true" to enable email notifications, "false" to disable. Defaults to "true".`,
						},
						"push": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: `Set to "all" to receive push notifications for all activity, "mention" for mentions and direct messages only, and "none" to disable. Defaults to "mention".`,
						},
						"desktop": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: `Set to "all" to receive desktop notifications for all activity, "mention" for mentions and direct messages only, and "none" to disable. Defaults to "all".`,
						},
						"desktop_sound": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: `Set to "true" to enable sound on desktop notifications, "false" to disable. Defaults to "true".`,
						},
						"mention_keys": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: `A comma-separated list of words to count as mentions. Defaults to username and @username.`,
						},
						"channel": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: `Set to "true" to enable channel-wide notifications (@channel, @all, etc.), "false" to disable. Defaults to "true".`,
						},
						"first_name": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: `Set to "true" to enable mentions for first name. Defaults to "true" if a first name is set, "false" otherwise.`,
						},
					},
				},
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
	if notifyProps, ok := d.GetOk("notify_props"); ok {
		np := notifyProps.([]interface{})
		if len(np) > 0 {
			user.NotifyProps = expandStringMap(np[0].(map[string]interface{}))
		}
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
	d.Set("password", user.Password)
	d.Set("auth_service", user.AuthService)
	d.Set("email", user.Email)
	d.Set("nickname", user.Nickname)
	d.Set("first_name", user.FirstName)
	d.Set("last_name", user.LastName)
	d.Set("locale", user.Locale)
	d.Set("auth_data", user.AuthData)
	d.Set("props", user.Props)
	d.Set("notify_props", []interface{}{user.NotifyProps})

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
	if notifyProps, ok := d.GetOk("notify_props"); ok {
		np := notifyProps.([]interface{})
		if len(np) > 0 {
			user.NotifyProps = expandStringMap(np[0].(map[string]interface{}))
		}
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
