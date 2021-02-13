package dsl

import (
	"github.com/hashicorp/hcl/v2"

	"github.com/wesovilabs/orion/internal/errors"
)

var schemaWhen = &hcl.BodySchema{
	Blocks:     handler.GetBlocksSpec(),
	Attributes: nil,
}

// When represents the stage in a schemaScenario in which we perform the actions
// or actions to be tested.
type When struct {
	*section
}

func decodeWhen(block *hcl.Block) (*When, errors.Error) {
	when := &When{
		section: newSection(blockWhen, block),
	}

	bodyContent, d := block.Body.Content(schemaWhen)
	if err := errors.EvalDiagnostics(d); err != nil {
		return nil, err
	}
	if err := when.populateAttributes(bodyContent.Attributes); err != nil {
		return nil, err
	}
	var err errors.Error
	if when.actions, err = handler.DecodePlugins(bodyContent.Blocks); err != nil {
		return nil, err
	}
	return when, nil
}

func (when *When) populateAttributes(attributes hcl.Attributes) errors.Error {
	if len(attributes) > 0 {
		return errors.ThrowUnsupportedArguments(blockWhen)
	}
	return nil
}
