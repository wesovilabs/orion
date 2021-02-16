package register

import (
	"github.com/hashicorp/hcl/v2"
	"github.com/wesovilabs/orion/actions"
	"github.com/wesovilabs/orion/actions/assert"
	"github.com/wesovilabs/orion/actions/block"
	"github.com/wesovilabs/orion/actions/call"
	"github.com/wesovilabs/orion/actions/http"
	"github.com/wesovilabs/orion/actions/mongo"
	pprint "github.com/wesovilabs/orion/actions/print"
	"github.com/wesovilabs/orion/actions/set"
)

var decoders = map[string]actions.Decoder{
	set.BlockSet:       new(set.Decoder),
	assert.BlockAssert: new(assert.Decoder),
	pprint.BlockPrint:  new(pprint.Decoder),
	http.BlockHTTP:     new(http.Decoder),
	mongo.BlockMongo:   new(mongo.Decoder),
	block.BlockBlock:   new(block.Decoder),
	call.BlockCall:     new(call.Decoder),
}

var schemaPlugins = &hcl.BodySchema{
	Attributes: nil,
	Blocks: []hcl.BlockHeaderSchema{
		decoders[set.BlockSet].BlockHeaderSchema(),
		decoders[assert.BlockAssert].BlockHeaderSchema(),
		decoders[pprint.BlockPrint].BlockHeaderSchema(),
		decoders[http.BlockHTTP].BlockHeaderSchema(),
		decoders[mongo.BlockMongo].BlockHeaderSchema(),
		decoders[block.BlockBlock].BlockHeaderSchema(),
		decoders[call.BlockCall].BlockHeaderSchema(),
	},
}
