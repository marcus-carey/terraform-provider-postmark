package provider

import (
	"context"
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

	// Example data value setting
	data.Id = types.StringValue("example-id")

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

func (r *senderSignatureResource) readFromAPI(_ context.Context, _ *resource_sender_signature.SenderSignatureModel) diag.Diagnostics {
	return diag.Diagnostics{diag.NewErrorDiagnostic("Not Implemented", "This function is not implemented yet")}
	/*res, err := r.client.GetSenderSignature(ctx, senderSignature.Id.ValueString())
	if err != nil {
		clientDiag := diag.NewErrorDiagnostic("Client Error", fmt.Sprintf("Unable to read senderSignature, got error: %s", err))
		return diag.Diagnostics{clientDiag}
	}

	return mapSenderSignatureResourceFromAPI(ctx, senderSignature, res)*/
}

func (r *senderSignatureResource) createFromAPI(_ context.Context, _ *resource_sender_signature.SenderSignatureModel) diag.Diagnostics {
	return diag.Diagnostics{diag.NewErrorDiagnostic("Not Implemented", "This function is not implemented yet")}
	/*body := mapSenderSignatureResourceToAPI(senderSignature)
	res, err := r.client.CreateSenderSignature(ctx, body)
	if err != nil {
		clientDiag := diag.NewErrorDiagnostic("Client Error", fmt.Sprintf("Unable to create senderSignature, got error: %s", err))
		return diag.Diagnostics{clientDiag}
	}

	if res.ID == 0 {
		clientDiag := diag.NewErrorDiagnostic("Client Error", "Unable to create senderSignature, got error: SenderSignature ID is 0.")
		return diag.Diagnostics{clientDiag}
	}

	return mapSenderSignatureResourceFromAPI(ctx, senderSignature, res)*/
}

func (r *senderSignatureResource) updateFromAPI(_ context.Context, _ *resource_sender_signature.SenderSignatureModel) diag.Diagnostics {
	return diag.Diagnostics{diag.NewErrorDiagnostic("Not Implemented", "This function is not implemented yet")}
	/*id := senderSignature.Id.ValueString()
	body := mapSenderSignatureResourceToAPI(senderSignature)
	body.ID, _ = strconv.ParseInt(id, 10, 64)
	res, err := r.client.EditSenderSignature(ctx, id, body)
	if err != nil {
		clientDiag := diag.NewErrorDiagnostic("Client Error", fmt.Sprintf("Unable to update senderSignature %s, got error: %s\nRequest Body:\n%#v", id, err, body))
		return diag.Diagnostics{clientDiag}
	}

	return mapSenderSignatureResourceFromAPI(ctx, senderSignature, res)*/
}

func (r *senderSignatureResource) deleteFromAPI(_ context.Context, _ *resource_sender_signature.SenderSignatureModel) diag.Diagnostics {
	return diag.Diagnostics{diag.NewErrorDiagnostic("Not Implemented", "This function is not implemented yet")}
}

/*func mapSenderSignatureResourceToAPI(senderSignature *resource_sender_signature.SenderSignatureModel) postmark.SenderSignature {
	//return postmark.SenderSignature{
	//	Name:                       senderSignature.Name.ValueString(),
	//}
}

func mapSenderSignatureResourceFromAPI(ctx context.Context, senderSignature *resource_sender_signature.SenderSignatureModel, res postmark.SenderSignature) diag.Diagnostics {
	return nil
}*/
