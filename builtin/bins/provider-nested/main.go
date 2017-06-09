package main

import (
	"github.com/hashicorp/terraform/builtin/providers/nested"
	"github.com/hashicorp/terraform/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: nested.Provider,
	})
}
