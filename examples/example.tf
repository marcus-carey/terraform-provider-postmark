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
  name = "Test-Terraform"
}

data "postmark_domain" "existing" {
  # This is the ID of an existing domain in your Postmark account
  id = var.existing_domain_id
}

data "postmark_sender_signature" "existing" {
  # This is the ID of an existing sender signature in your Postmark account
  id = var.existing_sender_signature_id
}

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

output "postmark_server_info_existing" {
  value = "The Postmark server name is ${data.postmark_server.existing.name} and the color is ${data.postmark_server.existing.color}.\nInbound email tracking is set to ${data.postmark_server.existing.track_links} and open tracking is set to ${data.postmark_server.existing.track_opens}.\nEmails sent to this server will be sent to the following email address: ${data.postmark_server.existing.inbound_address}."
}

# output "postmark_server_info_create" {
#   value = "The Postmark server name is ${postmark_server.test_create.name} and the color is ${postmark_server.test_create.color}.\nInbound email tracking is set to ${postmark_server.test_create.track_links} and open tracking is set to ${postmark_server.test_create.track_opens}.\nEmails sent to this server will be sent to the following email address: ${postmark_server.test_create.inbound_address}."
# }

output "postmark_domain_info_existing" {
  value = "The Postmark domain name is ${data.postmark_domain.existing.name} and the ID is ${data.postmark_domain.existing.id}. Return path domain is set to ${data.postmark_domain.existing.return_path_domain} using cname value ${data.postmark_domain.existing.return_path_domain_cname_value}."
}

# output "postmark_domain_info_create" {
#   value = "The Postmark domain name is ${postmark_domain.test_create.name} and the ID is ${postmark_domain.test_create.id}. Return path domain is set to ${postmark_domain.test_create.return_path_domain} using cname value ${postmark_domain.test_create.return_path_domain_cname_value}."
# }

output "postmark_sender_signature_info_existing" {
  value = "The Postmark sender signature ID is ${data.postmark_sender_signature.existing.id} and the email address is ${data.postmark_sender_signature.existing.from_email}."
}

# output "postmark_sender_signature_info_create" {
#   value = "The Postmark sender signature ID is ${postmark_sender_signature.test_create.id} and the email address is ${postmark_sender_signature.test_create.from_email}."
# }


