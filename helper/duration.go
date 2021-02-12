package helper

import (
	"time"

	"github.com/hashicorp/hcl/v2"

	"github.com/wesovilabs/orion/internal/errors"
)

// GetExpressionValueAsDuration evaluate the expression and return a duration value.
func GetExpressionValueAsDuration(ctx *hcl.EvalContext, expr hcl.Expression, def *time.Duration) (*time.Duration,
	errors.Error) {
	valueStr, err := GetExpressionValueAsString(ctx, expr, "")
	if err != nil {
		return nil, err
	}
	if valueStr == "" {
		return def, nil
	}
	value, parseErr := time.ParseDuration(valueStr)
	if parseErr != nil {
		return nil, errors.Unexpected(parseErr.Error()).AddMeta(errors.MetaLocation, expr.Range().String())
	}

	return &value, nil
}
