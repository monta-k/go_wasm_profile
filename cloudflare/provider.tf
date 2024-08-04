terraform {
  required_version = ">= 1.9"

  required_providers {
    cloudflare = {
      source = "cloudflare/cloudflare"
      version = "~> 4"
    }
  }
}

provider "cloudflare" {
  email      = var.cloudflare_email
  api_token = var.cloudflare_api_token
}