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
	"github.com/wesovilabs/orion/actions/sleep"
)

var decoders = map[string]actions.Decoder{
	assert.BlockAssert: new(assert.Decoder),
	block.BlockBlock:   new(block.Decoder),
	call.BlockCall:     new(call.Decoder),
	http.BlockHTTP:     new(http.Decoder),
	mongo.BlockMongo:   new(mongo.Decoder),
	pprint.BlockPrint:  new(pprint.Decoder),
	set.BlockSet:       new(set.Decoder),
	sleep.BlockSleep:   new(sleep.Decoder),
}

var schemaPlugins = &hcl.BodySchema{
	Attributes: nil,
	Blocks: []hcl.BlockHeaderSchema{
		decoders[assert.BlockAssert].BlockHeaderSchema(),
		decoders[block.BlockBlock].BlockHeaderSchema(),
		decoders[call.BlockCall].BlockHeaderSchema(),
		decoders[http.BlockHTTP].BlockHeaderSchema(),
		decoders[mongo.BlockMongo].BlockHeaderSchema(),
		decoders[pprint.BlockPrint].BlockHeaderSchema(),
		decoders[set.BlockSet].BlockHeaderSchema(),
		decoders[sleep.BlockSleep].BlockHeaderSchema(),
	},
}
