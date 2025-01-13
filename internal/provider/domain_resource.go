package provider

import (
	"context"
	"fmt"
	"strconv"
	"terraform-provider-postmark/internal/provider/resource_domain"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/mrz1836/postmark"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ resource.Resource                = &domainResource{}
	_ resource.ResourceWithConfigure   = &domainResource{}
	_ resource.ResourceWithImportState = &domainResource{}
)

func NewDomainResource() resource.Resource {
	return &domainResource{}
}

type domainResource struct {
	client *postmark.Client
}

func (r *domainResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_domain"
}

func (r *domainResource) Schema(ctx context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = resource_domain.DomainResourceSchema(ctx)
}

func (r *domainResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	r.client = GetResourcePostmarkClient(req, resp)
}

func (r *domainResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	// Retrieve import ID and save to id attribute
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *domainResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data resource_domain.DomainModel

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

func (r *domainResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data resource_domain.DomainModel

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

func (r *domainResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data resource_domain.DomainModel

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

func (r *domainResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data resource_domain.DomainModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Delete API call logic
	resp.Diagnostics.Append(r.deleteFromAPI(ctx, &data)...)
}

func (r *domainResource) readFromAPI(ctx context.Context, domain *resource_domain.DomainModel) diag.Diagnostics {
	res, err := r.client.GetDomain(ctx, TypeStringToInt64(domain.Id))
	if err != nil {
		clientDiag := diag.NewErrorDiagnostic("Client Error", fmt.Sprintf("Unable to read domain, got error: %s", err))
		return diag.Diagnostics{clientDiag}
	}

	return mapDomainResourceFromAPI(ctx, domain, res)
}

func (r *domainResource) createFromAPI(ctx context.Context, domain *resource_domain.DomainModel) diag.Diagnostics {
	body := postmark.DomainCreateRequest{
		Name:             domain.Name.ValueString(),
		ReturnPathDomain: domain.ReturnPathDomain.ValueString(),
	}
	res, err := r.client.CreateDomain(ctx, body)
	if err != nil {
		clientDiag := diag.NewErrorDiagnostic("Client Error", fmt.Sprintf("Unable to create domain, got error: %s", err))
		return diag.Diagnostics{clientDiag}
	}

	if res.ID == 0 {
		clientDiag := diag.NewErrorDiagnostic("Client Error", "Unable to create domain, got error: Domain ID is 0.")
		return diag.Diagnostics{clientDiag}
	}

	return mapDomainResourceFromAPI(ctx, domain, res)
}

func (r *domainResource) updateFromAPI(ctx context.Context, domain *resource_domain.DomainModel) diag.Diagnostics {
	body := postmark.DomainEditRequest{
		ReturnPathDomain: domain.ReturnPathDomain.ValueString(),
	}
	res, err := r.client.EditDomain(ctx, TypeStringToInt64(domain.Id), body)
	if err != nil {
		clientDiag := diag.NewErrorDiagnostic("Client Error", fmt.Sprintf("Unable to update domain %s, got error: %s", domain.Id.ValueString(), err))
		return diag.Diagnostics{clientDiag}
	}

	return mapDomainResourceFromAPI(ctx, domain, res)
}

func (r *domainResource) deleteFromAPI(ctx context.Context, domain *resource_domain.DomainModel) diag.Diagnostics {
	err := r.client.DeleteDomain(ctx, TypeStringToInt64(domain.Id))
	if err != nil {
		clientDiag := diag.NewErrorDiagnostic("Client Error", fmt.Sprintf("Unable to delete domain %s, got error: %s", domain.Id.ValueString(), err))
		return diag.Diagnostics{clientDiag}
	}

	return nil
}

func mapDomainResourceFromAPI(_ context.Context, domain *resource_domain.DomainModel, res postmark.DomainDetails) diag.Diagnostics {
	domain.Id = types.StringValue(strconv.FormatInt(res.ID, 10))
	domain.Name = types.StringValue(res.Name)
	domain.SpfHost = types.StringValue(res.SPFHost)
	domain.SpfTextValue = types.StringValue(res.SPFTextValue)
	domain.DkimVerified = types.BoolValue(res.DKIMVerified)
	domain.WeakDkim = types.BoolValue(res.WeakDKIM)
	domain.DkimHost = types.StringValue(res.DKIMHost)
	domain.DkimTextValue = types.StringValue(res.DKIMTextValue)
	domain.DkimPendingHost = types.StringValue(res.DKIMPendingHost)
	domain.DkimPendingTextValue = types.StringValue(res.DKIMPendingTextValue)
	domain.DkimRevokedHost = types.StringValue(res.DKIMRevokedHost)
	domain.DkimRevokedTextValue = types.StringValue(res.DKIMRevokedTextValue)
	domain.SafeToRemoveRevokedKeyFromDns = types.BoolValue(res.SafeToRemoveRevokedKeyFromDNS)
	domain.DkimUpdateStatus = types.StringValue(res.DKIMUpdateStatus)
	domain.ReturnPathDomain = types.StringValue(res.ReturnPathDomain)
	domain.ReturnPathDomainVerified = types.BoolValue(res.ReturnPathDomainVerified)
	domain.ReturnPathDomainCnameValue = types.StringValue(res.ReturnPathDomainCNAMEValue)

	return nil
}
