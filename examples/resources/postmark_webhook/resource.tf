data "postmark_server" "example" {
  name = "Test-Terraform"
}

resource "postmark_webhook" "example" {
  server_api_token = data.postmark_server.example.api_tokens[0]
  url              = "http://www.example.com/webhook-test-tracking"
  message_stream   = "outbound"

  http_auth = {
    username = "user"
    password = "pass"
  }

  http_headers = [
    {
      name  = "name"
      value = "value"
    }
  ]

  open_trigger = {
    enabled              = true
    post_first_open_only = false
  }

  click_trigger = {
    enabled = true
  }

  delivery_trigger = {
    enabled = true
  }

  bounce_trigger = {
    enabled         = false
    include_content = false
  }

  spam_complaint_trigger = {
    enabled         = false
    include_content = false
  }

  subscription_change_trigger = {
    enabled = false
  }
}
