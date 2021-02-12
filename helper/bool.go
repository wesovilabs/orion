// Package converter content relative to the package
package helper

import (
	"github.com/hashicorp/hcl/v2"
	"github.com/zclconf/go-cty/cty"

	"github.com/wesovilabs-tools/orion/internal/errors"
)

// GetValueAsBool return the value as a bool.
func GetValueAsBool(value cty.Value, def bool) bool {
	if !value.IsNull() {
		return value.True()
	}

	return def
}

// GetExpressionValueAsBool return the expression as bool.
func GetExpressionValueAsBool(ctx *hcl.EvalContext, expr hcl.Expression, def bool) (bool, errors.Error) {
	if expr == nil || expr.Range().Empty() {
		return def, nil
	}
	value, d := expr.Value(ctx)
	if err := errors.EvalDiagnostics(d); err != nil {
		return def, err.
			AddMeta(errors.MetaLocation, expr.Range().String())
	}
	if value.Type() == cty.Bool {
		return value.True(), nil
	}

	return false, nil
}
