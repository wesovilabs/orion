package dsl

import (
	"github.com/hashicorp/hcl/v2"
	"github.com/wesovilabs/orion/context"

	"github.com/wesovilabs/orion/internal/errors"
)

var schemaInput = &hcl.BodySchema{
	Blocks: []hcl.BlockHeaderSchema{
		{Type: blockArg, LabelNames: []string{labelName}},
	},
}

// Input represents the list of variables to be provided
type Input struct {
	args Args
}

// TotalArgs return the number of defined tags in the block
func (i *Input) TotalArgs() int {
	return len(i.args)
}

// AddArg add a new arg to the lis tof arguments
func (i *Input) AddArg(arg *Arg) {
	if i.args == nil {
		i.args = make(Args, 0)
	}
	i.args = append(i.args, arg)
}

func (i *Input) Execute(ctx context.FeatureContext) errors.Error {
	if len(i.args) == 0 {
		return errors.IncorrectUsage("'%s'  only can be declared if it contains one or more block '%s'", blockInput, blockArg)
	}
	for index := range i.args {
		if err := i.args[index].Execute(ctx); err != nil {
			return err
		}
	}
	return nil
}

// decodeInput decode block input
func decodeInput(b *hcl.Block) (*Input, errors.Error) {
	input := new(Input)
	bc, d := b.Body.Content(schemaInput)
	if err := errors.EvalDiagnostics(d); err != nil {
		return nil, err
	}
	if err := input.populateBlocks(bc.Blocks); err != nil {
		return nil, err
	}
	return input, nil
}

func (i *Input) populateBlocks(blocks hcl.Blocks) errors.Error {
	for index := range blocks {
		block := blocks[index]
		switch block.Type {
		case blockArg:
			arg, err := DecodeArg(block)
			if err != nil {
				return err
			}
			i.AddArg(arg)
		default:
			return errors.ThrowUnsupportedBlock(blockInput, block.Type)
		}
	}
	return nil
}
