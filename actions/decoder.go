package actions

import (
	"github.com/hashicorp/hcl/v2"
	"github.com/wesovilabs-tools/orion/internal/errors"
)

// Decoder interface to be implemented bu the plugin Decoders.
type Decoder interface {
	BlockHeaderSchema() hcl.BlockHeaderSchema
	DecodeBlock(block *hcl.Block) (Action, errors.Error)
}
