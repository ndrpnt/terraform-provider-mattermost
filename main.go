//go:generate terraform fmt -recursive ./examples/
//go:generate go run github.com/hashicorp/terraform-plugin-docs/cmd/tfplugindocs
//go:generate go run github.com/katbyte/terrafmt --verbose --fmtcompat fmt internal/provider

package main

import (
	"flag"

	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
	"github.com/hashicorp/terraform-provider-mattermost/internal/provider"
)

var (
	version = "dev"
	commit  = "unknown"
)

func main() {
	var debug bool
	flag.BoolVar(&debug, "debug", false, "set to true to run the provider with support for debuggers like delve")
	flag.Parse()

	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: provider.New(version),
		ProviderAddr: "registry.terraform.io/ndrpnt/mattermost",
		Debug:        debug,
	})
}
