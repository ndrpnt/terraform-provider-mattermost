package provider

import (
	"fmt"

	"github.com/mattermost/mattermost-server/v6/model"
)

func fmtErr(resp *model.Response, err error) error {
	return fmt.Errorf("request %s failed with status %d: %v", resp.RequestId, resp.StatusCode, err)
}
