package dsl

import (
	"fmt"
	"testing"

	"github.com/hashicorp/hcl/v2"
	"github.com/stretchr/testify/assert"
	"github.com/wesovilabs/orion/internal/oriontest"
)

var expectedBodiesDecodeErrs = map[string]struct{}{
	"testdata/body.hcl:6,3-7: Unsupported argument; An argument named \"name\" is not expected here.": {},
	"blcok 'body' must contain one action at least":                                                   {},
}

var expectedBodiesActions = []int{0, 0, 1, 2}

func TestBody_Execute(t *testing.T) {
	blocks := oriontest.ParseHCL("testdata/body.hcl", &hcl.BodySchema{
		Blocks: []hcl.BlockHeaderSchema{
			{Type: blockBody},
		},
	})
	for index := range blocks {
		block := blocks[index]
		body, err := decodeBody(block)
		if err != nil {
			fmt.Println(err.Message())
			assert.Contains(t, expectedBodiesDecodeErrs, err.Message())
			continue
		}
		assert.NotNil(t, body)
		assert.Len(t, body.actions, expectedBodiesActions[index])
	}
}
