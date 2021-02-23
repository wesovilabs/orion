package block

import (
	"github.com/stretchr/testify/assert"
	"github.com/wesovilabs/orion/testutil"
	"testing"
)

func TestDecoder_DecodeBlock(t *testing.T) {
	var input = `
	block {
	}
	`

	d := &Decoder{}

	content := testutil.GetStringAsBodyContent(input, BlockBlock, []string{})
	assert.NotNil(t, content)

	block := content.Blocks[0]
	action, err := d.DecodeBlock(block)
	assert.NotNil(t, err)
	assert.Nil(t, action)

	input = `
	block {
		print {
			msg = "Hi"
		}
	}
	`
	content = testutil.GetStringAsBodyContent(input, BlockBlock, []string{})
	assert.NotNil(t, content)

	block = content.Blocks[0]
	action, err = d.DecodeBlock(block)
	assert.Nil(t, err)
	assert.NotNil(t, action)
	assert.Equal(t, BlockBlock, action.String())
	actionBlock, ok := action.(*Block)
	assert.True(t, ok)
	assert.NotNil(t, actionBlock)
	assert.Len(t, actionBlock.actions, 1)

}
