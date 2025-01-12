resource "postmark_domain" "example" {
  name               = "example.com"
  return_path_domain = "pm-bounces.example.com"
}
