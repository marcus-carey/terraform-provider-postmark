terraform {
  required_providers {
    postmark = {
      source = "hashicorp.com/firebend/postmark"
    }
  }
}

variable "postmark_account_token" {
    type = string
}

provider "postmark" {
  account_token = var.postmark_account_token
}

data "postmark_server" "test" {
    id = 15125924
}

resource "postmark_server" "test" {
  name        = "Test-Terraform"
  color       = "green"
  track_links = "HtmlAndText"
  track_opens = true
}
