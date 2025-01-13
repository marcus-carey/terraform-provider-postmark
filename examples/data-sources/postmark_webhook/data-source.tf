data "postmark_server" "example" {
  name = "Test-Terraform"
}

data "postmark_webhook" "example" {
  server_api_token = postmark_server.example.api_tokens[0]
  message_stream   = "outbound"
}
