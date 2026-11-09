package provider

import "github.com/mrz1836/postmark"

type postmarkClientForTest struct{}

func (postmarkClientForTest) client(serverToken, accountToken string) *postmark.Client {
	return postmark.NewClient(serverToken, accountToken)
}
