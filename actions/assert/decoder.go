package assert

import (
	"github.com/hashicorp/hcl/v2"
	"github.com/wesovilabs/orion/actions"
	"github.com/wesovilabs/orion/internal/errors"
)

const (
	// BlockAssert identifier of the plugin.
	BlockAssert  = "assert"
	argAssertion = "assertion"
)

var schemaAssert = &hcl.BodySchema{
	Attributes: append(actions.BaseArguments, []hcl.AttributeSchema{
		{Name: argAssertion, Required: true},
	}...),
}

// Decoder implements interface plugin.Decoder.
type Decoder struct{}

// BlockHeaderSchema return the header schema for the plugin.
func (dec *Decoder) BlockHeaderSchema() hcl.BlockHeaderSchema {
	return hcl.BlockHeaderSchema{
		Type:       BlockAssert,
		LabelNames: nil,
	}
}

// DecodeBlock required method implementation by interface Decoder.
func (dec *Decoder) DecodeBlock(block *hcl.Block) (actions.Action, errors.Error) {
	bodyContent, d := block.Body.Content(schemaAssert)
	if err := errors.EvalDiagnostics(d); err != nil {
		return nil, err
	}
	assert := &Assert{Base: &actions.Base{}}
	if err := populateAttributes(assert, bodyContent.Attributes); err != nil {
		return nil, err
	}
	return assert, nil
}

func populateAttributes(assert *Assert, attrs hcl.Attributes) errors.Error {
	for name := range attrs {
		attribute := attrs[name]
		switch {
		case actions.IsPluginBaseArgument(name):
			if err := actions.SetBaseArgs(assert, attribute); err != nil {
				return err
			}
		case name == argAssertion:
			assert.SetAssertion(attribute.Expr)
		default:
			return errors.ThrowUnsupportedArgument(BlockAssert, name)
		}
	}
	return nil
}
