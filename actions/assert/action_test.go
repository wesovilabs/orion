package assert

import (
	"path"
	"testing"

	"github.com/wesovilabs/orion/context"

	"github.com/stretchr/testify/assert"
	"github.com/wesovilabs/orion/actions/shared"
	"github.com/zclconf/go-cty/cty"
)

var decoder = new(Decoder)

var assertPluginResult = []bool{true, false, false, true, true, true, false, true, false, false}

func TestAssert_Assertion(t *testing.T) {
	content, err := shared.GetBodyContent(path.Join("testdata", "assert-evaluation.hcl"), BlockAssert, []string{})
	assert.Nil(t, err)
	assert.NotNil(t, content)
	ctx := context.New(map[string]cty.Value{
		"person": cty.ObjectVal(map[string]cty.Value{
			"lastname": cty.StringVal("Robot"),
			"age":      cty.NumberIntVal(int64(20)),
		}),
	}, nil)

	for index := range content.Blocks {
		expected := assertPluginResult[index]
		plugin, err := decoder.DecodeBlock(content.Blocks[index])
		assert.Nil(t, err)
		assert.NotNil(t, plugin)
		a, ok := plugin.(*Assert)
		assert.True(t, ok)
		assert.NotNil(t, a)
		assert.NotNil(t, a.Assertion())
		err = a.Execute(ctx)
		if expected {
			assert.Nil(t, err)
			continue
		}
		assert.NotNil(t, err)
	}
}
