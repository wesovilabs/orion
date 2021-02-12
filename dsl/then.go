package dsl

import (
	"github.com/hashicorp/hcl/v2"
	assert "github.com/wesovilabs/orion/actions/assert"

	"github.com/wesovilabs/orion/internal/errors"
)

var schemaThen = &hcl.BodySchema{
	Blocks: []hcl.BlockHeaderSchema{
		{Type: assert.BlockAssert},
	},
	Attributes: nil,
}

// Then represents the stage in a schemaScenario in which the assertions take place
type Then struct {
	*section
}

func decodeThen(block *hcl.Block) (*Then, errors.Error) {
	then := &Then{
		section: newSection(blockThen, block),
	}
	bc, d := block.Body.Content(schemaThen)
	if err := errors.EvalDiagnostics(d); err != nil {
		return nil, err
	}
	if err := then.populateAttributes(bc.Attributes); err != nil {
		return nil, err
	}
	if err := then.populateActions(bc.Blocks); err != nil {
		return nil, err
	}
	return then, nil
}

func (then *Then) populateAttributes(attributes hcl.Attributes) errors.Error {
	if len(attributes) > 0 {
		return errors.ThrowUnsupportedArguments(blockThen)
	}
	return nil
}
