package dsl

import (
	"github.com/hashicorp/hcl/v2"

	"github.com/wesovilabs/orion/internal/errors"
)

const blockFunc = "func"

var schemaFunc = &hcl.BodySchema{
	Blocks: []hcl.BlockHeaderSchema{
		{Type: blockArg, LabelNames: []string{labelName}},
		{Type: blockBody},
		{Type: blockReturn},
	},
}

type Functions map[string]*Func

type Func struct {
	name string
	args Args
	ret  *Return
	body *Body
}

func decodeFunc(block *hcl.Block) (*Func, errors.Error) {
	var err errors.Error
	if len(block.Labels) != 1 {
		return nil, errors.ThrowMissingRequiredLabel(labelName)
	}
	function := &Func{
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
		case blockArg:
			function.args = make(Args, len(blocks))
			for i := range blocks {
				if function.args[i], err = DecodeArg(blocks[i]); err != nil {
					return nil, err
				}
			}
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
