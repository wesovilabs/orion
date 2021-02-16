package dsl

import (
	"github.com/hashicorp/hcl/v2"
	"github.com/wesovilabs/orion/internal/errors"
)

const blockFunc = "func"

var schemaFunc = &hcl.BodySchema{
	Blocks: []hcl.BlockHeaderSchema{
		{Type: blockInput},
		{Type: blockBody},
		{Type: blockReturn},
	},
}

type Functions map[string]*Function

type Function struct {
	name  string
	input *Input
	ret   *Return
	body  *Body
}

func (f *Function) Input() *Input {
	return f.input
}

func (f *Function) Body() *Body {
	return f.body
}

func (f *Function) Return() *Return {
	return f.ret
}

func decodeFunc(block *hcl.Block) (*Function, errors.Error) {
	var err errors.Error
	if len(block.Labels) != 1 {
		return nil, errors.ThrowMissingRequiredLabel(labelName)
	}
	function := &Function{
		name: block.Labels[0],
	}
	bodyContent, d := block.Body.Content(schemaFunc)
	if err = errors.EvalDiagnostics(d); err != nil {
		return nil, err
	}
	if len(bodyContent.Attributes) > 0 {
		return nil, errors.ThrowUnsupportedArguments(blockFunc)
	}
	for name, blocks := range bodyContent.Blocks.ByType() {
		switch name {
		case blockInput:
			if len(blocks) > 1 {
				return nil, errors.ThrowsExceeddedNumberOfBlocks(blockBody, 1)
			}
			input, err := decodeInput(blocks[0])
			if err != nil {
				return nil, err
			}
			function.input = input
		case blockBody:
			if len(blocks) > 1 {
				return nil, errors.ThrowsExceeddedNumberOfBlocks(blockBody, 1)
			}
			if len(blocks) == 1 {
				if function.body, err = decodeBody(blocks[0]); err != nil {
					return nil, err
				}
			}
		case blockReturn:
			if len(blocks) > 1 {
				return nil, errors.ThrowsExceeddedNumberOfBlocks(blockReturn, 1)
			}
			if len(blocks) == 1 {
				if function.ret, err = decodeReturn(blocks[0]); err != nil {
					return nil, err
				}
			}
		default:
			return nil, errors.ThrowUnsupportedBlock(blockFunc, name)
		}
	}
	return function, nil
}
