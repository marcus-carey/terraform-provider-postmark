package provider

import (
	"context"
	"fmt"
	"strconv"
	"terraform-provider-postmark/internal/provider/resource_webhook"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
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
	body := r.mapResourceToAPI(ctx, webhook)
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
	body := r.mapResourceToAPI(ctx, webhook)
	res, err := r.client.EditWebhook(ctx, TypeStringToInt(webhook.Id), body)
	if err != nil {
		clientDiag := diag.NewErrorDiagnostic("Client Error", fmt.Sprintf("Unable to update webhook %s, got error: %s", webhook.Id, err))
		return diag.Diagnostics{clientDiag}
	}

	return mapWebhookResourceFromAPI(ctx, webhook, res)
}

func (r *webhookResource) deleteFromAPI(ctx context.Context, webhook *resource_webhook.WebhookModel) diag.Diagnostics {
	r.client.ServerToken = webhook.ServerApiToken.ValueString()
	err := r.client.DeleteWebhook(ctx, TypeStringToInt(webhook.Id))
	if err != nil {
		clientDiag := diag.NewErrorDiagnostic("Client Error", fmt.Sprintf("Unable to delete webhook %s, got error: %s", webhook.Id.ValueString(), err))
		return diag.Diagnostics{clientDiag}
	}

	return nil
}

func (r *webhookResource) mapResourceToAPI(ctx context.Context, webhook *resource_webhook.WebhookModel) postmark.Webhook {
	httpAuth := &postmark.WebhookHTTPAuth{}
	if !webhook.HttpAuth.IsNull() {
		httpAuth.Username = webhook.HttpAuth.Username.ValueString()
		httpAuth.Password = webhook.HttpAuth.Password.ValueString()
	}

	httpHeaders := make([]postmark.Header, 0)
	if !webhook.HttpHeaders.IsNull() {
		httpHeaderValues := make([]resource_webhook.HttpHeadersValue, 0)
		_ = webhook.HttpHeaders.ElementsAs(ctx, httpHeaderValues, false)

		for _, element := range webhook.HttpHeaders.Elements() {
			header := element.(resource_webhook.HttpHeadersValue)
			httpHeaders = append(httpHeaders, postmark.Header{
				Name:  header.Name.ValueString(),
				Value: header.Value.ValueString(),
			})
		}
	}

	return postmark.Webhook{
		URL:           webhook.Url.ValueString(),
		MessageStream: webhook.MessageStream.ValueString(),
		HTTPAuth:      httpAuth,
		HTTPHeaders:   httpHeaders,
		Triggers: postmark.WebhookTrigger{
			Open: postmark.WebhookTriggerOpen{
				WebhookTriggerEnabled: postmark.WebhookTriggerEnabled{
					Enabled: webhook.OpenTrigger.Enabled.ValueBool(),
				},
				PostFirstOpenOnly: webhook.OpenTrigger.PostFirstOpenOnly.ValueBool(),
			},
			Click: postmark.WebhookTriggerEnabled{
				Enabled: webhook.ClickTrigger.Enabled.ValueBool(),
			},
			Delivery: postmark.WebhookTriggerEnabled{
				Enabled: webhook.DeliveryTrigger.Enabled.ValueBool(),
			},
			Bounce: postmark.WebhookTriggerIncContent{
				WebhookTriggerEnabled: postmark.WebhookTriggerEnabled{
					Enabled: webhook.BounceTrigger.Enabled.ValueBool(),
				},
				IncludeContent: webhook.BounceTrigger.IncludeContent.ValueBool(),
			},
			SpamComplaint: postmark.WebhookTriggerIncContent{
				WebhookTriggerEnabled: postmark.WebhookTriggerEnabled{
					Enabled: webhook.SpamComplaintTrigger.Enabled.ValueBool(),
				},
				IncludeContent: webhook.SpamComplaintTrigger.IncludeContent.ValueBool(),
			},
			SubscriptionChange: postmark.WebhookTriggerEnabled{
				Enabled: webhook.SubscriptionChangeTrigger.Enabled.ValueBool(),
			},
		},
	}
}

func mapWebhookResourceFromAPI(ctx context.Context, webhook *resource_webhook.WebhookModel, res postmark.Webhook) diag.Diagnostics {
	webhook.Id = types.StringValue(strconv.Itoa(res.ID))
	webhook.Url = types.StringValue(res.URL)
	webhook.MessageStream = types.StringValue(res.MessageStream)

	if res.HTTPAuth != nil && res.HTTPAuth.Username != "" && res.HTTPAuth.Password != "" {
		httpAuth, _ := resource_webhook.NewHttpAuthValue(
			resource_webhook.HttpAuthValue{}.AttributeTypes(ctx),
			map[string]attr.Value{
				"username": types.StringValue(res.HTTPAuth.Username),
				"password": types.StringValue(res.HTTPAuth.Password),
			})

		webhook.HttpAuth = httpAuth
	} else {
		webhook.HttpAuth = resource_webhook.NewHttpAuthValueNull()
	}

	httpHeaderValues := make([]resource_webhook.HttpHeadersValue, 0)
	for _, header := range res.HTTPHeaders {
		headerValue, _ := resource_webhook.NewHttpHeadersValue(
			resource_webhook.HttpHeadersValue{}.AttributeTypes(ctx),
			map[string]attr.Value{
				"name":  types.StringValue(header.Name),
				"value": types.StringValue(header.Value),
			})

		httpHeaderValues = append(httpHeaderValues, headerValue)
	}
	webhook.HttpHeaders, _ = types.ListValueFrom(ctx, webhook.HttpHeaders.ElementType(ctx), httpHeaderValues)

	webhook.OpenTrigger, _ = resource_webhook.NewOpenTriggerValue(
		resource_webhook.OpenTriggerValue{}.AttributeTypes(ctx),
		map[string]attr.Value{
			"enabled":              types.BoolValue(res.Triggers.Open.Enabled),
			"post_first_open_only": types.BoolValue(res.Triggers.Open.PostFirstOpenOnly),
		})

	webhook.ClickTrigger, _ = resource_webhook.NewClickTriggerValue(
		resource_webhook.ClickTriggerValue{}.AttributeTypes(ctx),
		map[string]attr.Value{
			"enabled": types.BoolValue(res.Triggers.Click.Enabled),
		})

	webhook.DeliveryTrigger, _ = resource_webhook.NewDeliveryTriggerValue(
		resource_webhook.DeliveryTriggerValue{}.AttributeTypes(ctx),
		map[string]attr.Value{
			"enabled": types.BoolValue(res.Triggers.Delivery.Enabled),
		})

	webhook.BounceTrigger, _ = resource_webhook.NewBounceTriggerValue(
		resource_webhook.BounceTriggerValue{}.AttributeTypes(ctx),
		map[string]attr.Value{
			"enabled":         types.BoolValue(res.Triggers.Bounce.Enabled),
			"include_content": types.BoolValue(res.Triggers.Bounce.IncludeContent),
		})

	webhook.SpamComplaintTrigger, _ = resource_webhook.NewSpamComplaintTriggerValue(
		resource_webhook.SpamComplaintTriggerValue{}.AttributeTypes(ctx),
		map[string]attr.Value{
			"enabled":         types.BoolValue(res.Triggers.SpamComplaint.Enabled),
			"include_content": types.BoolValue(res.Triggers.SpamComplaint.IncludeContent),
		})

	webhook.SubscriptionChangeTrigger, _ = resource_webhook.NewSubscriptionChangeTriggerValue(
		resource_webhook.SubscriptionChangeTriggerValue{}.AttributeTypes(ctx),
		map[string]attr.Value{
			"enabled": types.BoolValue(res.Triggers.SubscriptionChange.Enabled),
		})

	return nil
}
