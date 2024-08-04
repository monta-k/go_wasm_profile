resource "cloudflare_worker_script" "mnt_profile" {
  name = "mnt_profile"
  
  content = file("${path.module}/worker.js")
}