// Code generated by terraform-plugin-framework-generator DO NOT EDIT.

package datasource_domain

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func DomainDataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"dkim_host": schema.StringAttribute{
				Computed:            true,
				Description:         "DNS TXT host being used to validate messages sent in.",
				MarkdownDescription: "DNS TXT host being used to validate messages sent in.",
			},
			"dkim_pending_host": schema.StringAttribute{
				Computed:            true,
				Description:         "If a DKIM rotation has been initiated or this DKIM is from a new Domain, this field will show the pending DKIM DNS TXT host which has yet to be setup and confirmed at your registrar or DNS host.",
				MarkdownDescription: "If a DKIM rotation has been initiated or this DKIM is from a new Domain, this field will show the pending DKIM DNS TXT host which has yet to be setup and confirmed at your registrar or DNS host.",
			},
			"dkim_pending_text_value": schema.StringAttribute{
				Computed:            true,
				Description:         "Similar to the DKIMPendingHost field, this will show the DNS TXT value waiting to be confirmed at your registrar or DNS host.",
				MarkdownDescription: "Similar to the DKIMPendingHost field, this will show the DNS TXT value waiting to be confirmed at your registrar or DNS host.",
			},
			"dkim_revoked_host": schema.StringAttribute{
				Computed:            true,
				Description:         "Once a new DKIM has been confirmed at your registrar or DNS host, Postmark will revoke the old DKIM host in preparation for removing it permantly from the system.",
				MarkdownDescription: "Once a new DKIM has been confirmed at your registrar or DNS host, Postmark will revoke the old DKIM host in preparation for removing it permantly from the system.",
			},
			"dkim_revoked_text_value": schema.StringAttribute{
				Computed:            true,
				Description:         "Similar to DKIMRevokedHost, this field will show the DNS TXT value that will soon be removed from the Postmark system.",
				MarkdownDescription: "Similar to DKIMRevokedHost, this field will show the DNS TXT value that will soon be removed from the Postmark system.",
			},
			"dkim_text_value": schema.StringAttribute{
				Computed:            true,
				Description:         "DNS TXT value being used to validate messages sent in.",
				MarkdownDescription: "DNS TXT value being used to validate messages sent in.",
			},
			"dkim_update_status": schema.StringAttribute{
				Computed:            true,
				Description:         "While DKIM renewal or new DKIM operations are being conducted or setup, this field will indicate Pending. After all DNS TXT records are up to date and any pending renewal operations are finished, it will indicate Verified.",
				MarkdownDescription: "While DKIM renewal or new DKIM operations are being conducted or setup, this field will indicate Pending. After all DNS TXT records are up to date and any pending renewal operations are finished, it will indicate Verified.",
			},
			"dkim_verified": schema.BoolAttribute{
				Computed:            true,
				Description:         "Specifies whether DKIM has ever been verified for the domain or not. Once DKIM is verified, this response will stay true, even if the record is later removed from DNS.",
				MarkdownDescription: "Specifies whether DKIM has ever been verified for the domain or not. Once DKIM is verified, this response will stay true, even if the record is later removed from DNS.",
			},
			"id": schema.StringAttribute{
				Required:            true,
				Description:         "Unique ID of the Domain.",
				MarkdownDescription: "Unique ID of the Domain.",
			},
			"name": schema.StringAttribute{
				Computed:            true,
				Description:         "Domain Name",
				MarkdownDescription: "Domain Name",
			},
			"return_path_domain": schema.StringAttribute{
				Computed:            true,
				Description:         "The custom Return-Path for this domain, please read our support page.",
				MarkdownDescription: "The custom Return-Path for this domain, please read our support page.",
			},
			"return_path_domain_cname_value": schema.StringAttribute{
				Computed:            true,
				Description:         "The CNAME DNS record that Postmark expects to find at the ReturnPathDomain value.",
				MarkdownDescription: "The CNAME DNS record that Postmark expects to find at the ReturnPathDomain value.",
			},
			"return_path_domain_verified": schema.BoolAttribute{
				Computed:            true,
				Description:         "The verification state of the Return-Path domain. Tells you if the Return-Path is actively being used or still needs further action to be used.",
				MarkdownDescription: "The verification state of the Return-Path domain. Tells you if the Return-Path is actively being used or still needs further action to be used.",
			},
			"safe_to_remove_revoked_key_from_dns": schema.BoolAttribute{
				Computed:            true,
				Description:         "Indicates whether you may safely delete the old DKIM DNS TXT records at your registrar or DNS host. The new DKIM is now safely in use.",
				MarkdownDescription: "Indicates whether you may safely delete the old DKIM DNS TXT records at your registrar or DNS host. The new DKIM is now safely in use.",
			},
			"spf_host": schema.StringAttribute{
				Computed:            true,
				Description:         "Host name used for the SPF configuration.",
				MarkdownDescription: "Host name used for the SPF configuration.",
			},
			"spf_text_value": schema.StringAttribute{
				Computed:            true,
				Description:         "Value that can be optionally setup with your DNS host.",
				MarkdownDescription: "Value that can be optionally setup with your DNS host.",
			},
			"weak_dkim": schema.BoolAttribute{
				Computed:            true,
				Description:         "DKIM is using a strength weaker than 1024 bit. If so, it’s possible to request a new DKIM using the RequestNewDKIM function.",
				MarkdownDescription: "DKIM is using a strength weaker than 1024 bit. If so, it’s possible to request a new DKIM using the RequestNewDKIM function.",
			},
		},
	}
}

type DomainModel struct {
	DkimHost                      types.String `tfsdk:"dkim_host"`
	DkimPendingHost               types.String `tfsdk:"dkim_pending_host"`
	DkimPendingTextValue          types.String `tfsdk:"dkim_pending_text_value"`
	DkimRevokedHost               types.String `tfsdk:"dkim_revoked_host"`
	DkimRevokedTextValue          types.String `tfsdk:"dkim_revoked_text_value"`
	DkimTextValue                 types.String `tfsdk:"dkim_text_value"`
	DkimUpdateStatus              types.String `tfsdk:"dkim_update_status"`
	DkimVerified                  types.Bool   `tfsdk:"dkim_verified"`
	Id                            types.String `tfsdk:"id"`
	Name                          types.String `tfsdk:"name"`
	ReturnPathDomain              types.String `tfsdk:"return_path_domain"`
	ReturnPathDomainCnameValue    types.String `tfsdk:"return_path_domain_cname_value"`
	ReturnPathDomainVerified      types.Bool   `tfsdk:"return_path_domain_verified"`
	SafeToRemoveRevokedKeyFromDns types.Bool   `tfsdk:"safe_to_remove_revoked_key_from_dns"`
	SpfHost                       types.String `tfsdk:"spf_host"`
	SpfTextValue                  types.String `tfsdk:"spf_text_value"`
	WeakDkim                      types.Bool   `tfsdk:"weak_dkim"`
}
