package helper

import (
	"github.com/hashicorp/hcl/v2"
	"github.com/zclconf/go-cty/cty"

	"github.com/wesovilabs/orion/internal/errors"
)

// GetExpressionValueAsInt evaluate the expression and return a int value.
func GetExpressionValueAsInt(ctx *hcl.EvalContext, expr hcl.Expression, def int) (int, errors.Error) {
	if expr == nil || expr.Range().Empty() {
		return def, nil
	}
	value, diagnostics := expr.Value(ctx)
	if diagnostics != nil && diagnostics.HasErrors() {
		return 0, errors.IncorrectUsage(diagnostics.Error()).AddMeta(errors.MetaLocation, expr.Range().String())
	}
	if value.Type() == cty.Number {
		valFloat := value.AsBigFloat()
		out, _ := valFloat.Int64()

		return int(out), nil
	}

	return def, nil
}
