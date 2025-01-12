package provider

import (
	"context"
	"fmt"
	"strconv"
	"terraform-provider-postmark/internal/provider/datasource_server"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/mrz1836/postmark"
)

var _ datasource.DataSource = (*serverDataSource)(nil)

func NewServerDataSource() datasource.DataSource {
	return &serverDataSource{}
}

type serverDataSource struct {
	client *postmark.Client
}

func (d *serverDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_server"
}

func (d *serverDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	d.client = GetDataSourcePostmarkClient(req, resp)
}

func (d *serverDataSource) Schema(ctx context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = datasource_server.ServerDataSourceSchema(ctx)
}

func (d *serverDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data datasource_server.ServerModel

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

func (d *serverDataSource) readFromAPI(ctx context.Context, server *datasource_server.ServerModel) diag.Diagnostics {
	name := server.Name.ValueString()
	if server.Id.IsNull() && name == "" {
		return diag.Diagnostics{diag.NewErrorDiagnostic("Invalid Configuration", "Either the server ID or server name must be set")}
	}

	res := postmark.Server{}

	if !server.Id.IsNull() {

		server, err := d.client.GetServer(ctx, TypeStringToInt64(server.Id))
		if err != nil {
			clientDiag := diag.NewErrorDiagnostic("Client Error", fmt.Sprintf("Unable to read server, got error: %s", err))
			return diag.Diagnostics{clientDiag}
		}

		res = server
	} else {
		search, err := d.client.GetServers(ctx, 20, 0, name)
		if err != nil || len(search.Servers) == 0 {
			clientDiag := diag.NewErrorDiagnostic("Client Error", fmt.Sprintf("Unable to read server, got error: %s", err))
			return diag.Diagnostics{clientDiag}
		}

		// Find server with name that is exact match
		for _, s := range search.Servers {
			if s.Name == name {
				res = s
				break
			}
		}

		if res.ID == 0 {
			clientDiag := diag.NewErrorDiagnostic("Client Error", fmt.Sprintf("Unable to read server, no server found with name: %s", name))
			return diag.Diagnostics{clientDiag}
		}
	}

	server.Id = types.StringValue(strconv.FormatInt(res.ID, 10))
	server.Name = types.StringValue(res.Name)
	server.Color = types.StringValue(res.Color)

	apiTokens, parseDiags := parseListType(ctx, server.ApiTokens, res.APITokens)
	if parseDiags.HasError() {
		return parseDiags
	}
	server.ApiTokens = apiTokens

	server.DeliveryType = types.StringValue(res.DeliveryType)
	server.InboundAddress = types.StringValue(res.InboundAddress)
	server.InboundDomain = types.StringValue(res.InboundDomain)
	server.InboundHash = types.StringValue(res.InboundHash)
	server.InboundHookUrl = types.StringValue(res.InboundHookURL)
	server.InboundSpamThreshold = types.Int64Value(res.InboundSpamThreshold)
	server.PostFirstOpenOnly = types.BoolValue(res.PostFirstOpenOnly)
	server.RawEmailEnabled = types.BoolValue(res.RawEmailEnabled)
	server.ServerLink = types.StringValue(res.ServerLink)
	server.SmtpApiActivated = types.BoolValue(res.SMTPAPIActivated)
	server.IncludeBounceContentInHook = types.BoolValue(res.IncludeBounceContentInHook)
	server.EnableSmtpApiErrorHooks = types.BoolValue(res.EnableSMTPAPIErrorHooks)

	return nil
}
