package main

import (
	"github.com/benjvi/terraform-provider-tkgi/tkgi"
	"github.com/hashicorp/terraform-plugin-sdk/plugin"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: func() terraform.ResourceProvider {
			return tkgi.Provider()
		},
	})
}
