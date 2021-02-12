package dsl

import (
	"github.com/hashicorp/hcl/v2"

	"github.com/wesovilabs-tools/orion/internal/errors"
)

const (
	blockReturn = "return"
	argValue    = "value"
)

var schemaReturn = &hcl.BodySchema{
	Attributes: []hcl.AttributeSchema{
		{Name: argValue, Required: true},
	},
}

type Return struct {
	value hcl.Expression
}

// DecodeBlock inherited method from interface Decoder.
func decodeReturn(block *hcl.Block) (*Return, errors.Error) {
	bodyContent, d := block.Body.Content(schemaReturn)
	if err := errors.EvalDiagnostics(d); err != nil {
		return nil, err
	}
	if len(bodyContent.Blocks) > 0 {
		return nil, errors.IncorrectUsage("blocks are not permited in '%s'", blockReturn)
	}
	ret := &Return{}
	for name, value := range bodyContent.Attributes {
		if name != argValue {
			return nil, errors.ThrowUnsupportedArgument(blockReturn, name)
		}
		ret.value = value.Expr
	}
	return ret, nil
}
