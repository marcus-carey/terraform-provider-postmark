terraform {
  required_version = ">= 1.3.1"
  required_providers {
    postmark = {
      source = "marcus.carey/terraform/postmark"
    }
  }
}

variable "postmark_account_token" {
  type = string
}

variable "existing_server_id" {
  type = string
}

provider "postmark" {
  account_token = var.postmark_account_token
}

data "postmark_server" "existing" {
  # This is the ID of an existing server in your Postmark account
  # id = var.existing_server_id

  # This is the Name of an existing server in your Postmark account
  name = "Test"
}

resource "postmark_server" "test_create" {
  name        = "Test-Terraform"
  color       = "green"
  track_links = "HtmlAndText"
  track_opens = true
}

output "test_name" {
  value = "Existing Postmark server name: ${data.postmark_server.existing.name}"
}

output "postmark_info" {
  value = "The Postmark server name is ${postmark_server.test_create.name} and the color is ${postmark_server.test_create.color}.\nInbound email tracking is set to ${postmark_server.test_create.track_links} and open tracking is set to ${postmark_server.test_create.track_opens}.\nEmails sent to this server will be sent to the following email address: ${postmark_server.test_create.inbound_address}."
}
