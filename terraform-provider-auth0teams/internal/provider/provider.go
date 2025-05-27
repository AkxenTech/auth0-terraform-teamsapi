package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"api_token": {
				Type:      schema.TypeString,
				Required:  true,
				Sensitive: true,
			},
			"team_slug": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"auth0teams_tenant": resourceTenant(),
		},
		ConfigureContextFunc: configureProvider,
	}
}

func configureProvider(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	token := d.Get("api_token").(string)
	teamSlug := d.Get("team_slug").(string)
	return &Client{Token: token, TeamSlug: teamSlug}, nil
}

type Client struct {
	Token    string
	TeamSlug string
}
