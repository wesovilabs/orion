package dsl

import (
	"testing"

	"github.com/hashicorp/hcl/v2"
	"github.com/stretchr/testify/assert"
	"github.com/wesovilabs/orion/context"
	"github.com/wesovilabs/orion/internal/errors"
	"github.com/wesovilabs/orion/internal/oriontest"
	"github.com/zclconf/go-cty/cty"
)

var expectedArgsDecodeErrs = map[string]struct{}{
	"variable lastname must be provided": {},
	"testdata/args.hcl:11,3-10: Unsupported argument; An argument named \"unknown\" is not expected here.": {},
	"variables are not permitted in attribute description":                                                 {},
	"variable gender must be provided":                                                                     {},
	"testdata/args.hcl:27,21-28: Unknown variable; There is no variable named \"unknown\".":                {},
}

var expectedArgs = map[string]struct {
	description string
	value       cty.Value
	err         errors.Error
}{
	"firstname": {
		description: "firstname of person",
		value:       cty.StringVal("John"),
	},
	"age": {
		value: cty.NumberIntVal(6),
	},
	"elements": {
		value: cty.ListVal([]cty.Value{
			cty.StringVal("d"), cty.StringVal("e"), cty.StringVal("f"),
		}),
	},
	"city": {
		value: cty.StringVal("London"),
	},
	"gender": {
		value: cty.StringVal("${unknown}"),
	},
}

var orionCtx = context.New(map[string]cty.Value{
	"elements": cty.ListVal([]cty.Value{
		cty.StringVal("d"), cty.StringVal("e"), cty.StringVal("f"),
	}),
	"city": cty.StringVal("London"),
})

func TestArg_Execute(t *testing.T) {
	blocks := oriontest.ParseHCL("testdata/args.hcl", &hcl.BodySchema{
		Blocks: []hcl.BlockHeaderSchema{
			{Type: blockArg, LabelNames: []string{labelName}},
		},
	})
	for index := range blocks {
		block := blocks[index]
		arg, err := decodeArg(block)
		if err != nil {
			assert.Contains(t, expectedArgsDecodeErrs, err.Message())
			continue
		}
		assert.NotNil(t, arg)
		expected := expectedArgs[arg.name]
		err = arg.Execute(orionCtx)
		if err != nil {
			assert.Contains(t, expectedArgsDecodeErrs, err.Message())
			continue
		}
		assert.True(t, expected.value.Equals(orionCtx.EvalContext().Variables[arg.name]).True(), "unexpected value for arg with name '%s'.", arg.name)
	}
}
