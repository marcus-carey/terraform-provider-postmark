# Webhooks belong to a server, so import them as <server_id>:<webhook_id>.
# The provider looks up the server's API token with the account token.
terraform import postmark_webhook.example 12345:123
