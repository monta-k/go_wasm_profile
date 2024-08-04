terraform {
  required_version = ">= 1.9"

  cloud {
    organization = "mnt_profile" 
    hostname = "app.terraform.io"

    workspaces { 
      name = "mnt_profile" 
    } 
  }

  required_providers {
    cloudflare = {
      source = "cloudflare/cloudflare"
      version = "~> 4"
    }
  }
}

provider "cloudflare" {
  api_token = var.cloudflare_api_token
}