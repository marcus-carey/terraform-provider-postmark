package provider

import (
	"context"
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

	// Example data value setting
	data.Id = types.StringValue("example-id")

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

func (r *domainResource) readFromAPI(_ context.Context, _ *resource_domain.DomainModel) diag.Diagnostics {
	return diag.Diagnostics{diag.NewErrorDiagnostic("Not Implemented", "This function is not implemented yet")}
	/*	res, err := r.client.GetDomain(ctx, domain.Id.ValueString())
		if err != nil {
			clientDiag := diag.NewErrorDiagnostic("Client Error", fmt.Sprintf("Unable to read domain, got error: %s", err))
			return diag.Diagnostics{clientDiag}
		}

		return mapDomainResourceFromAPI(ctx, domain, res)*/
}

func (r *domainResource) createFromAPI(_ context.Context, _ *resource_domain.DomainModel) diag.Diagnostics {
	return diag.Diagnostics{diag.NewErrorDiagnostic("Not Implemented", "This function is not implemented yet")}
	/*	body := mapDomainResourceToAPI(domain)
		res, err := r.client.CreateDomain(ctx, body)
		if err != nil {
			clientDiag := diag.NewErrorDiagnostic("Client Error", fmt.Sprintf("Unable to create domain, got error: %s", err))
			return diag.Diagnostics{clientDiag}
		}

		if res.ID == 0 {
			clientDiag := diag.NewErrorDiagnostic("Client Error", "Unable to create domain, got error: Domain ID is 0.")
			return diag.Diagnostics{clientDiag}
		}

		return mapDomainResourceFromAPI(ctx, domain, res)*/
}

func (r *domainResource) updateFromAPI(_ context.Context, _ *resource_domain.DomainModel) diag.Diagnostics {
	return diag.Diagnostics{diag.NewErrorDiagnostic("Not Implemented", "This function is not implemented yet")}
	/*	id := domain.Id.ValueString()
		body := mapDomainResourceToAPI(domain)
		body.ID, _ = strconv.ParseInt(id, 10, 64)
		res, err := r.client.EditDomain(ctx, id, body)
		if err != nil {
			clientDiag := diag.NewErrorDiagnostic("Client Error", fmt.Sprintf("Unable to update domain %s, got error: %s\nRequest Body:\n%#v", id, err, body))
			return diag.Diagnostics{clientDiag}
		}

		return mapDomainResourceFromAPI(ctx, domain, res)*/
}

func (r *domainResource) deleteFromAPI(_ context.Context, _ *resource_domain.DomainModel) diag.Diagnostics {
	return diag.Diagnostics{diag.NewErrorDiagnostic("Not Implemented", "This function is not implemented yet")}
}

/*func mapDomainResourceToAPI(domain *resource_domain.DomainModel) postmark.Domain {
	//return postmark.Domain{
	//	Name:                       domain.Name.ValueString(),
	//}
}

func mapDomainResourceFromAPI(ctx context.Context, domain *resource_domain.DomainModel, res postmark.Domain) diag.Diagnostics {
	return nil
}*/
