package provider

import (
	"strconv"

	"github.com/hashicorp/terraform-plugin-framework/types"
)

func TypeStringToInt(value types.String) int {
	valueInt, _ := strconv.ParseInt(value.ValueString(), 10, 64)
	return int(valueInt)
}
