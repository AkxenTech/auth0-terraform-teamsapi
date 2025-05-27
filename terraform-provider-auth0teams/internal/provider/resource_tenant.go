package provider

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type CreateTenantRequest struct {
	TenantID        string `json:"tenant_id,omitempty"`
	TenantName      string `json:"tenant_name,omitempty"`
	AdminEmail      string `json:"admin_email"`
	Region          string `json:"region,omitempty"`
	Environment     string `json:"environment,omitempty"`
	EnvironmentType string `json:"environment_type,omitempty"`
}

type ManagementClient struct {
	ClientName   string `json:"client_name"`
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
}

type CreateTenantResponse struct {
	TenantID         string           `json:"tenant_id"`
	TenantFQDN       string           `json:"tenant_fqdn"`
	Environment      string           `json:"environment"`
	Region           string           `json:"region"`
	CreatedAt        string           `json:"created_at"`
	ManagementClient ManagementClient `json:"management_client"`
}

func resourceTenant() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceTenantCreate,
		ReadContext:   resourceTenantRead,
		UpdateContext: resourceTenantUpdate,
		DeleteContext: resourceTenantDelete,

		Schema: map[string]*schema.Schema{
			// Inputs
			"tenant_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"admin_email": {
				Type:     schema.TypeString,
				Required: true,
			},
			"region": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"environment": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"environment_type": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "development",
			},
			// Computed outputs
			"tenant_fqdn": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"created_at": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"client_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"client_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"client_secret": {
				Type:      schema.TypeString,
				Computed:  true,
				Sensitive: true,
			},
		},
	}
}

func resourceTenantCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*Client)

	payload := CreateTenantRequest{
		TenantName:      d.Get("tenant_name").(string),
		AdminEmail:      d.Get("admin_email").(string),
		Region:          d.Get("region").(string),
		Environment:     d.Get("environment").(string),
		EnvironmentType: d.Get("environment_type").(string),
	}

	body, err := json.Marshal(payload)
	if err != nil {
		return diag.FromErr(err)
	}

	url := fmt.Sprintf("https://%s.teams.auth0.com/api/tenants", client.TeamSlug)
	req, err := http.NewRequest("POST", url, strings.NewReader(string(body)))
	if err != nil {
		return diag.FromErr(err)
	}
	req.Header.Set("Authorization", "Bearer "+client.Token)
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return diag.FromErr(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 201 {
		return diag.Errorf("Failed to create tenant: %s", resp.Status)
	}

	var result CreateTenantResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return diag.FromErr(err)
	}

	d.SetId(result.TenantID)
	d.Set("tenant_id", result.TenantID)
	d.Set("tenant_fqdn", result.TenantFQDN)
	d.Set("environment", result.Environment)
	d.Set("created_at", result.CreatedAt)
	d.Set("client_name", result.ManagementClient.ClientName)
	d.Set("client_id", result.ManagementClient.ClientID)
	d.Set("client_secret", result.ManagementClient.ClientSecret)
	return nil
}

func resourceTenantRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return nil
}

func resourceTenantUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return nil
}

func resourceTenantDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*Client)
	id := d.Id()
	url := fmt.Sprintf("https://%s.teams.auth0.com/api/tenants/%s", client.TeamSlug, id)

	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return diag.FromErr(err)
	}
	req.Header.Set("Authorization", "Bearer "+client.Token)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return diag.FromErr(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 204 {
		return diag.Errorf("Failed to delete tenant: %s", resp.Status)
	}

	d.SetId("")
	return nil
}
