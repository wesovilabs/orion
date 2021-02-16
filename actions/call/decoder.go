package call

import (
	"github.com/hashicorp/hcl/v2"
	log "github.com/sirupsen/logrus"
	"github.com/wesovilabs/orion/actions"
	"github.com/wesovilabs/orion/internal/errors"
)

const (
	// BlockCall identifier of action.
	BlockCall = "call"
	blockWith = "with"
	labelName = "name"
	argAs     = "as"
)

var schemaCall = &hcl.BodySchema{
	Attributes: append(actions.BaseArguments, []hcl.AttributeSchema{
		{Name: argAs, Required: false},
	}...),
	Blocks: []hcl.BlockHeaderSchema{
		{Type: blockWith, LabelNames: nil},
	},
}

// Decoder implement interface Decoder.
type Decoder struct{}

// BlockHeaderSchema return the header schema for the plugin.
func (dec *Decoder) BlockHeaderSchema() hcl.BlockHeaderSchema {
	return hcl.BlockHeaderSchema{
		Type:       BlockCall,
		LabelNames: []string{labelName},
	}
}

// DecodeBlock inherited method from interface Decoder.
func (dec *Decoder) DecodeBlock(block *hcl.Block) (actions.Action, errors.Error) {
	log.Tracef("it starts decoding of block %s", BlockCall)
	bodyContent, d := block.Body.Content(schemaCall)
	if err := errors.EvalDiagnostics(d); err != nil {
		return nil, err
	}
	if len(block.Labels) != 1 {
		return nil, errors.ThrowMissingRequiredLabel(BlockCall)
	}
	call := &Call{
		name: block.Labels[0],
		Base: &actions.Base{},
	}
	if err := call.populateAttributes(bodyContent.Attributes); err != nil {
		return nil, err
	}
	if err := call.populateBlocks(bodyContent.Blocks); err != nil {
		return nil, err
	}
	return call, nil
}
