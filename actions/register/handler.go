// Package plugins contain the types and method to deal with plugins
package register

import (
	"sync"

	"github.com/hashicorp/hcl/v2"
	"github.com/wesovilabs-tools/orion/actions"
	"github.com/wesovilabs-tools/orion/internal/errors"
)

var handler Handler

var once sync.Once

// Get return an instance of the Handler interface.
func GetHandler() Handler {
	once.Do(func() {
		handler = new(manager)
	})
	return handler
}

// Handler interface to obtain the plugin to be performed.
type Handler interface {
	DecodePlugin(block *hcl.Block) (actions.Action, errors.Error)
	DecodePlugins(blocks hcl.Blocks) (actions.Actions, errors.Error)
	GetBlocksSpec() []hcl.BlockHeaderSchema
}

type manager struct{}

func (m *manager) DecodePlugin(block *hcl.Block) (actions.Action, errors.Error) {
	dec, ok := decoders[block.Type]
	if !ok {
		return nil, errors.ThrowUnsupportedBlock("", block.Type)
	}
	return dec.DecodeBlock(block)
}

func (m *manager) DecodePlugins(blocks hcl.Blocks) (actions.Actions, errors.Error) {
	actions := make(actions.Actions, len(blocks))
	for i, block := range blocks {
		action, err := m.DecodePlugin(block)
		if err != nil {
			return nil, err
		}
		action.SetKind(block.Type)
		actions[i] = action
	}
	return actions, nil
}

func (m *manager) GetBlocksSpec() []hcl.BlockHeaderSchema {
	return schemaPlugins.Blocks
}
