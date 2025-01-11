package provider

import (
	"context"
	"terraform-provider-postmark/internal/provider/datasource_sender_signature"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ datasource.DataSource = (*senderSignatureDataSource)(nil)

func NewSenderSignatureDataSource() datasource.DataSource {
	return &senderSignatureDataSource{}
}

type senderSignatureDataSource struct{}

func (d *senderSignatureDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_sender_signature"
}

func (d *senderSignatureDataSource) Schema(ctx context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = datasource_sender_signature.SenderSignatureDataSourceSchema(ctx)
}

func (d *senderSignatureDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data datasource_sender_signature.SenderSignatureModel

	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Read API call logic

	// Example data value setting
	data.Id = types.StringValue("example-id")

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
