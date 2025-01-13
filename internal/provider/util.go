package provider

import (
	"context"
	"strconv"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func TypeStringToInt64(value types.String) int64 {
	valueInt, _ := strconv.ParseInt(value.ValueString(), 10, 64)
	return valueInt
}

func TypeStringToInt(value types.String) int {
	return int(TypeStringToInt64(value))
}

func parseListType(ctx context.Context, list types.List, val []string) (basetypes.ListValue, diag.Diagnostics) {
	listType := list.ElementType(ctx)

	if len(val) == 0 {
		return types.ListNull(listType), nil
	}
	parsed, diags := types.ListValueFrom(ctx, listType, val)

	if diags.HasError() {
		return types.ListNull(listType), diags
	}

	return parsed, nil
}
