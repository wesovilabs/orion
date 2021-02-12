package assert

import (
	"path"
	"testing"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclsyntax"
	"github.com/stretchr/testify/assert"
	"github.com/wesovilabs/orion/actions/shared"
)

var assertDecoderPath = path.Join(
	"testdata", "assert-decoder.hcl",
)

var assertDecoderBlocks = []*Assert{
	{
		assertion: &hclsyntax.TemplateExpr{
			SrcRange: hcl.Range{
				Filename: assertDecoderPath,
				Start:    shared.CreatePos(2, 3, 11),
				End:      shared.CreatePos(2, 19, 27),
			},
		},
	},
	{
		assertion: &hclsyntax.TemplateExpr{
			SrcRange: hcl.Range{
				Filename: assertDecoderPath,
				Start:    shared.CreatePos(6, 3, 42),
				End:      shared.CreatePos(6, 28, 67),
			},
		},
	},
}

func TestDecoder_DecodeBlock(t *testing.T) {

	content, err := shared.GetBodyContent(assertDecoderPath, BlockAssert, []string{})
	assert.Nil(t, err)
	assert.NotNil(t, content)

	assert.Len(t, content.Blocks, len(assertDecoderBlocks))
	for index := range content.Blocks {
		block := content.Blocks[index]
		expectedBlock := assertDecoderBlocks[index]
		assertBlockAssert(t, block, expectedBlock)
	}
}

func assertBlockAssert(t *testing.T, block *hcl.Block, a *Assert) {
	assert.Len(t, block.Labels, 0)
	assert.Equal(t, BlockAssert, block.Type)
	c, d := block.Body.Content(schemaAssert)
	assert.Nil(t, d)
	assert.NotNil(t, c)
	for name := range c.Attributes {
		attribute := c.Attributes[name]
		switch name {
		case argAssertion:
			assert.NotNil(t, a.assertion)
			assert.Equal(t, a.assertion.Range().Start, attribute.Range.Start)
			assert.Equal(t, a.assertion.Range().End, attribute.Range.End)
		default:
			assert.Failf(t, "error", "unexpected attribute %s", name)
		}
	}
}
