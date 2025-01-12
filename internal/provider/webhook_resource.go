package provider

import (
	"context"
	"fmt"
	"terraform-provider-postmark/internal/provider/resource_webhook"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/mrz1836/postmark"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ resource.Resource                = &webhookResource{}
	_ resource.ResourceWithConfigure   = &webhookResource{}
	_ resource.ResourceWithImportState = &webhookResource{}
)

func NewWebhookResource() resource.Resource {
	return &webhookResource{}
}

type webhookResource struct {
	client *postmark.Client
}

func (r *webhookResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_webhook"
}

func (r *webhookResource) Schema(ctx context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = resource_webhook.WebhookResourceSchema(ctx)
}

func (r *webhookResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	r.client = GetResourcePostmarkClient(req, resp)
}

func (r *webhookResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	// Retrieve import ID and save to id attribute
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *webhookResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data resource_webhook.WebhookModel

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

func (r *webhookResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data resource_webhook.WebhookModel

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

func (r *webhookResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data resource_webhook.WebhookModel

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

func (r *webhookResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data resource_webhook.WebhookModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Delete API call logic
	resp.Diagnostics.Append(r.deleteFromAPI(ctx, &data)...)
}

func (r *webhookResource) readFromAPI(ctx context.Context, webhook *resource_webhook.WebhookModel) diag.Diagnostics {
	r.client.ServerToken = webhook.ServerApiToken.ValueString()
	res, err := r.client.GetWebhook(ctx, TypeStringToInt(webhook.Id))
	if err != nil {
		clientDiag := diag.NewErrorDiagnostic("Client Error", fmt.Sprintf("Unable to read webhook, got error: %s", err))
		return diag.Diagnostics{clientDiag}
	}

	return mapWebhookResourceFromAPI(ctx, webhook, res)
}

func (r *webhookResource) createFromAPI(ctx context.Context, webhook *resource_webhook.WebhookModel) diag.Diagnostics {
	r.client.ServerToken = webhook.ServerApiToken.ValueString()
	body := mapWebhookResourceToAPI(webhook)
	res, err := r.client.CreateWebhook(ctx, body)
	if err != nil {
		clientDiag := diag.NewErrorDiagnostic("Client Error", fmt.Sprintf("Unable to create webhook, got error: %s", err))
		return diag.Diagnostics{clientDiag}
	}

	if res.ID == 0 {
		clientDiag := diag.NewErrorDiagnostic("Client Error", "Unable to create webhook, got error: Webhook ID is 0.")
		return diag.Diagnostics{clientDiag}
	}

	return mapWebhookResourceFromAPI(ctx, webhook, res)
}

func (r *webhookResource) updateFromAPI(ctx context.Context, webhook *resource_webhook.WebhookModel) diag.Diagnostics {
	r.client.ServerToken = webhook.ServerApiToken.ValueString()
	id := TypeStringToInt(webhook.Id)
	body := mapWebhookResourceToAPI(webhook)
	body.ID = id
	res, err := r.client.EditWebhook(ctx, id, body)
	if err != nil {
		clientDiag := diag.NewErrorDiagnostic("Client Error", fmt.Sprintf("Unable to update webhook %d, got error: %s\nRequest Body:\n%#v", id, err, body))
		return diag.Diagnostics{clientDiag}
	}

	return mapWebhookResourceFromAPI(ctx, webhook, res)
}

func (r *webhookResource) deleteFromAPI(_ context.Context, webhook *resource_webhook.WebhookModel) diag.Diagnostics {
	r.client.ServerToken = webhook.ServerApiToken.ValueString()
	return diag.Diagnostics{diag.NewErrorDiagnostic("Not Implemented", "This function is not implemented yet")}
}

func mapWebhookResourceToAPI(_ *resource_webhook.WebhookModel) postmark.Webhook {
	return postmark.Webhook{
		// TODO - Map the webhook model to the API webhook
	}
}

func mapWebhookResourceFromAPI(_ context.Context, _ *resource_webhook.WebhookModel, _ postmark.Webhook) diag.Diagnostics {
	// TODO - Map the API webhook to the webhook model
	return nil
}
