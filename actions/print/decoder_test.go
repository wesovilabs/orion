package pprint

import (
	"path"
	"testing"

	"github.com/wesovilabs-tools/orion/actions/shared"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclsyntax"
	"github.com/stretchr/testify/assert"
)

const scenario3 = "scenario3.hcl"

func createPos(line, col, byte int) hcl.Pos {
	return hcl.Pos{
		Line:   line,
		Column: col,
		Byte:   byte,
	}
}

var inputPath = path.Join(featuresDir, scenario3)
var blocks = []*Print{
	{
		prefix: &hclsyntax.TemplateExpr{
			SrcRange: hcl.Range{
				Filename: inputPath,
				Start:    createPos(3, 3, 41),
				End:      createPos(3, 23, 61),
			},
		},
		msg: &hclsyntax.TemplateExpr{
			SrcRange: hcl.Range{
				Filename: inputPath,
				Start:    createPos(2, 3, 10),
				End:      createPos(2, 31, 38),
			},
		},
	},
	{
		msg: &hclsyntax.TemplateExpr{
			SrcRange: hcl.Range{
				Filename: inputPath,
				Start:    createPos(7, 3, 75),
				End:      createPos(7, 31, 103),
			},
		},
	},
	{
		prefix: &hclsyntax.TemplateExpr{
			SrcRange: hcl.Range{
				Filename: inputPath,
				Start:    createPos(11, 3, 117),
				End:      createPos(11, 23, 137),
			},
		},
		msg: &hclsyntax.TemplateExpr{
			SrcRange: hcl.Range{
				Filename: inputPath,
				Start:    createPos(12, 3, 140),
				End:      createPos(12, 31, 168),
			},
		},
		timestamp: &hclsyntax.TemplateExpr{
			SrcRange: hcl.Range{
				Filename: inputPath,
				Start:    createPos(13, 3, 171),
				End:      createPos(13, 21, 189),
			},
		},
		timestampFormat: &hclsyntax.TemplateExpr{
			SrcRange: hcl.Range{
				Filename: inputPath,
				Start:    createPos(14, 3, 192),
				End:      createPos(14, 33, 222),
			},
		},
		format: &hclsyntax.TemplateExpr{
			SrcRange: hcl.Range{
				Filename: inputPath,
				Start:    createPos(15, 3, 225),
				End:      createPos(15, 18, 240),
			},
		},
	},
}

func TestDecoder_DecodeBlock(t *testing.T) {
	content, err := shared.GetBodyContent(inputPath, BlockPrint, []string{})
	assert.Nil(t, err)
	assert.NotNil(t, content)

	assert.Len(t, content.Blocks, len(blocks))
	for index := range content.Blocks {
		block := content.Blocks[index]
		expectedBlock := blocks[index]
		assertBlockPrint(t, block, expectedBlock)
	}
}

func assertBlockPrint(t *testing.T, block *hcl.Block, print *Print) {
	assert.Len(t, block.Labels, 0)
	assert.Equal(t, BlockPrint, block.Type)
	c, d := block.Body.Content(schemaPrint)
	assert.Nil(t, d)
	assert.NotNil(t, c)
	for name := range c.Attributes {
		attribute := c.Attributes[name]
		switch name {
		case AttributeMsg:
			assert.NotNil(t, print.msg)
			assert.Equal(t, print.msg.Range().Start, attribute.Range.Start)
			assert.Equal(t, print.msg.Range().End, attribute.Range.End)
		case AttributePrefix:
			assert.NotNil(t, print.prefix)
			assert.Equal(t, print.prefix.Range().Start, attribute.Range.Start)
			assert.Equal(t, print.prefix.Range().End, attribute.Range.End)
		case AttributeFormat:
			assert.NotNil(t, print.format)
			assert.Equal(t, print.format.Range().Start, attribute.Range.Start)
			assert.Equal(t, print.format.Range().End, attribute.Range.End)
		case AttributeTimestamp:
			assert.NotNil(t, print.timestamp)
			assert.Equal(t, print.timestamp.Range().Start, attribute.Range.Start)
			assert.Equal(t, print.timestamp.Range().End, attribute.Range.End)
		case AttributeTimestampFormat:
			assert.NotNil(t, print.timestampFormat)
			assert.Equal(t, print.timestampFormat.Range().Start, attribute.Range.Start)
			assert.Equal(t, print.timestampFormat.Range().End, attribute.Range.End)
		default:
			assert.Failf(t, "error", "unexpected attribute %s", name)
		}
	}
}
