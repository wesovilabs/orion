// Package set contain types and method to deal with this plugin
package set

import (
	"github.com/hashicorp/hcl/v2"
	log "github.com/sirupsen/logrus"
	"github.com/wesovilabs/orion/actions"
	"github.com/wesovilabs/orion/context"
	"github.com/wesovilabs/orion/helper"
	"github.com/wesovilabs/orion/internal/errors"
)

// Set common block used to define variables.
type Set struct {
	*actions.Base
	name  string
	index hcl.Expression
	value hcl.Expression
}

func (set *Set) populateAttributes(attrs hcl.Attributes) errors.Error {
	for name := range attrs {
		attribute := attrs[name]
		switch {
		case actions.IsPluginBaseArgument(name):
			if err := actions.SetBaseArgs(set, attribute); err != nil {
				return err
			}
		case name == argValue:
			set.value = attribute.Expr
		case name == argIndex:
			set.index = attribute.Expr
		default:
			return errors.ThrowUnsupportedArgument(BlockSet, name)
		}
	}
	return nil
}

// Execute method to run the plugin.
func (set *Set) Execute(ctx context.OrionContext) errors.Error {
	log.Debugf("[%s] It sets value for variable %s", BlockSet, set.name)
	key := set.name
	if set.index != nil {
		index, d := set.index.Value(ctx.EvalContext())
		if err := errors.EvalDiagnostics(d); err != nil {
			return err
		}
		indexV, _ := index.AsBigFloat().Int64()
		return helper.EvaluateArrayItemExpression(ctx.EvalContext(), set.name, int(indexV), set.value)
	}
	return helper.EvalUnorderedExpression(ctx.EvalContext(), map[string]hcl.Expression{
		key: set.value,
	})
}
