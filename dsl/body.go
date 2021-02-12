package dsl

import (
	"github.com/hashicorp/hcl/v2"
	actions2 "github.com/wesovilabs/orion/actions"
	"github.com/wesovilabs/orion/internal/errors"
)

const blockBody = "body"

var schemaBody = &hcl.BodySchema{
	Blocks: handler.GetBlocksSpec(),
}

type Body struct {
	actions actions2.Actions
}

func decodeBody(block *hcl.Block) (*Body, errors.Error) {
	bodyContent, d := block.Body.Content(schemaBody)
	if err := errors.EvalDiagnostics(d); err != nil {
		return nil, err
	}
	if len(bodyContent.Attributes) > 0 {
		return nil, errors.ThrowUnsupportedArguments(blockBody)
	}
	actions, err := handler.DecodePlugins(bodyContent.Blocks)
	if err != nil {
		return nil, err
	}
	return &Body{
		actions: actions,
	}, nil
}
