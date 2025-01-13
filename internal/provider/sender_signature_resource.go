package provider

import (
	"context"
	"fmt"
	"strconv"
	"terraform-provider-postmark/internal/provider/resource_sender_signature"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/mrz1836/postmark"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ resource.Resource                = &senderSignatureResource{}
	_ resource.ResourceWithConfigure   = &senderSignatureResource{}
	_ resource.ResourceWithImportState = &senderSignatureResource{}
)

func NewSenderSignatureResource() resource.Resource {
	return &senderSignatureResource{}
}

type senderSignatureResource struct {
	client *postmark.Client
}

func (r *senderSignatureResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_sender_signature"
}

func (r *senderSignatureResource) Schema(ctx context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = resource_sender_signature.SenderSignatureResourceSchema(ctx)
}

func (r *senderSignatureResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	r.client = GetResourcePostmarkClient(req, resp)
}

func (r *senderSignatureResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	// Retrieve import ID and save to id attribute
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *senderSignatureResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data resource_sender_signature.SenderSignatureModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Create API call logic
	resp.Diagnostics.Append(r.createFromAPI(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *senderSignatureResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data resource_sender_signature.SenderSignatureModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Read API call logic
	resp.Diagnostics.Append(r.readFromAPI(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *senderSignatureResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data resource_sender_signature.SenderSignatureModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Update API call logic
	resp.Diagnostics.Append(r.updateFromAPI(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *senderSignatureResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data resource_sender_signature.SenderSignatureModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Delete API call logic
	resp.Diagnostics.Append(r.deleteFromAPI(ctx, &data)...)
}

func (r *senderSignatureResource) readFromAPI(ctx context.Context, senderSignature *resource_sender_signature.SenderSignatureModel) diag.Diagnostics {
	res, err := r.client.GetSenderSignature(ctx, TypeStringToInt64(senderSignature.Id))
	if err != nil {
		clientDiag := diag.NewErrorDiagnostic("Client Error", fmt.Sprintf("Unable to read sender signature, got error: %s", err))
		return diag.Diagnostics{clientDiag}
	}

	return mapSenderSignatureResourceFromAPI(ctx, senderSignature, res)
}

func (r *senderSignatureResource) createFromAPI(ctx context.Context, senderSignature *resource_sender_signature.SenderSignatureModel) diag.Diagnostics {
	body := postmark.SenderSignatureCreateRequest{
		FromEmail:                senderSignature.FromEmail.ValueString(),
		Name:                     senderSignature.Name.ValueString(),
		ReplyToEmail:             senderSignature.ReplyToEmail.ValueString(),
		ReturnPathDomain:         senderSignature.ReturnPathDomain.ValueString(),
		ConfirmationPersonalNote: senderSignature.ConfirmationPersonalNote.ValueString(),
	}
	res, err := r.client.CreateSenderSignature(ctx, body)
	if err != nil {
		clientDiag := diag.NewErrorDiagnostic("Client Error", fmt.Sprintf("Unable to create sender signature, got error: %s", err))
		return diag.Diagnostics{clientDiag}
	}

	if res.ID == 0 {
		clientDiag := diag.NewErrorDiagnostic("Client Error", "Unable to create sender signature, got error: SenderSignature ID is 0.")
		return diag.Diagnostics{clientDiag}
	}

	return mapSenderSignatureResourceFromAPI(ctx, senderSignature, res)
}

func (r *senderSignatureResource) updateFromAPI(ctx context.Context, senderSignature *resource_sender_signature.SenderSignatureModel) diag.Diagnostics {
	body := postmark.SenderSignatureEditRequest{
		Name:                     senderSignature.Name.ValueString(),
		ReplyToEmail:             senderSignature.ReplyToEmail.ValueString(),
		ReturnPathDomain:         senderSignature.ReturnPathDomain.ValueString(),
		ConfirmationPersonalNote: senderSignature.ConfirmationPersonalNote.ValueString(),
	}
	res, err := r.client.EditSenderSignature(ctx, TypeStringToInt64(senderSignature.Id), body)
	if err != nil {
		clientDiag := diag.NewErrorDiagnostic("Client Error", fmt.Sprintf("Unable to update sender signature %s, got error: %s\nRequest Body:\n%#v", senderSignature.Id.ValueString(), err, body))
		return diag.Diagnostics{clientDiag}
	}

	return mapSenderSignatureResourceFromAPI(ctx, senderSignature, res)
}

func (r *senderSignatureResource) deleteFromAPI(ctx context.Context, sender *resource_sender_signature.SenderSignatureModel) diag.Diagnostics {
	err := r.client.DeleteSenderSignature(ctx, TypeStringToInt64(sender.Id))
	if err != nil {
		clientDiag := diag.NewErrorDiagnostic("Client Error", fmt.Sprintf("Unable to delete sender signature %s, got error: %s", sender.Id.ValueString(), err))
		return diag.Diagnostics{clientDiag}
	}

	return nil
}

func mapSenderSignatureResourceFromAPI(_ context.Context, sender *resource_sender_signature.SenderSignatureModel, res postmark.SenderSignatureDetails) diag.Diagnostics {
	sender.Id = types.StringValue(strconv.Itoa(int(res.ID)))
	sender.Domain = types.StringValue(res.Domain)
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
