package block

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wesovilabs/orion/testutil"
)

var testDecoder = &Decoder{}

func assertDecodeBlockError(t *testing.T, input string) {
	content := testutil.GetStringAsBodyContent(input, BlockBlock, []string{})
	assert.NotNil(t, content)
	block := content.Blocks[0]
	action, err := testDecoder.DecodeBlock(block)
	assert.NotNil(t, err)
	assert.Nil(t, action)
}

func TestDecoder_DecodeBlock(t *testing.T) {
	input := `
	block {
		print {
			msg = "Hi"
		}
	}
	`
	content := testutil.GetStringAsBodyContent(input, BlockBlock, []string{})
	assert.NotNil(t, content)

	block := content.Blocks[0]
	action, err := testDecoder.DecodeBlock(block)
	assert.Nil(t, err)
	assert.NotNil(t, action)
	assert.Equal(t, BlockBlock, action.String())
	actionBlock, ok := action.(*Block)
	assert.True(t, ok)
	assert.NotNil(t, actionBlock)
	assert.Len(t, actionBlock.actions, 1)

	input = `
	block {
	}`
	assertDecodeBlockError(t, input)
	input = `
	block {
		unsupported = 123
	}`
	assertDecodeBlockError(t, input)
	input = `
	block {
		name = unknown
	}`
	assertDecodeBlockError(t, input)
	input = `
	block {
		unknown {}
	}`
	assertDecodeBlockError(t, input)
}
