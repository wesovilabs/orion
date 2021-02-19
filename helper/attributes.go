// Package value contains types and methods to deal with cty value
package helper

import (
	"regexp"
	"strings"

	"github.com/hashicorp/hcl/v2"
	"github.com/zclconf/go-cty/cty"

	"github.com/wesovilabs/orion/internal/errors"
)

const (
	errMissingVariable = "Unknown variable"
	twoArgs            = 2
)

var unknownVarRegExp = regexp.MustCompile(`There is no variable named \"([a-zA-Z0-9_]+)\"`)

// EvalAttribute evaluate the hcl attribute.
func EvalAttribute(ctx *hcl.EvalContext, attribute *hcl.Attribute) (cty.Value, errors.Error) {
	value, d := attribute.Expr.Value(ctx)
	if err := errors.EvalDiagnostics(d); err != nil {
		return cty.NilVal, err
	}

	return value, nil
}

// EvalUnorderedExpression evaluate the the list of expressions.
func EvalUnorderedExpression(ctx *hcl.EvalContext, expressions map[string]hcl.Expression) errors.Error {
	pendingExpressions := make(map[string]hcl.Expression)
	for name, expr := range expressions {
		value, diagnostics := expr.Value(ctx)
		if diagnostics != nil && diagnostics.HasErrors() {
			if err := checkExpr(diagnostics, expressions); err != nil {
				return err
			}
			pendingExpressions[name] = expr
			continue
		}

		ctx.Variables[name] = value
	}

	if len(pendingExpressions) > 0 {
		return EvalUnorderedExpression(ctx, pendingExpressions)
	}

	return nil
}

// nolint
// EvaluateArrayItemExpression evaluate the the list of expressions.
func EvaluateArrayItemExpression(ctx *hcl.EvalContext, name string, index int, val hcl.Expression) errors.Error {
	arrayValue := ctx.Variables[name]
	if arrayValue == cty.NilVal {
		return errors.IncorrectUsage("variable '%s' doesn't exist", name)
	}
	valueSlice := arrayValue.AsValueSlice()
	item, d := val.Value(ctx)
	if err := errors.EvalDiagnostics(d); err != nil {
		return err
	}
	for i := len(valueSlice) - 1; i < index; i++ {
		valueSlice = append(valueSlice, cty.EmptyObjectVal)
	}
	if IsMap(item) {
		newItem := item.AsValueMap()
		old := valueSlice[index].AsValueMap()
		if old != nil {
			for k, v := range newItem {
				old[k] = v
			}
			item = cty.ObjectVal(old)
		} else {
			item = cty.ObjectVal(newItem)
		}

	}
	valueSlice[index] = item
	ctx.Variables[name] = cty.TupleVal(valueSlice)
	return nil
}

func checkExpr(diagnostics hcl.Diagnostics, expressions map[string]hcl.Expression) errors.Error {
	diagnostic := diagnostics[0]
	name := ""
	if strings.EqualFold(diagnostic.Summary, errMissingVariable) {
		if matches := unknownVarRegExp.FindStringSubmatch(diagnostic.Detail); len(matches) == twoArgs {
			name = matches[1]
		}
	}
	if _, ok := expressions[name]; !ok {
		return errors.Unexpected(diagnostics.Error()).ThroughBy(diagnostics.Errs()[0])
	}

	return nil
}
