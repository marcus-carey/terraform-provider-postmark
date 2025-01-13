package provider

import (
	"context"
	"fmt"
	"strconv"
	"terraform-provider-postmark/internal/provider/datasource_domain"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/mrz1836/postmark"
)

var _ datasource.DataSource = (*domainDataSource)(nil)

func NewDomainDataSource() datasource.DataSource {
	return &domainDataSource{}
}

type domainDataSource struct {
	client *postmark.Client
}

func (d *domainDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_domain"
}

func (d *domainDataSource) Schema(ctx context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = datasource_domain.DomainDataSourceSchema(ctx)
}

func (d *domainDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	d.client = GetDataSourcePostmarkClient(req, resp)
}

func (d *domainDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data datasource_domain.DomainModel

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

func (d *domainDataSource) readFromAPI(ctx context.Context, domain *datasource_domain.DomainModel) diag.Diagnostics {
	res, err := d.client.GetDomain(ctx, TypeStringToInt64(domain.Id))
	if err != nil {
		clientDiag := diag.NewErrorDiagnostic("Client Error", fmt.Sprintf("Unable to read domain, got error: %s", err))
		return diag.Diagnostics{clientDiag}
	}

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
