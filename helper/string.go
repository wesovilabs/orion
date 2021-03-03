package helper

import (
	"fmt"

	"github.com/hashicorp/hcl/v2"
	"github.com/zclconf/go-cty/cty"

	"github.com/wesovilabs/orion/internal/errors"
)

// GetExpressionValueAsString return the value fo the expression as a string.
func GetExpressionValueAsString(ctx *hcl.EvalContext, expr hcl.Expression, def string) (string, errors.Error) {
	if expr == nil || expr.Range().Empty() {
		return def, nil
	}
	value, d := expr.Value(ctx)
	if err := errors.EvalDiagnostics(d); err != nil {
		return "", err
	}
	return ToString(value)
}

// ToString converts a given cty.Value into a string.
func ToString(v cty.Value) (string, errors.Error) {
	if v.IsNull() {
		return "", nil
	}
	switch {
	case IsMap(v):
		result := ""
		valueMap := v.AsValueMap()
		for k, v := range valueMap {
			item, err := ToString(v)
			if err != nil {
				return "", err
			}
			result += k + ":[" + item + "]"
		}
		return result, nil
	case IsSlice(v):
		result := ""
		slice := v.AsValueSlice()
		for index := range slice {
			valueItem := slice[index]
			if index != 0 {
				result += ", "
			}
			item, err := ToString(valueItem)
			if err != nil {
				return "", err
			}
			result += item
		}
		return result, nil
	case v.Type() == cty.String:
		return v.AsString(), nil
	case v.Type() == cty.Number:
		return v.AsBigFloat().String(), nil
	case v.Type() == cty.Bool:
		return fmt.Sprintf("%v", v.True()), nil
	}
	return "", errors.Unexpected("value type %s is not supported", v.Type().GoString())
}

// AttributeToStringWithoutContext convert a hc Attribute into a string value.
func AttributeToStringWithoutContext(attribute *hcl.Attribute) (string, errors.Error) {
	if len(attribute.Expr.Variables()) > 0 {
		return "", errors.IncorrectUsage("variables are not permitted in attribute %s", attribute.Name)
	}
	val, err := EvalAttribute(nil, attribute)
	if err != nil {
		return "", err
	}
	return ToStrictString(val)
}
