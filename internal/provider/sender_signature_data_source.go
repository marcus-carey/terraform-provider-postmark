package provider

import (
	"context"
	"fmt"
	"strconv"
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

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (d *senderSignatureDataSource) readFromAPI(ctx context.Context, sender *datasource_sender_signature.SenderSignatureModel) diag.Diagnostics {
	res, err := d.client.GetSenderSignature(ctx, TypeStringToInt64(sender.Id))
	if err != nil {
		clientDiag := diag.NewErrorDiagnostic("Client Error", fmt.Sprintf("Unable to read sender signature, got error: %s", err))
		return diag.Diagnostics{clientDiag}
	}

	// Set the data
	sender.Id = types.StringValue(strconv.Itoa(int(res.ID)))
	sender.Name = types.StringValue(res.Name)
	sender.FromEmail = types.StringValue(res.FromEmail)
	sender.ReplyToEmail = types.StringValue(res.ReplyToEmail)
	sender.Confirmed = types.BoolValue(res.Confirmed)
	sender.SpfHost = types.StringValue(res.SPFHost)
	sender.SpfTextValue = types.StringValue(res.SPFTextValue)
	sender.DkimVerified = types.BoolValue(res.DKIMVerified)
	sender.WeakDkim = types.BoolValue(res.WeakDKIM)
	sender.DkimHost = types.StringValue(res.DKIMHost)
	sender.DkimTextValue = types.StringValue(res.DKIMTextValue)
	sender.DkimPendingHost = types.StringValue(res.DKIMPendingHost)
	sender.DkimPendingTextValue = types.StringValue(res.DKIMPendingTextValue)
	sender.DkimRevokedHost = types.StringValue(res.DKIMRevokedHost)
	sender.DkimRevokedTextValue = types.StringValue(res.DKIMRevokedTextValue)
	sender.SafeToRemoveRevokedKeyFromDns = types.BoolValue(res.SafeToRemoveRevokedKeyFromDNS)
	sender.DkimUpdateStatus = types.StringValue(res.DKIMUpdateStatus)
	sender.ReturnPathDomain = types.StringValue(res.ReturnPathDomain)
	sender.ReturnPathDomainVerified = types.BoolValue(res.ReturnPathDomainVerified)
	sender.ReturnPathDomainCnameValue = types.StringValue(res.ReturnPathDomainCNAMEValue)
	sender.ConfirmationPersonalNote = types.StringValue(res.ConfirmationPersonalNote)

	return nil
}
