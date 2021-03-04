package call

import (
	"github.com/hashicorp/hcl/v2"
	"github.com/wesovilabs/orion/actions"
	"github.com/wesovilabs/orion/context"
	"github.com/wesovilabs/orion/helper"
	"github.com/wesovilabs/orion/internal/errors"
	"github.com/zclconf/go-cty/cty"
)

// Call used to invoke functions.
type Call struct {
	*actions.Base
	name string
	as   string
	with hcl.Attributes
}

func (c *Call) populateAttributes(attrs hcl.Attributes) errors.Error {
	for name := range attrs {
		attribute := attrs[name]
		switch {
		case actions.IsCommonAttribute(name):
			if err := actions.SetBaseArgs(c, attribute); err != nil {
				return err
			}
		case name == argAs:
			as, err := helper.AttributeToStringWithoutContext(attribute)
			if err != nil {
				return err
			}
			c.as = as
		default:
			return errors.ThrowUnsupportedArgument(BlockCall, name)
		}
	}
	return nil
}

func (c *Call) populateBlocks(blocks hcl.Blocks) errors.Error {
	if len(blocks) > 1 {
		return errors.ThrowsExceededNumberOfBlocks(blockWith, 1)
	}
	if len(blocks) == 0 {
		return nil
	}
	attributes, d := blocks[0].Body.JustAttributes()
	if err := errors.EvalDiagnostics(d); err != nil {
		return err
	}
	c.with = attributes
	return nil
}

func cloneEvalContext(ctx *hcl.EvalContext) *hcl.EvalContext {
	vars := make(map[string]cty.Value)
	for name, val := range ctx.Variables {
		vars[name] = val
	}
	return &hcl.EvalContext{
		Variables: vars,
		Functions: ctx.Functions,
	}
}

func (c *Call) Execute(ctx context.OrionContext) errors.Error {
	fn, ok := ctx.Functions()[c.name]
	oldVariables := cloneEvalContext(ctx.EvalContext()).Variables
	for index := range c.with {
		arg := c.with[index]
		val, d := arg.Expr.Value(ctx.EvalContext())
		if err := errors.EvalDiagnostics(d); err != nil {
			return err
		}
		ctx.EvalContext().Variables[arg.Name] = val
	}
	if !ok {
		return errors.IncorrectUsage("missing required function '%s'", c.name)
	}
	defer func() {
		oldVariables[c.as] = ctx.EvalContext().Variables[c.as]
		ctx.EvalContext().Variables = oldVariables
	}()
	return fn(ctx, c.as)
}
