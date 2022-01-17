package provider

import (
	"context"
	"net/http"

	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
)

const (
	conflictsWithErr         = "There was a conflict detected."
	conflictsWithDescription =
)

type conflictsWithValidator struct {
	conflicts []*tftypes.AttributePath
}

func ConflictsWith(paths ...*tftypes.AttributePath) tfsdk.AttributeValidator {
	http.HandlerFunc()
	return conflictsWithValidator{
		conflicts: paths,
	}
}

// /////////

type Validator func(context.Context, tfsdk.ValidateAttributeRequest, *tfsdk.ValidateAttributeResponse)

func (v Validator) Description(context.Context) string {
	return "No description provided."
}

func (v Validator) MarkdownDescription(context.Context) string {
	return "No description provided."
}

func (v Validator) Validate(ctx context.Context, req tfsdk.ValidateAttributeRequest, resp *tfsdk.ValidateAttributeResponse) {
	v(ctx, req, resp)
}
