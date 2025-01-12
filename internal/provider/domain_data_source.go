package provider

import (
	"context"
	"terraform-provider-postmark/internal/provider/datasource_domain"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/mrz1836/postmark"
)

var _ datasource.DataSource = (*domainDataSource)(nil)

func NewDomainDataSource() datasource.DataSource {
	return &domainDataSource{}
}

type domainDataSource struct {
	client *postmark.Client
}

func (d *domainDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_domain"
}

func (d *domainDataSource) Schema(ctx context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = datasource_domain.DomainDataSourceSchema(ctx)
}

func (d *domainDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	d.client = GetDataSourcePostmarkClient(req, resp)
}

func (d *domainDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data datasource_domain.DomainModel

	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Read API call logic
	resp.Diagnostics.Append(d.readFromAPI(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Example data value setting
	data.Id = types.StringValue("example-id")

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (d *domainDataSource) readFromAPI(_ context.Context, _ *datasource_domain.DomainModel) diag.Diagnostics {
	return diag.Diagnostics{diag.NewErrorDiagnostic("Not Implemented", "This function is not implemented yet")}
	// TODO - Implement the API call to read the domain
	/*	res, err := d.client.GetDomain(ctx, domain.Id.ValueString())
		if err != nil {
			clientDiag := diag.NewErrorDiagnostic("Client Error", fmt.Sprintf("Unable to read domain, got error: %s", err))
			return diag.Diagnostics{clientDiag}
		}

		// TODO - Implement the API response to the model

		return nil*/
}
