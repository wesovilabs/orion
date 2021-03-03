package sleep

import (
	"github.com/hashicorp/hcl/v2"
	"github.com/wesovilabs/orion/actions"
	"github.com/wesovilabs/orion/internal/errors"
)

const (
	// BlockSleep identifier of the sleep block.
	BlockSleep = "sleep"
	// AttributeDuration name of the argument used to specify the sleep duration.
	AttributeDuration = "duration"
)

var schemaSleep = &hcl.BodySchema{
	Attributes: append(actions.BaseArguments, []hcl.AttributeSchema{
		{Name: AttributeDuration, Required: true},
	}...),
}

// Decoder implements interface plugin.Decoder.
type Decoder struct{}

// BlockHeaderSchema return the header schema for the plugin.
func (dec *Decoder) BlockHeaderSchema() hcl.BlockHeaderSchema {
	return hcl.BlockHeaderSchema{
		Type:       BlockSleep,
		LabelNames: nil,
	}
}

// DecodeBlock required to implement the plugin interface.
func (dec *Decoder) DecodeBlock(block *hcl.Block) (actions.Action, errors.Error) {
	bodyContent, d := block.Body.Content(schemaSleep)
	if err := errors.EvalDiagnostics(d); err != nil {
		return nil, err
	}
	assert := &Sleep{Base: &actions.Base{}}
	if err := populateAttributes(assert, bodyContent.Attributes); err != nil {
		return nil, err
	}
	return assert, nil
}

func populateAttributes(s *Sleep, attrs hcl.Attributes) errors.Error {
	for name := range attrs {
		attribute := attrs[name]

		switch {
		case actions.IsCommonAttribute(name):
			if err := actions.SetBaseArgs(s, attribute); err != nil {
				return err
			}
		case name == AttributeDuration:
			s.SetDuration(attribute.Expr)
		default:
			return errors.ThrowUnsupportedArgument(BlockSleep, name)
		}
	}

	return nil
}
