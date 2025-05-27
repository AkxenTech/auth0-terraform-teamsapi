# auth0teams_tenant Resource

Manages an Auth0 Teams tenant using the Auth0 Teams API. This resource allows you to create and destroy tenants scoped to a specific team.

---

## Example Usage

```hcl
provider "auth0teams" {
  api_token = var.api_token
  team_slug = var.team_slug
}

resource "auth0teams_tenant" "example" {
  tenant_name      = "acme-dev"
  admin_email      = "admin@acme.com"
  region           = "us"
  environment_type = "development"
}
```

---

## Argument Reference

The following arguments are supported:

- `tenant_name` *(Optional)* – Friendly display name of the tenant.
- `admin_email` *(Required)* – Email address of the tenant administrator.
- `region` *(Required)* – Geographic region for the tenant (e.g. `us`, `eu`) for public cloud setups.
- `environment` *(Required)* – Slug of the environment, for private cloud setups.
- `environment_type` *(Optional)* – Type of environment (`development`, `staging`, `production`). Defaults to `"development"`.

---

## Attributes Reference

The following attributes are exported:

- `tenant_id` – Internal UUID assigned to the tenant.
- `tenant_fqdn` – Fully qualified domain name of the tenant.
- `created_at` – Timestamp of tenant creation in RFC3339 format.
- `client_id` – Client ID of the created management application.
- `client_secret` – Client secret of the management application (sensitive).
- `client_name` – Name of the management client associated with the tenant.

---

## Create

The tenant is created via a `POST /api/tenants` call with the supplied configuration values.

- If `environment_type` is not provided, the API may assign a default.
- On success, the provider sets the resource ID to `tenant_id` and returns all computed attributes.

---

## Delete

The tenant is deleted via a `DELETE /api/tenants/{tenant_id}` request.

This is triggered when:

- You run `terraform destroy`
- The resource block is removed from your configuration and `terraform apply` is executed

Terraform will:
- Send a DELETE request using the tenant ID
- Remove the resource from state
- Ignore 404 errors if the tenant has already been deleted externally
```