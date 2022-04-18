package provider

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/mattermost/mattermost-server/v6/model"
)

func fmtErr(resp *model.Response, err error) error {
	if resp == nil {
		return err
	}

	return fmt.Errorf("request %s failed with status %d: %v", resp.RequestId, resp.StatusCode, err)
}

func expandStringMap(m map[string]interface{}) map[string]string {
	r := make(map[string]string, len(m))
	for k, v := range m {
		r[k] = fmt.Sprintf("%v", v) // works for most types, unlike v.(string)
	}

	return r
}

// Check whether an attribute is "set", i.e. both known and not null.
//
// The implementation is quite wasteful but there doesn't seem to be a batter
// API available to check for Null/Unknown generically on all types.
func isSet(x attr.Value) bool {
	v, err := x.ToTerraformValue(nil)
	return err != nil && v.IsFullyKnown() && !v.IsNull()
}
