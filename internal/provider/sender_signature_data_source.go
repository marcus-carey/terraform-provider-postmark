package provider

import (
	"context"
	"terraform-provider-postmark/internal/provider/datasource_sender_signature"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/mrz1836/postmark"
)

var _ datasource.DataSource = (*senderSignatureDataSource)(nil)

func NewSenderSignatureDataSource() datasource.DataSource {
	return &senderSignatureDataSource{}
}

type senderSignatureDataSource struct {
	client *postmark.Client
}

func (d *senderSignatureDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_sender_signature"
}

func (d *senderSignatureDataSource) Schema(ctx context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = datasource_sender_signature.SenderSignatureDataSourceSchema(ctx)
}

func (d *senderSignatureDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	d.client = GetDataSourcePostmarkClient(req, resp)
}

func (d *senderSignatureDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data datasource_sender_signature.SenderSignatureModel

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

func (d *senderSignatureDataSource) readFromAPI(_ context.Context, _ *datasource_sender_signature.SenderSignatureModel) diag.Diagnostics {
	return diag.Diagnostics{diag.NewErrorDiagnostic("Not Implemented", "This function is not implemented yet")}
	// TODO - Implement the API call to read the sender signature
	/*	res, err := d.client.GetSenderSignature(ctx, sender.Id.ValueString())
		if err != nil {
			clientDiag := diag.NewErrorDiagnostic("Client Error", fmt.Sprintf("Unable to read domain, got error: %s", err))
			return diag.Diagnostics{clientDiag}
		}

		// TODO - Implement the API response to the model

		return nil*/
}
