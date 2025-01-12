resource "postmark_server" "example" {
  name                           = "Staging Testing"
  color                          = "red"
  smtp_api_activated             = true
  raw_email_enabled              = false
  delivery_type                  = "Live"
  inbound_hook_url               = "http://hooks.example.com/inbound"
  bounce_hook_url                = "http://hooks.example.com/bounce"
  open_hook_url                  = "http://hooks.example.com/open"
  delivery_hook_url              = "http://hooks.example.com/delivery"
  post_first_open_only           = false
  inbound_domain                 = ""
  inbound_spam_threshold         = 5
  track_opens                    = false
  track_links                    = "None"
  include_bounce_content_in_hook = true
  click_hook_url                 = "http://hooks.example.com/click"
  enable_smtp_api_error_hooks    = false
}
