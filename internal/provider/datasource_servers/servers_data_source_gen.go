// Code generated by terraform-plugin-framework-generator DO NOT EDIT.

package datasource_servers

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func ServersDataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"name": schema.StringAttribute{
				Required:            true,
				Description:         "Filter by a specific server name",
				MarkdownDescription: "Filter by a specific server name",
			},
			"servers": schema.ListNestedAttribute{
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"api_tokens": schema.ListAttribute{
							ElementType: types.StringType,
							Computed:    true,
						},
						"color": schema.StringAttribute{
							Computed: true,
						},
						"id": schema.Int64Attribute{
							Computed: true,
						},
						"inbound_address": schema.StringAttribute{
							Computed: true,
						},
						"inbound_domain": schema.StringAttribute{
							Computed: true,
						},
						"inbound_hash": schema.StringAttribute{
							Computed: true,
						},
						"inbound_hook_url": schema.StringAttribute{
							Computed: true,
						},
						"inbound_spam_threshold": schema.Int64Attribute{
							Computed: true,
						},
						"name": schema.StringAttribute{
							Computed: true,
						},
						"post_first_open_only": schema.BoolAttribute{
							Computed: true,
						},
						"raw_email_enabled": schema.BoolAttribute{
							Computed: true,
						},
						"server_link": schema.StringAttribute{
							Computed: true,
						},
						"smtp_api_activated": schema.BoolAttribute{
							Computed: true,
						},
						"track_links": schema.StringAttribute{
							Computed: true,
						},
						"track_opens": schema.BoolAttribute{
							Computed: true,
						},
					},
					CustomType: ServersType{
						ObjectType: types.ObjectType{
							AttrTypes: ServersValue{}.AttributeTypes(ctx),
						},
					},
				},
				Computed: true,
			},
			"total_count": schema.Int64Attribute{
				Computed: true,
			},
		},
	}
}

type ServersModel struct {
	Name       types.String `tfsdk:"name"`
	Servers    types.List   `tfsdk:"servers"`
	TotalCount types.Int64  `tfsdk:"total_count"`
}

var _ basetypes.ObjectTypable = ServersType{}

type ServersType struct {
	basetypes.ObjectType
}

func (t ServersType) Equal(o attr.Type) bool {
	other, ok := o.(ServersType)

	if !ok {
		return false
	}

	return t.ObjectType.Equal(other.ObjectType)
}

func (t ServersType) String() string {
	return "ServersType"
}

func (t ServersType) ValueFromObject(ctx context.Context, in basetypes.ObjectValue) (basetypes.ObjectValuable, diag.Diagnostics) {
	var diags diag.Diagnostics

	attributes := in.Attributes()

	apiTokensAttribute, ok := attributes["api_tokens"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`api_tokens is missing from object`)

		return nil, diags
	}

	apiTokensVal, ok := apiTokensAttribute.(basetypes.ListValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`api_tokens expected to be basetypes.ListValue, was: %T`, apiTokensAttribute))
	}

	colorAttribute, ok := attributes["color"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`color is missing from object`)

		return nil, diags
	}

	colorVal, ok := colorAttribute.(basetypes.StringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`color expected to be basetypes.StringValue, was: %T`, colorAttribute))
	}

	idAttribute, ok := attributes["id"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`id is missing from object`)

		return nil, diags
	}

	idVal, ok := idAttribute.(basetypes.Int64Value)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`id expected to be basetypes.Int64Value, was: %T`, idAttribute))
	}

	inboundAddressAttribute, ok := attributes["inbound_address"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`inbound_address is missing from object`)

		return nil, diags
	}

	inboundAddressVal, ok := inboundAddressAttribute.(basetypes.StringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`inbound_address expected to be basetypes.StringValue, was: %T`, inboundAddressAttribute))
	}

	inboundDomainAttribute, ok := attributes["inbound_domain"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`inbound_domain is missing from object`)

		return nil, diags
	}

	inboundDomainVal, ok := inboundDomainAttribute.(basetypes.StringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`inbound_domain expected to be basetypes.StringValue, was: %T`, inboundDomainAttribute))
	}

	inboundHashAttribute, ok := attributes["inbound_hash"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`inbound_hash is missing from object`)

		return nil, diags
	}

	inboundHashVal, ok := inboundHashAttribute.(basetypes.StringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`inbound_hash expected to be basetypes.StringValue, was: %T`, inboundHashAttribute))
	}

	inboundHookUrlAttribute, ok := attributes["inbound_hook_url"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`inbound_hook_url is missing from object`)

		return nil, diags
	}

	inboundHookUrlVal, ok := inboundHookUrlAttribute.(basetypes.StringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`inbound_hook_url expected to be basetypes.StringValue, was: %T`, inboundHookUrlAttribute))
	}

	inboundSpamThresholdAttribute, ok := attributes["inbound_spam_threshold"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`inbound_spam_threshold is missing from object`)

		return nil, diags
	}

	inboundSpamThresholdVal, ok := inboundSpamThresholdAttribute.(basetypes.Int64Value)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`inbound_spam_threshold expected to be basetypes.Int64Value, was: %T`, inboundSpamThresholdAttribute))
	}

	nameAttribute, ok := attributes["name"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`name is missing from object`)

		return nil, diags
	}

	nameVal, ok := nameAttribute.(basetypes.StringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`name expected to be basetypes.StringValue, was: %T`, nameAttribute))
	}

	postFirstOpenOnlyAttribute, ok := attributes["post_first_open_only"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`post_first_open_only is missing from object`)

		return nil, diags
	}

	postFirstOpenOnlyVal, ok := postFirstOpenOnlyAttribute.(basetypes.BoolValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`post_first_open_only expected to be basetypes.BoolValue, was: %T`, postFirstOpenOnlyAttribute))
	}

	rawEmailEnabledAttribute, ok := attributes["raw_email_enabled"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`raw_email_enabled is missing from object`)

		return nil, diags
	}

	rawEmailEnabledVal, ok := rawEmailEnabledAttribute.(basetypes.BoolValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`raw_email_enabled expected to be basetypes.BoolValue, was: %T`, rawEmailEnabledAttribute))
	}

	serverLinkAttribute, ok := attributes["server_link"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`server_link is missing from object`)

		return nil, diags
	}

	serverLinkVal, ok := serverLinkAttribute.(basetypes.StringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`server_link expected to be basetypes.StringValue, was: %T`, serverLinkAttribute))
	}

	smtpApiActivatedAttribute, ok := attributes["smtp_api_activated"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`smtp_api_activated is missing from object`)

		return nil, diags
	}

	smtpApiActivatedVal, ok := smtpApiActivatedAttribute.(basetypes.BoolValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`smtp_api_activated expected to be basetypes.BoolValue, was: %T`, smtpApiActivatedAttribute))
	}

	trackLinksAttribute, ok := attributes["track_links"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`track_links is missing from object`)

		return nil, diags
	}

	trackLinksVal, ok := trackLinksAttribute.(basetypes.StringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`track_links expected to be basetypes.StringValue, was: %T`, trackLinksAttribute))
	}

	trackOpensAttribute, ok := attributes["track_opens"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`track_opens is missing from object`)

		return nil, diags
	}

	trackOpensVal, ok := trackOpensAttribute.(basetypes.BoolValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`track_opens expected to be basetypes.BoolValue, was: %T`, trackOpensAttribute))
	}

	if diags.HasError() {
		return nil, diags
	}

	return ServersValue{
		ApiTokens:            apiTokensVal,
		Color:                colorVal,
		Id:                   idVal,
		InboundAddress:       inboundAddressVal,
		InboundDomain:        inboundDomainVal,
		InboundHash:          inboundHashVal,
		InboundHookUrl:       inboundHookUrlVal,
		InboundSpamThreshold: inboundSpamThresholdVal,
		Name:                 nameVal,
		PostFirstOpenOnly:    postFirstOpenOnlyVal,
		RawEmailEnabled:      rawEmailEnabledVal,
		ServerLink:           serverLinkVal,
		SmtpApiActivated:     smtpApiActivatedVal,
		TrackLinks:           trackLinksVal,
		TrackOpens:           trackOpensVal,
		state:                attr.ValueStateKnown,
	}, diags
}

func NewServersValueNull() ServersValue {
	return ServersValue{
		state: attr.ValueStateNull,
	}
}

func NewServersValueUnknown() ServersValue {
	return ServersValue{
		state: attr.ValueStateUnknown,
	}
}

func NewServersValue(attributeTypes map[string]attr.Type, attributes map[string]attr.Value) (ServersValue, diag.Diagnostics) {
	var diags diag.Diagnostics

	// Reference: https://github.com/hashicorp/terraform-plugin-framework/issues/521
	ctx := context.Background()

	for name, attributeType := range attributeTypes {
		attribute, ok := attributes[name]

		if !ok {
			diags.AddError(
				"Missing ServersValue Attribute Value",
				"While creating a ServersValue value, a missing attribute value was detected. "+
					"A ServersValue must contain values for all attributes, even if null or unknown. "+
					"This is always an issue with the provider and should be reported to the provider developers.\n\n"+
					fmt.Sprintf("ServersValue Attribute Name (%s) Expected Type: %s", name, attributeType.String()),
			)

			continue
		}

		if !attributeType.Equal(attribute.Type(ctx)) {
			diags.AddError(
				"Invalid ServersValue Attribute Type",
				"While creating a ServersValue value, an invalid attribute value was detected. "+
					"A ServersValue must use a matching attribute type for the value. "+
					"This is always an issue with the provider and should be reported to the provider developers.\n\n"+
					fmt.Sprintf("ServersValue Attribute Name (%s) Expected Type: %s\n", name, attributeType.String())+
					fmt.Sprintf("ServersValue Attribute Name (%s) Given Type: %s", name, attribute.Type(ctx)),
			)
		}
	}

	for name := range attributes {
		_, ok := attributeTypes[name]

		if !ok {
			diags.AddError(
				"Extra ServersValue Attribute Value",
				"While creating a ServersValue value, an extra attribute value was detected. "+
					"A ServersValue must not contain values beyond the expected attribute types. "+
					"This is always an issue with the provider and should be reported to the provider developers.\n\n"+
					fmt.Sprintf("Extra ServersValue Attribute Name: %s", name),
			)
		}
	}

	if diags.HasError() {
		return NewServersValueUnknown(), diags
	}

	apiTokensAttribute, ok := attributes["api_tokens"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`api_tokens is missing from object`)

		return NewServersValueUnknown(), diags
	}

	apiTokensVal, ok := apiTokensAttribute.(basetypes.ListValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`api_tokens expected to be basetypes.ListValue, was: %T`, apiTokensAttribute))
	}

	colorAttribute, ok := attributes["color"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`color is missing from object`)

		return NewServersValueUnknown(), diags
	}

	colorVal, ok := colorAttribute.(basetypes.StringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`color expected to be basetypes.StringValue, was: %T`, colorAttribute))
	}

	idAttribute, ok := attributes["id"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`id is missing from object`)

		return NewServersValueUnknown(), diags
	}

	idVal, ok := idAttribute.(basetypes.Int64Value)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`id expected to be basetypes.Int64Value, was: %T`, idAttribute))
	}

	inboundAddressAttribute, ok := attributes["inbound_address"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`inbound_address is missing from object`)

		return NewServersValueUnknown(), diags
	}

	inboundAddressVal, ok := inboundAddressAttribute.(basetypes.StringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`inbound_address expected to be basetypes.StringValue, was: %T`, inboundAddressAttribute))
	}

	inboundDomainAttribute, ok := attributes["inbound_domain"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`inbound_domain is missing from object`)

		return NewServersValueUnknown(), diags
	}

	inboundDomainVal, ok := inboundDomainAttribute.(basetypes.StringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`inbound_domain expected to be basetypes.StringValue, was: %T`, inboundDomainAttribute))
	}

	inboundHashAttribute, ok := attributes["inbound_hash"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`inbound_hash is missing from object`)

		return NewServersValueUnknown(), diags
	}

	inboundHashVal, ok := inboundHashAttribute.(basetypes.StringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`inbound_hash expected to be basetypes.StringValue, was: %T`, inboundHashAttribute))
	}

	inboundHookUrlAttribute, ok := attributes["inbound_hook_url"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`inbound_hook_url is missing from object`)

		return NewServersValueUnknown(), diags
	}

	inboundHookUrlVal, ok := inboundHookUrlAttribute.(basetypes.StringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`inbound_hook_url expected to be basetypes.StringValue, was: %T`, inboundHookUrlAttribute))
	}

	inboundSpamThresholdAttribute, ok := attributes["inbound_spam_threshold"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`inbound_spam_threshold is missing from object`)

		return NewServersValueUnknown(), diags
	}

	inboundSpamThresholdVal, ok := inboundSpamThresholdAttribute.(basetypes.Int64Value)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`inbound_spam_threshold expected to be basetypes.Int64Value, was: %T`, inboundSpamThresholdAttribute))
	}

	nameAttribute, ok := attributes["name"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`name is missing from object`)

		return NewServersValueUnknown(), diags
	}

	nameVal, ok := nameAttribute.(basetypes.StringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`name expected to be basetypes.StringValue, was: %T`, nameAttribute))
	}

	postFirstOpenOnlyAttribute, ok := attributes["post_first_open_only"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`post_first_open_only is missing from object`)

		return NewServersValueUnknown(), diags
	}

	postFirstOpenOnlyVal, ok := postFirstOpenOnlyAttribute.(basetypes.BoolValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`post_first_open_only expected to be basetypes.BoolValue, was: %T`, postFirstOpenOnlyAttribute))
	}

	rawEmailEnabledAttribute, ok := attributes["raw_email_enabled"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`raw_email_enabled is missing from object`)

		return NewServersValueUnknown(), diags
	}

	rawEmailEnabledVal, ok := rawEmailEnabledAttribute.(basetypes.BoolValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`raw_email_enabled expected to be basetypes.BoolValue, was: %T`, rawEmailEnabledAttribute))
	}

	serverLinkAttribute, ok := attributes["server_link"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`server_link is missing from object`)

		return NewServersValueUnknown(), diags
	}

	serverLinkVal, ok := serverLinkAttribute.(basetypes.StringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`server_link expected to be basetypes.StringValue, was: %T`, serverLinkAttribute))
	}

	smtpApiActivatedAttribute, ok := attributes["smtp_api_activated"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`smtp_api_activated is missing from object`)

		return NewServersValueUnknown(), diags
	}

	smtpApiActivatedVal, ok := smtpApiActivatedAttribute.(basetypes.BoolValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`smtp_api_activated expected to be basetypes.BoolValue, was: %T`, smtpApiActivatedAttribute))
	}

	trackLinksAttribute, ok := attributes["track_links"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`track_links is missing from object`)

		return NewServersValueUnknown(), diags
	}

	trackLinksVal, ok := trackLinksAttribute.(basetypes.StringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`track_links expected to be basetypes.StringValue, was: %T`, trackLinksAttribute))
	}

	trackOpensAttribute, ok := attributes["track_opens"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`track_opens is missing from object`)

		return NewServersValueUnknown(), diags
	}

	trackOpensVal, ok := trackOpensAttribute.(basetypes.BoolValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`track_opens expected to be basetypes.BoolValue, was: %T`, trackOpensAttribute))
	}

	if diags.HasError() {
		return NewServersValueUnknown(), diags
	}

	return ServersValue{
		ApiTokens:            apiTokensVal,
		Color:                colorVal,
		Id:                   idVal,
		InboundAddress:       inboundAddressVal,
		InboundDomain:        inboundDomainVal,
		InboundHash:          inboundHashVal,
		InboundHookUrl:       inboundHookUrlVal,
		InboundSpamThreshold: inboundSpamThresholdVal,
		Name:                 nameVal,
		PostFirstOpenOnly:    postFirstOpenOnlyVal,
		RawEmailEnabled:      rawEmailEnabledVal,
		ServerLink:           serverLinkVal,
		SmtpApiActivated:     smtpApiActivatedVal,
		TrackLinks:           trackLinksVal,
		TrackOpens:           trackOpensVal,
		state:                attr.ValueStateKnown,
	}, diags
}

func NewServersValueMust(attributeTypes map[string]attr.Type, attributes map[string]attr.Value) ServersValue {
	object, diags := NewServersValue(attributeTypes, attributes)

	if diags.HasError() {
		// This could potentially be added to the diag package.
		diagsStrings := make([]string, 0, len(diags))

		for _, diagnostic := range diags {
			diagsStrings = append(diagsStrings, fmt.Sprintf(
				"%s | %s | %s",
				diagnostic.Severity(),
				diagnostic.Summary(),
				diagnostic.Detail()))
		}

		panic("NewServersValueMust received error(s): " + strings.Join(diagsStrings, "\n"))
	}

	return object
}

func (t ServersType) ValueFromTerraform(ctx context.Context, in tftypes.Value) (attr.Value, error) {
	if in.Type() == nil {
		return NewServersValueNull(), nil
	}

	if !in.Type().Equal(t.TerraformType(ctx)) {
		return nil, fmt.Errorf("expected %s, got %s", t.TerraformType(ctx), in.Type())
	}

	if !in.IsKnown() {
		return NewServersValueUnknown(), nil
	}

	if in.IsNull() {
		return NewServersValueNull(), nil
	}

	attributes := map[string]attr.Value{}

	val := map[string]tftypes.Value{}

	err := in.As(&val)

	if err != nil {
		return nil, err
	}

	for k, v := range val {
		a, err := t.AttrTypes[k].ValueFromTerraform(ctx, v)

		if err != nil {
			return nil, err
		}

		attributes[k] = a
	}

	return NewServersValueMust(ServersValue{}.AttributeTypes(ctx), attributes), nil
}

func (t ServersType) ValueType(ctx context.Context) attr.Value {
	return ServersValue{}
}

var _ basetypes.ObjectValuable = ServersValue{}

type ServersValue struct {
	ApiTokens            basetypes.ListValue   `tfsdk:"api_tokens"`
	Color                basetypes.StringValue `tfsdk:"color"`
	Id                   basetypes.Int64Value  `tfsdk:"id"`
	InboundAddress       basetypes.StringValue `tfsdk:"inbound_address"`
	InboundDomain        basetypes.StringValue `tfsdk:"inbound_domain"`
	InboundHash          basetypes.StringValue `tfsdk:"inbound_hash"`
	InboundHookUrl       basetypes.StringValue `tfsdk:"inbound_hook_url"`
	InboundSpamThreshold basetypes.Int64Value  `tfsdk:"inbound_spam_threshold"`
	Name                 basetypes.StringValue `tfsdk:"name"`
	PostFirstOpenOnly    basetypes.BoolValue   `tfsdk:"post_first_open_only"`
	RawEmailEnabled      basetypes.BoolValue   `tfsdk:"raw_email_enabled"`
	ServerLink           basetypes.StringValue `tfsdk:"server_link"`
	SmtpApiActivated     basetypes.BoolValue   `tfsdk:"smtp_api_activated"`
	TrackLinks           basetypes.StringValue `tfsdk:"track_links"`
	TrackOpens           basetypes.BoolValue   `tfsdk:"track_opens"`
	state                attr.ValueState
}

func (v ServersValue) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	attrTypes := make(map[string]tftypes.Type, 15)

	var val tftypes.Value
	var err error

	attrTypes["api_tokens"] = basetypes.ListType{
		ElemType: types.StringType,
	}.TerraformType(ctx)
	attrTypes["color"] = basetypes.StringType{}.TerraformType(ctx)
	attrTypes["id"] = basetypes.Int64Type{}.TerraformType(ctx)
	attrTypes["inbound_address"] = basetypes.StringType{}.TerraformType(ctx)
	attrTypes["inbound_domain"] = basetypes.StringType{}.TerraformType(ctx)
	attrTypes["inbound_hash"] = basetypes.StringType{}.TerraformType(ctx)
	attrTypes["inbound_hook_url"] = basetypes.StringType{}.TerraformType(ctx)
	attrTypes["inbound_spam_threshold"] = basetypes.Int64Type{}.TerraformType(ctx)
	attrTypes["name"] = basetypes.StringType{}.TerraformType(ctx)
	attrTypes["post_first_open_only"] = basetypes.BoolType{}.TerraformType(ctx)
	attrTypes["raw_email_enabled"] = basetypes.BoolType{}.TerraformType(ctx)
	attrTypes["server_link"] = basetypes.StringType{}.TerraformType(ctx)
	attrTypes["smtp_api_activated"] = basetypes.BoolType{}.TerraformType(ctx)
	attrTypes["track_links"] = basetypes.StringType{}.TerraformType(ctx)
	attrTypes["track_opens"] = basetypes.BoolType{}.TerraformType(ctx)

	objectType := tftypes.Object{AttributeTypes: attrTypes}

	switch v.state {
	case attr.ValueStateKnown:
		vals := make(map[string]tftypes.Value, 15)

		val, err = v.ApiTokens.ToTerraformValue(ctx)

		if err != nil {
			return tftypes.NewValue(objectType, tftypes.UnknownValue), err
		}

		vals["api_tokens"] = val

		val, err = v.Color.ToTerraformValue(ctx)

		if err != nil {
			return tftypes.NewValue(objectType, tftypes.UnknownValue), err
		}

		vals["color"] = val

		val, err = v.Id.ToTerraformValue(ctx)

		if err != nil {
			return tftypes.NewValue(objectType, tftypes.UnknownValue), err
		}

		vals["id"] = val

		val, err = v.InboundAddress.ToTerraformValue(ctx)

		if err != nil {
			return tftypes.NewValue(objectType, tftypes.UnknownValue), err
		}

		vals["inbound_address"] = val

		val, err = v.InboundDomain.ToTerraformValue(ctx)

		if err != nil {
			return tftypes.NewValue(objectType, tftypes.UnknownValue), err
		}

		vals["inbound_domain"] = val

		val, err = v.InboundHash.ToTerraformValue(ctx)

		if err != nil {
			return tftypes.NewValue(objectType, tftypes.UnknownValue), err
		}

		vals["inbound_hash"] = val

		val, err = v.InboundHookUrl.ToTerraformValue(ctx)

		if err != nil {
			return tftypes.NewValue(objectType, tftypes.UnknownValue), err
		}

		vals["inbound_hook_url"] = val

		val, err = v.InboundSpamThreshold.ToTerraformValue(ctx)

		if err != nil {
			return tftypes.NewValue(objectType, tftypes.UnknownValue), err
		}

		vals["inbound_spam_threshold"] = val

		val, err = v.Name.ToTerraformValue(ctx)

		if err != nil {
			return tftypes.NewValue(objectType, tftypes.UnknownValue), err
		}

		vals["name"] = val

		val, err = v.PostFirstOpenOnly.ToTerraformValue(ctx)

		if err != nil {
			return tftypes.NewValue(objectType, tftypes.UnknownValue), err
		}

		vals["post_first_open_only"] = val

		val, err = v.RawEmailEnabled.ToTerraformValue(ctx)

		if err != nil {
			return tftypes.NewValue(objectType, tftypes.UnknownValue), err
		}

		vals["raw_email_enabled"] = val

		val, err = v.ServerLink.ToTerraformValue(ctx)

		if err != nil {
			return tftypes.NewValue(objectType, tftypes.UnknownValue), err
		}

		vals["server_link"] = val

		val, err = v.SmtpApiActivated.ToTerraformValue(ctx)

		if err != nil {
			return tftypes.NewValue(objectType, tftypes.UnknownValue), err
		}

		vals["smtp_api_activated"] = val

		val, err = v.TrackLinks.ToTerraformValue(ctx)

		if err != nil {
			return tftypes.NewValue(objectType, tftypes.UnknownValue), err
		}

		vals["track_links"] = val

		val, err = v.TrackOpens.ToTerraformValue(ctx)

		if err != nil {
			return tftypes.NewValue(objectType, tftypes.UnknownValue), err
		}

		vals["track_opens"] = val

		if err := tftypes.ValidateValue(objectType, vals); err != nil {
			return tftypes.NewValue(objectType, tftypes.UnknownValue), err
		}

		return tftypes.NewValue(objectType, vals), nil
	case attr.ValueStateNull:
		return tftypes.NewValue(objectType, nil), nil
	case attr.ValueStateUnknown:
		return tftypes.NewValue(objectType, tftypes.UnknownValue), nil
	default:
		panic(fmt.Sprintf("unhandled Object state in ToTerraformValue: %s", v.state))
	}
}

func (v ServersValue) IsNull() bool {
	return v.state == attr.ValueStateNull
}

func (v ServersValue) IsUnknown() bool {
	return v.state == attr.ValueStateUnknown
}

func (v ServersValue) String() string {
	return "ServersValue"
}

func (v ServersValue) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	var diags diag.Diagnostics

	var apiTokensVal basetypes.ListValue
	switch {
	case v.ApiTokens.IsUnknown():
		apiTokensVal = types.ListUnknown(types.StringType)
	case v.ApiTokens.IsNull():
		apiTokensVal = types.ListNull(types.StringType)
	default:
		var d diag.Diagnostics
		apiTokensVal, d = types.ListValue(types.StringType, v.ApiTokens.Elements())
		diags.Append(d...)
	}

	if diags.HasError() {
		return types.ObjectUnknown(map[string]attr.Type{
			"api_tokens": basetypes.ListType{
				ElemType: types.StringType,
			},
			"color":                  basetypes.StringType{},
			"id":                     basetypes.Int64Type{},
			"inbound_address":        basetypes.StringType{},
			"inbound_domain":         basetypes.StringType{},
			"inbound_hash":           basetypes.StringType{},
			"inbound_hook_url":       basetypes.StringType{},
			"inbound_spam_threshold": basetypes.Int64Type{},
			"name":                   basetypes.StringType{},
			"post_first_open_only":   basetypes.BoolType{},
			"raw_email_enabled":      basetypes.BoolType{},
			"server_link":            basetypes.StringType{},
			"smtp_api_activated":     basetypes.BoolType{},
			"track_links":            basetypes.StringType{},
			"track_opens":            basetypes.BoolType{},
		}), diags
	}

	attributeTypes := map[string]attr.Type{
		"api_tokens": basetypes.ListType{
			ElemType: types.StringType,
		},
		"color":                  basetypes.StringType{},
		"id":                     basetypes.Int64Type{},
		"inbound_address":        basetypes.StringType{},
		"inbound_domain":         basetypes.StringType{},
		"inbound_hash":           basetypes.StringType{},
		"inbound_hook_url":       basetypes.StringType{},
		"inbound_spam_threshold": basetypes.Int64Type{},
		"name":                   basetypes.StringType{},
		"post_first_open_only":   basetypes.BoolType{},
		"raw_email_enabled":      basetypes.BoolType{},
		"server_link":            basetypes.StringType{},
		"smtp_api_activated":     basetypes.BoolType{},
		"track_links":            basetypes.StringType{},
		"track_opens":            basetypes.BoolType{},
	}

	if v.IsNull() {
		return types.ObjectNull(attributeTypes), diags
	}

	if v.IsUnknown() {
		return types.ObjectUnknown(attributeTypes), diags
	}

	objVal, diags := types.ObjectValue(
		attributeTypes,
		map[string]attr.Value{
			"api_tokens":             apiTokensVal,
			"color":                  v.Color,
			"id":                     v.Id,
			"inbound_address":        v.InboundAddress,
			"inbound_domain":         v.InboundDomain,
			"inbound_hash":           v.InboundHash,
			"inbound_hook_url":       v.InboundHookUrl,
			"inbound_spam_threshold": v.InboundSpamThreshold,
			"name":                   v.Name,
			"post_first_open_only":   v.PostFirstOpenOnly,
			"raw_email_enabled":      v.RawEmailEnabled,
			"server_link":            v.ServerLink,
			"smtp_api_activated":     v.SmtpApiActivated,
			"track_links":            v.TrackLinks,
			"track_opens":            v.TrackOpens,
		})

	return objVal, diags
}

func (v ServersValue) Equal(o attr.Value) bool {
	other, ok := o.(ServersValue)

	if !ok {
		return false
	}

	if v.state != other.state {
		return false
	}

	if v.state != attr.ValueStateKnown {
		return true
	}

	if !v.ApiTokens.Equal(other.ApiTokens) {
		return false
	}

	if !v.Color.Equal(other.Color) {
		return false
	}

	if !v.Id.Equal(other.Id) {
		return false
	}

	if !v.InboundAddress.Equal(other.InboundAddress) {
		return false
	}

	if !v.InboundDomain.Equal(other.InboundDomain) {
		return false
	}

	if !v.InboundHash.Equal(other.InboundHash) {
		return false
	}

	if !v.InboundHookUrl.Equal(other.InboundHookUrl) {
		return false
	}

	if !v.InboundSpamThreshold.Equal(other.InboundSpamThreshold) {
		return false
	}

	if !v.Name.Equal(other.Name) {
		return false
	}

	if !v.PostFirstOpenOnly.Equal(other.PostFirstOpenOnly) {
		return false
	}

	if !v.RawEmailEnabled.Equal(other.RawEmailEnabled) {
		return false
	}

	if !v.ServerLink.Equal(other.ServerLink) {
		return false
	}

	if !v.SmtpApiActivated.Equal(other.SmtpApiActivated) {
		return false
	}

	if !v.TrackLinks.Equal(other.TrackLinks) {
		return false
	}

	if !v.TrackOpens.Equal(other.TrackOpens) {
		return false
	}

	return true
}

func (v ServersValue) Type(ctx context.Context) attr.Type {
	return ServersType{
		basetypes.ObjectType{
			AttrTypes: v.AttributeTypes(ctx),
		},
	}
}

func (v ServersValue) AttributeTypes(ctx context.Context) map[string]attr.Type {
	return map[string]attr.Type{
		"api_tokens": basetypes.ListType{
			ElemType: types.StringType,
		},
		"color":                  basetypes.StringType{},
		"id":                     basetypes.Int64Type{},
		"inbound_address":        basetypes.StringType{},
		"inbound_domain":         basetypes.StringType{},
		"inbound_hash":           basetypes.StringType{},
		"inbound_hook_url":       basetypes.StringType{},
		"inbound_spam_threshold": basetypes.Int64Type{},
		"name":                   basetypes.StringType{},
		"post_first_open_only":   basetypes.BoolType{},
		"raw_email_enabled":      basetypes.BoolType{},
		"server_link":            basetypes.StringType{},
		"smtp_api_activated":     basetypes.BoolType{},
		"track_links":            basetypes.StringType{},
		"track_opens":            basetypes.BoolType{},
	}
}
