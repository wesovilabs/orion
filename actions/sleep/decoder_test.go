package sleep

import (
	"path"
	"testing"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclsyntax"
	"github.com/stretchr/testify/assert"
	"github.com/wesovilabs/orion/actions/shared"
)

var assertDecoderPath = path.Join(
	featuresDir, featureScenario,
)

var assertDecoderBlocks = []*Sleep{
	{
		duration: &hclsyntax.TemplateExpr{
			SrcRange: hcl.Range{
				Filename: assertDecoderPath,
				Start:    shared.CreatePos(2, 3, 10),
				End:      shared.CreatePos(2, 20, 27),
			},
		},
	},
}

func TestDecoder_DecodeBlock(t *testing.T) {
	content, err := shared.GetBodyContent(assertDecoderPath, BlockSleep, []string{})
	assert.Nil(t, err)
	assert.NotNil(t, content)

	assert.Len(t, content.Blocks, len(assertDecoderBlocks))
	for index := range content.Blocks {
		block := content.Blocks[index]
		expectedBlock := assertDecoderBlocks[index]
		assertBlockAssert(t, block, expectedBlock)
	}
}

func assertBlockAssert(t *testing.T, block *hcl.Block, s *Sleep) {
	assert.Len(t, block.Labels, 0)
	assert.Equal(t, BlockSleep, block.Type)
	c, d := block.Body.Content(schemaSleep)
	assert.Nil(t, d)
	assert.NotNil(t, c)
	for name := range c.Attributes {
		attribute := c.Attributes[name]
		switch name {
		case AttributeDuration:
			assert.NotNil(t, s.duration)
			assert.Equal(t, s.duration.Range().Start, attribute.Range.Start)
			assert.Equal(t, s.duration.Range().End, attribute.Range.End)
		default:
			assert.Failf(t, "error", "unexpected attribute %s", name)
		}
	}
}
