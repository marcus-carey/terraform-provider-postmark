package provider

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/mrz1836/postmark"
)

func GetDataSourcePostmarkClient(req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) *postmark.Client {
	if req.ProviderData == nil {
		return nil
	}

	client, ok := req.ProviderData.(*postmark.Client)

	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Resource Configure Type",
			fmt.Sprintf("Expected *http.Client, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return nil
	}

	return client
}
