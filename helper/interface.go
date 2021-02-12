package helper

import (
	"github.com/hashicorp/hcl/v2"
	"github.com/zclconf/go-cty/cty"

	"github.com/wesovilabs/orion/internal/errors"
)

// GetExpressionValueAsInterface evaluate an hcl expresion and converts into an interface.
func GetExpressionValueAsInterface(ctx *hcl.EvalContext, expr hcl.Expression, def interface{}) (interface{},
	errors.Error) {
	if expr == nil || expr.Range().Empty() {
		return def, nil
	}
	value, diagnostics := expr.Value(ctx)
	if diagnostics != nil && diagnostics.HasErrors() {
		return nil, errors.IncorrectUsage(diagnostics.Error()).AddMeta(errors.MetaLocation, expr.Range().String())
	}

	return ToInterface(value)
}

// ToInterface converts a given cty.Value into an interface.
func ToInterface(v cty.Value) (interface{}, errors.Error) {
	switch {
	case IsMap(v):
		result := make(map[string]interface{})
		valueMap := v.AsValueMap()
		for k, v := range valueMap {
			item, err := ToInterface(v)
			if err != nil {
				return "", err
			}
			result[k] = item
		}

		return result, nil
	case IsSlice(v):
		slice := v.AsValueSlice()
		result := make([]interface{}, len(slice))
		for index := range slice {
			valueItem := slice[index]
			item, err := ToString(valueItem)
			if err != nil {
				return "", err
			}
			result[index] = item
		}

		return result, nil
	case v.Type() == cty.String:
		return v.AsString(), nil
	case v.Type() == cty.Number:
		valueFloat := v.AsBigFloat()
		if valueFloat.IsInf() {
			vFloat, _ := valueFloat.Int64()
			return vFloat, nil
		}
		result, _ := valueFloat.Float64()

		return result, nil
	case v.Type() == cty.Bool:
		return v.True(), nil
	}

	return "", errors.Unexpected("v type %s is not supported", v.Type().GoString())
}
