package provider

import (
	"context"
	"terraform-provider-postmark/internal/provider/datasource_webhook"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/mrz1836/postmark"
)

var _ datasource.DataSource = (*webhookDataSource)(nil)

func NewWebhookDataSource() datasource.DataSource {
	return &webhookDataSource{}
}

type webhookDataSource struct {
	client *postmark.Client
}

func (d *webhookDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_webhook"
}

func (d *webhookDataSource) Schema(ctx context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = datasource_webhook.WebhookDataSourceSchema(ctx)
}

func (d *webhookDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	d.client = GetDataSourcePostmarkClient(req, resp)
}

func (d *webhookDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data datasource_webhook.WebhookModel

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

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (d *webhookDataSource) readFromAPI(_ context.Context, webhook *datasource_webhook.WebhookModel) diag.Diagnostics {
	d.client.ServerToken = webhook.ServerApiToken.ValueString()
	return diag.Diagnostics{diag.NewErrorDiagnostic("Not Implemented", "This function is not implemented yet")}
	// TODO - Implement the API call to read the webhook signature
	/*	res, err := d.client.GetSenderSignature(ctx, sender.Id.ValueString())
		if err != nil {
			clientDiag := diag.NewErrorDiagnostic("Client Error", fmt.Sprintf("Unable to read domain, got error: %s", err))
			return diag.Diagnostics{clientDiag}
		}

		// TODO - Implement the API response to the model

		return nil*/
}
