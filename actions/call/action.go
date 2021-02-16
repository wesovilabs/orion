package call

import (
	"github.com/hashicorp/hcl/v2"
	"github.com/wesovilabs/orion/actions"
	"github.com/wesovilabs/orion/context"
	"github.com/wesovilabs/orion/helper"
	"github.com/wesovilabs/orion/internal/errors"
)

// Call used to invoke functions
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
		case actions.IsPluginBaseArgument(name):
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
		return errors.ThrowsExceeddedNumberOfBlocks(blockWith, 1)
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

func (c *Call) Execute(ctx context.OrionContext) errors.Error {
	return nil
}
