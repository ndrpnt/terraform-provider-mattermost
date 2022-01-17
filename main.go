//go:generate terraform fmt -recursive ./examples/
//go:generate go run github.com/hashicorp/terraform-plugin-docs/cmd/tfplugindocs
//go:generate go run github.com/katbyte/terrafmt --verbose --fmtcompat fmt internal/provider

package main

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-provider-mattermost/internal/provider"
)

var (
	version = "dev"
	commit  = "unknown"
)

func main() {
	opts := tfsdk.ServeOpts{
		Name: "registry.terraform.io/ndrpnt/mattermost",
	}

	err := tfsdk.Serve(context.Background(), provider.New(version), opts)
	if err != nil {
		log.Fatal(err.Error())
	}
}
