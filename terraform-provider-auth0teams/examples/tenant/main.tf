terraform {
  required_providers {
    auth0teams = {
      source  = "auth0/auth0teams"
      version = "0.1.0"
    }
  }
}

variable "api_token" {
  description = "API token to authenticate with Auth0 Teams"
  type        = string
}

variable "team_slug" {
  description = "The slug of the team"
  type        = string
}

provider "auth0teams" {
  api_token = var.api_token
  team_slug = var.team_slug
}

resource "auth0teams_tenant" "auth0_acme" {
  tenant_name      = "acme-development-tenant"
  admin_email      = "user@example.com"
  region           = "us"
  environment_type = "development"
}

output "tenant_id" {
  value = auth0teams_tenant.auth0_acme.id
}

output "client_name" {
  value     = auth0teams_tenant.auth0_acme.client_name
}

output "client_id" {
  value     = auth0teams_tenant.auth0_acme.client_id
}

output "client_secret" {
  value     = auth0teams_tenant.auth0_acme.client_secret
  sensitive = true
}