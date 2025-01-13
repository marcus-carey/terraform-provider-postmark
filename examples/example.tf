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

variable "existing_server_name" {
  type = string
}

variable "existing_domain_id" {
  type = string
}

variable "existing_sender_signature_id" {
  type = string
}

provider "postmark" {
  account_token = var.postmark_account_token
}

data "postmark_server" "existing" {
  # This is the Name of an existing server in your Postmark account
  name = var.existing_server_name
}

data "postmark_domain" "existing" {
  # This is the ID of an existing domain in your Postmark account
  id = var.existing_domain_id
}

data "postmark_sender_signature" "existing" {
  # This is the ID of an existing sender signature in your Postmark account
  id = var.existing_sender_signature_id
}

data "postmark_webhook" "outbound" {
  server_api_token = data.postmark_server.existing.api_tokens[0]
  message_stream   = "outbound"
}

# resource "postmark_webhook" "example" {
#   server_api_token = data.postmark_server.existing.api_tokens[0]
#   url              = "http://www.example.com/webhook-test-tracking"
#   message_stream   = "outbound"
#
#   http_auth = {
#     username = "user"
#     password = "pass"
#   }
# 
#   http_headers = [
#     {
#       name  = "x-custom-header"
#       value = "test_value"
#     }
#   ]
#
#   open_trigger = {
#     enabled              = true
#     post_first_open_only = false
#   }
#
#   click_trigger = {
#     enabled = true
#   }
#
#   delivery_trigger = {
#     enabled = true
#   }
#
#   bounce_trigger = {
#     enabled         = false
#     include_content = false
#   }
#
#   spam_complaint_trigger = {
#     enabled         = false
#     include_content = false
#   }
#
#   subscription_change_trigger = {
#     enabled = false
#   }
# }

# resource "postmark_server" "test_create" {
#   name                   = "Test-Terraform"
#   color                  = "green"
#   delivery_type          = "Live"
#   inbound_hook_url       = ""
#   track_links            = "HtmlAndText"
#   inbound_spam_threshold = 5
#   track_opens            = true
#   post_first_open_only   = false
#   raw_email_enabled      = false
#   smtp_api_activated     = true
# }

# resource "postmark_domain" "test_create" {
#   name               = "test.${data.postmark_domain.existing.name}"
#   return_path_domain = "pm-bounces.test.${data.postmark_domain.existing.name}"
# }

# resource "postmark_sender_signature" "test_create" {
#   name                       = "Terraform"
#   from_email                 = "terraform@${data.postmark_domain.existing.name}"
# }


