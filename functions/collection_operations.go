package functions

import (
	"github.com/zclconf/go-cty/cty"
	"github.com/zclconf/go-cty/cty/function"

	"github.com/wesovilabs-tools/orion/helper"
)

const (
	opFirst = "first"
	opLast  = "last"
)

var (
	first = function.New(collectionFixedPosition(opFirst, func(slice []cty.Value) cty.Value {
		return slice[0]
	}))
	last = function.New(collectionFixedPosition(opLast, func(slice []cty.Value) cty.Value {
		return slice[len(slice)-1]
	}))
)

func collectionFixedPosition(operation string, fn func([]cty.Value) cty.Value) *function.Spec {
	return &function.Spec{
		VarParam: nil,
		Params: []function.Parameter{
			{Type: cty.List(cty.DynamicPseudoType)},
		},
		Type: func(args []cty.Value) (cty.Type, error) {
			var err error
			if len(args) > 1 {
				err = invalidArgs(operation, 1)
			}
			if !helper.IsSlice(args[0]) {
				err = invalidArgType(operation, 0, "slice")
			}
			return cty.DynamicPseudoType, err
		},
		Impl: func(args []cty.Value, retType cty.Type) (cty.Value, error) {
			slice := args[0].AsValueSlice()
			return fn(slice), nil
		},
	}
}
