package main

import (
	"github.com/AkxenTech/auth0-terraform-teamsapi/terraform-provider-auth0teams/internal/provider"
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: provider.Provider,
	})
}
