package set

import (
	"github.com/hashicorp/hcl/v2"
	log "github.com/sirupsen/logrus"
	"github.com/wesovilabs-tools/orion/actions"
	"github.com/wesovilabs-tools/orion/internal/errors"
)

const (
	// BlockSet identifier of plugin.
	BlockSet = "set"
	// LabelName label required by block set.
	LabelName = "name"
	argValue  = "value"
	argIndex  = "arrayIndex"
)

var schemaSet = &hcl.BodySchema{
	Attributes: append(actions.BaseArguments, []hcl.AttributeSchema{
		{Name: argValue, Required: true},
		{Name: argIndex, Required: false},
	}...),
}

// Decoder implement interface Decoder.
type Decoder struct{}

// BlockHeaderSchema return the header schema for the plugin.
func (dec *Decoder) BlockHeaderSchema() hcl.BlockHeaderSchema {
	return hcl.BlockHeaderSchema{
		Type:       BlockSet,
		LabelNames: []string{LabelName},
	}
}

// DecodeBlock inherited method from interface Decoder.
func (dec *Decoder) DecodeBlock(block *hcl.Block) (actions.Action, errors.Error) {
	log.Tracef("it starts decoding of block %s", BlockSet)
	bodyContent, d := block.Body.Content(schemaSet)
	if err := errors.EvalDiagnostics(d); err != nil {
		return nil, err
	}
	if len(block.Labels) != 1 {
		return nil, errors.ThrowMissingRequiredLabel(BlockSet)
	}
	set := &Set{
		name: block.Labels[0],
		Base: &actions.Base{},
	}
	if err := set.populateAttributes(bodyContent.Attributes); err != nil {
		return nil, err
	}
	if len(bodyContent.Blocks) > 0 {
		return nil, errors.ThrowBlocksAreNotPermitted(BlockSet)
	}
	return set, nil
}
