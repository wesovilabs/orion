package dsl

import (
	"github.com/hashicorp/hcl/v2"
	log "github.com/sirupsen/logrus"
	"github.com/wesovilabs/orion/context"
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
	if len(block.Labels) != 1 {
		return nil, errors.ThrowMissingRequiredLabel(labelName)
	}
	function := &Function{
		name: block.Labels[0],
	}
	bodyContent, d := block.Body.Content(schemaFunc)
	if err := errors.EvalDiagnostics(d); err != nil {
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
				body, err := decodeBody(blocks[0])
				if err != nil {
					return nil, err
				}
				function.body = body
			}
		case blockReturn:
			if len(blocks) > 1 {
				return nil, errors.ThrowsExceeddedNumberOfBlocks(blockReturn, 1)
			}
			if len(blocks) == 1 {
				ret, err := decodeReturn(blocks[0])
				if err != nil {
					return nil, err
				}
				function.ret = ret
			}
		default:
			return nil, errors.ThrowUnsupportedBlock(blockFunc, name)
		}
	}
	return function, nil
}

// nolint
func (f *Function) runFunction(ctx context.OrionContext, out string) errors.Error {
	if input := f.Input(); input != nil {
		if err := input.Execute(ctx); err != nil {
			return err
		}
	}
	actions := f.Body().Actions()
	for index := range actions {
		action := actions[index]
		if action.ShouldExecute(ctx.EvalContext()) {
			if err := action.Execute(ctx); err != nil {
				return err
			}
			continue
		}
		log.Debugf("action %s is skipped!", action)
	}
	if f.Return() != nil {
		result, d := f.Return().Value().Value(ctx.EvalContext())
		if err := errors.EvalDiagnostics(d); err != nil {
			return err
		}
		ctx.EvalContext().Variables[out] = result
	}
	return nil
}

func (functions Functions) Append(newFunctions Functions) {
	for name, value := range newFunctions {
		functions[name] = value
	}
}
