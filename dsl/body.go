package dsl

import (
	"github.com/hashicorp/hcl/v2"
	"github.com/wesovilabs/orion/actions"
	"github.com/wesovilabs/orion/internal/errors"
)

const blockBody = "body"

var schemaBody = &hcl.BodySchema{
	Blocks: handler.GetBlocksSpec(),
}

type Body struct {
	actions actions.Actions
}

func (b *Body) Actions() actions.Actions {
	return b.actions
}

func decodeBody(block *hcl.Block) (*Body, errors.Error) {
	bodyContent, d := block.Body.Content(schemaBody)
	if err := errors.EvalDiagnostics(d); err != nil {
		return nil, err
	}
	if len(bodyContent.Attributes) > 0 {
		return nil, errors.ThrowUnsupportedArguments(blockBody)
	}
	if len(bodyContent.Blocks) == 0 {
		return nil, errors.IncorrectUsage("blcok '%s' must contain one action at least", blockBody)
	}
	actions, err := handler.DecodePlugins(bodyContent.Blocks)
	if err != nil {
		return nil, err
	}
	return &Body{
		actions: actions,
	}, nil
}
