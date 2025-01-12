resource "postmark_sender_signature" "example" {
  name                       = "John Doe"
  from_email                 = "john.doe@example.com"
  reply_to_email             = "reply@example.com"
  return_path_domain         = "pm-bounces.example.com"
  confirmation_personal_note = "This is a note visible to the recipient to provide context of what Postmark is."
}
