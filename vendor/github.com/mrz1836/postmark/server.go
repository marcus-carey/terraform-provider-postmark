package postmark

import (
	"context"
	"net/http"
)

// GetCurrentServer gets details for the server associated
// with the currently in-use server API Key
func (client *Client) GetCurrentServer(ctx context.Context) (Server, error) {
	res := Server{}
	err := client.doRequest(ctx, parameters{
		Method:    http.MethodGet,
		Path:      "server",
		TokenType: serverToken,
	}, &res)

	return res, err
}

// EditCurrentServer updates details for the server associated
// with the currently in-use server API Key
func (client *Client) EditCurrentServer(ctx context.Context, server Server) (Server, error) {
	// res := Server{}
	err := client.doRequest(ctx, parameters{
		Method:    http.MethodPut,
		Path:      "server",
		TokenType: serverToken,
	}, &server)
	return server, err
}
