package block

import (
	"github.com/hashicorp/hcl/v2"
	log "github.com/sirupsen/logrus"
	"github.com/wesovilabs/orion/actions"
	"github.com/wesovilabs/orion/actions/assert"
	"github.com/wesovilabs/orion/actions/call"
	"github.com/wesovilabs/orion/actions/http"
	"github.com/wesovilabs/orion/actions/mongo"
	pprint "github.com/wesovilabs/orion/actions/print"
	"github.com/wesovilabs/orion/actions/set"
	"github.com/wesovilabs/orion/context"
	"github.com/wesovilabs/orion/internal/errors"
)

const (
	// BlockBlock identifier of plugin.
	BlockBlock = "block"
)

var decoders = map[string]actions.Decoder{
	set.BlockSet:       new(set.Decoder),
	assert.BlockAssert: new(assert.Decoder),
	pprint.BlockPrint:  new(pprint.Decoder),
	http.BlockHTTP:     new(http.Decoder),
	mongo.BlockMongo:   new(mongo.Decoder),
	call.BlockCall:     new(call.Decoder),
}

var schemaBlock = &hcl.BodySchema{
	Attributes: actions.BaseArguments,
	Blocks: []hcl.BlockHeaderSchema{
		decoders[assert.BlockAssert].BlockHeaderSchema(),
		decoders[set.BlockSet].BlockHeaderSchema(),
		decoders[pprint.BlockPrint].BlockHeaderSchema(),
		decoders[http.BlockHTTP].BlockHeaderSchema(),
		decoders[mongo.BlockMongo].BlockHeaderSchema(),
		decoders[call.BlockCall].BlockHeaderSchema(),
	},
}

// Set common block used to define variables.
type Block struct {
	*actions.Base
	actions actions.Actions
}

func (b *Block) populateAttributes(attrs hcl.Attributes) errors.Error {
	for name := range attrs {
		attribute := attrs[name]
		switch {
		case actions.IsPluginBaseArgument(name):
			if err := actions.SetBaseArgs(b, attribute); err != nil {
				return err
			}
		default:
			return errors.ThrowUnsupportedArgument(BlockBlock, name)
		}
	}
	return nil
}

func (b *Block) populateBlocks(blocks hcl.Blocks) errors.Error {
	for index := range blocks {
		block := blocks[index]
		dec, ok := decoders[block.Type]
		if !ok {
			return errors.ThrowUnsupportedBlock("", block.Type)
		}
		action, err := dec.DecodeBlock(block)
		if err != nil {
			return err
		}
		action.SetKind(block.Type)
		b.actions = append(b.actions, action)
	}
	return nil
}

// Execute method to run the plugin.
func (b *Block) Execute(ctx context.OrionContext) errors.Error {
	if len(b.actions) == 0 {
		return errors.IncorrectUsage("block '%s' cannot be empty. It must contain one action at least", BlockBlock)
	}
	return actions.Execute(ctx, b.Base, func(ctx context.OrionContext) errors.Error {
		for index := range b.actions {
			action := b.actions[index]
			if action.ShouldExecute(ctx.EvalContext()) {
				if err := action.Execute(ctx); err != nil {
					return err
				}
				continue
			}
			log.Debugf("action %s is skipped!", action)
		}
		return nil
	})
}

// Decoder implement interface Decoder.
type Decoder struct{}

// BlockHeaderSchema return the header schema for the plugin.
func (dec *Decoder) BlockHeaderSchema() hcl.BlockHeaderSchema {
	return hcl.BlockHeaderSchema{
		Type: BlockBlock,
	}
}

// DecodeBlock inherited method from interface Decoder.
func (dec *Decoder) DecodeBlock(block *hcl.Block) (actions.Action, errors.Error) {
	log.Tracef("it starts decoding of block %s", BlockBlock)
	bodyContent, d := block.Body.Content(schemaBlock)
	if err := errors.EvalDiagnostics(d); err != nil {
		return nil, err
	}
	b := &Block{
		actions: make(actions.Actions, 0),
		Base:    &actions.Base{},
	}
	b.SetKind(BlockBlock)
	if err := b.populateAttributes(bodyContent.Attributes); err != nil {
		return nil, err
	}
	if len(bodyContent.Blocks)==0{
		return nil, errors.IncorrectUsage("block of type `%s` cannot be empty ",BlockBlock)
	}
	if err := b.populateBlocks(bodyContent.Blocks); err != nil {
		return nil, err
	}
	log.Tracef("block %s is decoded successfully", BlockBlock)
	return b, nil
}
