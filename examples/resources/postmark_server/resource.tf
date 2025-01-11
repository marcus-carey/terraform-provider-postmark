resource "postmark_server" "test" {
  name        = "Test-Terraform"
  color       = "green"
  track_links = "HtmlAndText"
  track_opens = true
}
