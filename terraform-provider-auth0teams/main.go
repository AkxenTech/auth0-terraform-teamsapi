package main

import (
	"github.com/atko-scratch/auth0-teams-api-terraform-hack25/terraform-provider-auth0teams/internal/provider"
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: provider.Provider,
	})
}
