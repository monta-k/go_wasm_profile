resource "cloudflare_worker_script" "mnt_profile" {
  account_id = var.cloudflare_account_id
  name = "mnt_profile"
  content = file("${path.module}/worker.js")
  module     = true
}

resource "cloudflare_worker_route" "mnt_profile_route" {
  pattern = "mnt_profile.monta-k0523.workers.dev"
  enabled = true
  script_name = cloudflare_worker_script.mnt_profile.name
}