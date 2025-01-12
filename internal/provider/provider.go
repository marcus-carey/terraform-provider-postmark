package provider

import (
	"context"
	"terraform-provider-postmark/internal/provider/provider_postmark"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/mrz1836/postmark"
)

var _ provider.Provider = &postmarkProvider{}

func New(version string) func() provider.Provider {
	return func() provider.Provider {
		return &postmarkProvider{
			version: version,
		}
	}
}

type postmarkProvider struct {
	version string
}

func (p *postmarkProvider) Schema(ctx context.Context, _ provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = provider_postmark.PostmarkProviderSchema(ctx)
}

func (p *postmarkProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	var data provider_postmark.PostmarkModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	println("Account Token:", data.AccountToken.String())
	client := postmark.NewClient("[SERVER-TOKEN]", data.AccountToken.ValueString())

	// TODO determine how to test connection to fail earlier

	resp.DataSourceData = client
	resp.ResourceData = client
}

func (p *postmarkProvider) Metadata(_ context.Context, _ provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "postmark"
	resp.Version = p.version
}

func (p *postmarkProvider) DataSources(_ context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{
		NewDomainDataSource,
		NewSenderSignatureDataSource,
		NewServerDataSource,
		NewWebhookDataSource,
	}
}

func (p *postmarkProvider) Resources(_ context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		NewDomainResource,
		NewSenderSignatureResource,
		NewServerResource,
		NewWebhookResource,
	}
}
