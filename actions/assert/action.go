// Package assert contains types and methods used by the plugin
package assert

import (
	"github.com/hashicorp/hcl/v2"
	"github.com/wesovilabs/orion/actions"
	"github.com/wesovilabs/orion/context"
	"github.com/wesovilabs/orion/helper"
	"github.com/wesovilabs/orion/internal/errors"
)

const (
	defaultAssertion = false
)

// Assert plugin definition.
type Assert struct {
	*actions.Base
	assertion hcl.Expression
}

// SetAssertion establish the assertion.
func (a *Assert) SetAssertion(assertion hcl.Expression) {
	a.assertion = assertion
}

// Assertion obtain the assertion.
func (a *Assert) Assertion() hcl.Expression {
	return a.assertion
}

// Execute execute the assert plugin.
func (a *Assert) Execute(ctx context.FeatureContext) errors.Error {
	return actions.Execute(ctx, a.Base, func(ctx context.FeatureContext) errors.Error {
		assertion, err := helper.GetExpressionValueAsBool(ctx.EvalContext(), a.assertion, defaultAssertion)
		if err != nil {
			return err
		}
		if !assertion {
			err := errors.Unexpected("assertion is not satisfied!")
			if a.assertion != nil {
				return err.AddMeta(errors.MetaLocation, a.assertion.Range().String())
			}
			return err
		}
		return nil
	})
}
