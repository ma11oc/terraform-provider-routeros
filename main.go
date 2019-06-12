package main

import (
	"github.com/hashicorp/terraform/plugin"
	"github.com/ma11oc/terraform-provider-routeros/routeros"
)

func main() {
	defer func() {
		if routeros.RouterOSClient != nil {
			routeros.RouterOSClient.Close()
			routeros.RouterOSClient = nil
		}
	}()

	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: routeros.Provider,
	})
}
