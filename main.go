package main

import (
	"github.com/cgrotz/terraform-provider-bis/bis"

	"github.com/hashicorp/terraform-plugin-sdk/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: bis.Provider})
}
