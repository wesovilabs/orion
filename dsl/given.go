package dsl

import (
	"github.com/hashicorp/hcl/v2"
	"github.com/wesovilabs/orion/actions/set"

	"github.com/wesovilabs/orion/internal/errors"
)

var schemaGiven = &hcl.BodySchema{
	Blocks: []hcl.BlockHeaderSchema{
		{Type: set.BlockSet, LabelNames: []string{set.LabelName}},
	},
	Attributes: nil,
}

// Given represents the stage in a schemaScenario in which the preconditions
// must be declared and the data are initialized.
type Given struct {
	*section
}

func decodeGiven(block *hcl.Block) (*Given, errors.Error) {
	given := &Given{
		section: newSection(blockGiven, block),
	}

	bc, d := block.Body.Content(schemaGiven)
	if err := errors.EvalDiagnostics(d); err != nil {
		return nil, err
	}
	if err := given.populateAttributes(bc.Attributes); err != nil {
		return nil, err
	}
	if err := given.populateActions(bc.Blocks); err != nil {
		return nil, err
	}
	return given, nil
}

func (given *Given) populateAttributes(attributes hcl.Attributes) errors.Error {
	if len(attributes) > 0 {
		return errors.ThrowUnsupportedArguments(blockThen)
	}
	return nil
}
