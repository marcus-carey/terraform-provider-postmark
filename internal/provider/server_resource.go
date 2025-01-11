package provider

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/mrz1836/postmark"
	"strconv"
	"terraform-provider-postmark/internal/provider/resource_server"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ resource.Resource                = &serverResource{}
	_ resource.ResourceWithConfigure   = &serverResource{}
	_ resource.ResourceWithImportState = &serverResource{}
)

func NewServerResource() resource.Resource {
	return &serverResource{}
}

type serverResource struct {
	client *postmark.Client
}

func (r *serverResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_server"
}

func (r *serverResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = resource_server.ServerResourceSchema(ctx)
}

func (r *serverResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*postmark.Client)

	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Resource Configure Type",
			fmt.Sprintf("Expected *http.Client, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	r.client = client
}

func (r *serverResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	// Retrieve import ID and save to id attribute
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *serverResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data resource_server.ServerModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Create API call logic
	resp.Diagnostics.Append(r.createFromApi(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *serverResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data resource_server.ServerModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Read API call logic
	resp.Diagnostics.Append(r.readFromApi(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *serverResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data resource_server.ServerModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Update API call logic
	resp.Diagnostics.Append(r.updateFromApi(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *serverResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data resource_server.ServerModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Delete API call logic
	resp.Diagnostics.Append(r.deleteFromApi(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *serverResource) readFromApi(ctx context.Context, server *resource_server.ServerModel) diag.Diagnostics {
	res, err := r.client.GetServer(context.Background(), server.Id.ValueString())

	if err != nil {
		clientDiag := diag.NewErrorDiagnostic("Client Error", fmt.Sprintf("Unable to read server, got error: %s", err))
		return diag.Diagnostics{clientDiag}
	}

	return mapServerResourceFromApi(ctx, server, res)
}

func (r *serverResource) createFromApi(ctx context.Context, server *resource_server.ServerModel) diag.Diagnostics {
	body := mapServerResourceToApi(ctx, server)
	res, err := r.client.CreateServer(context.Background(), body)

	if err != nil {
		clientDiag := diag.NewErrorDiagnostic("Client Error", fmt.Sprintf("Unable to create server, got error: %s", err))
		return diag.Diagnostics{clientDiag}
	}

	if res.ID == 0 {
		clientDiag := diag.NewErrorDiagnostic("Client Error", "Unable to create server, got error: Server ID is 0.")
		return diag.Diagnostics{clientDiag}
	}

	return mapServerResourceFromApi(ctx, server, res)
}

func (r *serverResource) updateFromApi(ctx context.Context, server *resource_server.ServerModel) diag.Diagnostics {
	id := server.Id.ValueString()
	body := mapServerResourceToApi(ctx, server)
	body.ID, _ = strconv.ParseInt(id, 10, 64)
	res, err := r.client.EditServer(context.Background(), id, body)

	if err != nil {
		clientDiag := diag.NewErrorDiagnostic("Client Error", fmt.Sprintf("Unable to update server %s, got error: %s\nRequest Body:\n%#v", id, err, body))
		return diag.Diagnostics{clientDiag}
	}

	return mapServerResourceFromApi(ctx, server, res)
}

func (r *serverResource) deleteFromApi(ctx context.Context, server *resource_server.ServerModel) diag.Diagnostics {
	// TODO - Implement server deletion, but catch error, since not all servers can be deleted via the api
	clientDiag := diag.NewWarningDiagnostic("Server Deletion Not Supported", "Server must be deleted manually in the Postmark UI.")
	return diag.Diagnostics{clientDiag}
}

func mapServerResourceToApi(ctx context.Context, server *resource_server.ServerModel) postmark.Server {
	return postmark.Server{
		Name:                       server.Name.ValueString(),
		Color:                      server.Color.ValueString(),
		SMTPAPIActivated:           server.SmtpApiActivated.ValueBool(),
		RawEmailEnabled:            server.RawEmailEnabled.ValueBool(),
		DeliveryType:               server.DeliveryType.ValueString(),
		InboundHookURL:             server.InboundHookUrl.ValueString(),
		PostFirstOpenOnly:          server.PostFirstOpenOnly.ValueBool(),
		InboundDomain:              server.InboundDomain.ValueString(),
		InboundSpamThreshold:       server.InboundSpamThreshold.ValueInt64(),
		TrackOpens:                 server.TrackOpens.ValueBool(),
		TrackLinks:                 server.TrackLinks.ValueString(),
		IncludeBounceContentInHook: server.IncludeBounceContentInHook.ValueBool(),
		EnableSMTPAPIErrorHooks:    server.EnableSmtpApiErrorHooks.ValueBool(),
	}
}

func mapServerResourceFromApi(ctx context.Context, server *resource_server.ServerModel, res postmark.Server) diag.Diagnostics {
	server.Id = types.StringValue(strconv.FormatInt(res.ID, 10))
	server.Name = types.StringValue(res.Name)

	apiTokenDiags := server.ApiTokens.ElementsAs(ctx, &res.APITokens, false)
	server.ApiTokens, apiTokenDiags = types.ListValueFrom(ctx, server.ApiTokens.ElementType(ctx), res.APITokens)

	if apiTokenDiags.HasError() {
		return apiTokenDiags
	}

	server.Color = types.StringValue(res.Color)
	server.SmtpApiActivated = types.BoolValue(res.SMTPAPIActivated)
	server.RawEmailEnabled = types.BoolValue(res.RawEmailEnabled)
	server.DeliveryType = types.StringValue(res.DeliveryType)
	server.ServerLink = types.StringValue(res.ServerLink)
	server.InboundAddress = types.StringValue(res.InboundAddress)
	server.InboundHookUrl = types.StringValue(res.InboundHookURL)
	server.PostFirstOpenOnly = types.BoolValue(res.PostFirstOpenOnly)
	server.InboundDomain = types.StringValue(res.InboundDomain)
	server.InboundHash = types.StringValue(res.InboundHash)
	server.InboundSpamThreshold = types.Int64Value(res.InboundSpamThreshold)
	server.TrackOpens = types.BoolValue(res.TrackOpens)
	server.TrackLinks = types.StringValue(res.TrackLinks)
	server.IncludeBounceContentInHook = types.BoolValue(res.IncludeBounceContentInHook)
	server.EnableSmtpApiErrorHooks = types.BoolValue(res.EnableSMTPAPIErrorHooks)

	return nil
}
