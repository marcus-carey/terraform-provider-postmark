// The Postmark server data source allows access to the details of a server.
// This data source can be used to get information about an existing server in your Postmark account.

// Lookup by Server ID
data "postmark_server" "example" {
  id = 1111111
}

// Lookup by Server Name
data "postmark_server" "example" {
  name = "Test"
}
