package testutils

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestCheckResourceAttrf(name, key, format string, a ...interface{}) resource.TestCheckFunc {
	return resource.TestCheckResourceAttr(name, key, fmt.Sprintf(format, a...))
}
