resource "cloudflare_worker_script" "mnt_profile" {
  account_id = var.cloudflare_account_id
  name = "mnt_profile"
  content = file("${path.module}/worker.js")
}