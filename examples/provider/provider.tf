variable "postmark_account_token" {
  type = string
}

provider "postmark" {
  account_token = var.postmark_account_token
}
