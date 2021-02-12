package functions

import (
	"github.com/zclconf/go-cty/cty"
	"github.com/zclconf/go-cty/cty/function"

	"github.com/wesovilabs/orion/helper"
	"github.com/wesovilabs/orion/internal/errors"
)

const opLen = "len"

var lenOp = function.New(&function.Spec{
	VarParam: nil,
	Params: []function.Parameter{
		{
			Type: cty.DynamicPseudoType,
		},
	},
	Type: func(args []cty.Value) (t cty.Type, err error) {
		t = cty.Number
		if len(args) != 1 {
			err = errors.IncorrectUsage("`len` function must retrieve 1 arguments")
			return
		}
		if !helper.IsSlice(args[0]) {
			err = errors.IncorrectUsage("argument must be a slice")
		}
		return
	},
	Impl: func(args []cty.Value, retType cty.Type) (cty.Value, error) {
		listLen := len(args[0].AsValueSlice())
		return cty.NumberFloatVal(float64(listLen)), nil
	},
})
