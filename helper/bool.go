// Package converter content relative to the package
package helper

import (
	"github.com/hashicorp/hcl/v2"
	"github.com/wesovilabs/orion/internal/errors"
	"github.com/zclconf/go-cty/cty"
)

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

	return def, nil
}
