package provider

import (
	"context"
	"fmt"
	"strconv"
	"terraform-provider-postmark/internal/provider/datasource_webhook"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
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

func (d *webhookDataSource) readFromAPI(ctx context.Context, webhook *datasource_webhook.WebhookModel) diag.Diagnostics {
	d.client.ServerToken = webhook.ServerApiToken.ValueString()
	webhooks, err := d.client.ListWebhooks(ctx, webhook.MessageStream.ValueString())
	if err != nil || len(webhooks) == 0 {
		clientDiag := diag.NewErrorDiagnostic("Client Error", fmt.Sprintf("Unable to read webhook, got error: %s", err))
		return diag.Diagnostics{clientDiag}
	}

	res := webhooks[0]

	webhook.Id = types.StringValue(strconv.Itoa(res.ID))
	webhook.Url = types.StringValue(res.URL)
	webhook.MessageStream = types.StringValue(res.MessageStream)

	if res.HTTPAuth != nil && res.HTTPAuth.Username != "" && res.HTTPAuth.Password != "" {
		httpAuth, _ := datasource_webhook.NewHttpAuthValue(
			datasource_webhook.HttpAuthValue{}.AttributeTypes(ctx),
			map[string]attr.Value{
				"username": types.StringValue(res.HTTPAuth.Username),
				"password": types.StringValue(res.HTTPAuth.Password),
			})

		webhook.HttpAuth = httpAuth
	} else {
		webhook.HttpAuth = datasource_webhook.NewHttpAuthValueNull()
	}

	httpHeaderValues := make([]datasource_webhook.HttpHeadersValue, 0)
	for _, header := range res.HTTPHeaders {
		headerValue, _ := datasource_webhook.NewHttpHeadersValue(
			datasource_webhook.HttpHeadersValue{}.AttributeTypes(ctx),
			map[string]attr.Value{
				"name":  types.StringValue(header.Name),
				"value": types.StringValue(header.Value),
			})

		httpHeaderValues = append(httpHeaderValues, headerValue)
	}
	webhook.HttpHeaders, _ = types.ListValueFrom(ctx, webhook.HttpHeaders.ElementType(ctx), httpHeaderValues)

	webhook.OpenTrigger, _ = datasource_webhook.NewOpenTriggerValue(
		datasource_webhook.OpenTriggerValue{}.AttributeTypes(ctx),
		map[string]attr.Value{
			"enabled":              types.BoolValue(res.Triggers.Open.Enabled),
			"post_first_open_only": types.BoolValue(res.Triggers.Open.PostFirstOpenOnly),
		})

	webhook.ClickTrigger, _ = datasource_webhook.NewClickTriggerValue(
		datasource_webhook.ClickTriggerValue{}.AttributeTypes(ctx),
		map[string]attr.Value{
			"enabled": types.BoolValue(res.Triggers.Click.Enabled),
		})

	webhook.DeliveryTrigger, _ = datasource_webhook.NewDeliveryTriggerValue(
		datasource_webhook.DeliveryTriggerValue{}.AttributeTypes(ctx),
		map[string]attr.Value{
			"enabled": types.BoolValue(res.Triggers.Delivery.Enabled),
		})

	webhook.BounceTrigger, _ = datasource_webhook.NewBounceTriggerValue(
		datasource_webhook.BounceTriggerValue{}.AttributeTypes(ctx),
		map[string]attr.Value{
			"enabled":         types.BoolValue(res.Triggers.Bounce.Enabled),
			"include_content": types.BoolValue(res.Triggers.Bounce.IncludeContent),
		})

	webhook.SpamComplaintTrigger, _ = datasource_webhook.NewSpamComplaintTriggerValue(
		datasource_webhook.SpamComplaintTriggerValue{}.AttributeTypes(ctx),
		map[string]attr.Value{
			"enabled":         types.BoolValue(res.Triggers.SpamComplaint.Enabled),
			"include_content": types.BoolValue(res.Triggers.SpamComplaint.IncludeContent),
		})

	webhook.SubscriptionChangeTrigger, _ = datasource_webhook.NewSubscriptionChangeTriggerValue(
		datasource_webhook.SubscriptionChangeTriggerValue{}.AttributeTypes(ctx),
		map[string]attr.Value{
			"enabled": types.BoolValue(res.Triggers.SubscriptionChange.Enabled),
		})

	return nil
}
