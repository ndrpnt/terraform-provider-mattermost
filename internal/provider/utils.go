package provider

import (
	"fmt"

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
