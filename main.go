//go:generate terraform fmt -recursive ./examples/
//go:generate go run github.com/hashicorp/terraform-plugin-docs/cmd/tfplugindocs
//go:generate go run github.com/katbyte/terrafmt --verbose --fmtcompat fmt internal/provider

package main

import (
	"context"
	"flag"
	"fmt"
	"os"

	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-provider-mattermost/internal/provider"
)

var (
	version = "unknown"
	commit  = "unknown"
)

func main() {
	var debug bool
	flag.BoolVar(&debug, "debug", false, "set to true to run the provider with support for debuggers like delve")
	flag.Parse()

	ctx := context.Background()
	err := tfsdk.Serve(ctx, provider.New(version), tfsdk.ServeOpts{
		Name:  "registry.terraform.io/ndrpnt/mattermost",
		Debug: debug,
	})
	if err != nil {
		tflog.Error(ctx, fmt.Sprintf("Cannot serve provider: %v", err))
		os.Exit(1)
	}
}
