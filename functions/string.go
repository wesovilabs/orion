package functions

import (
	"strings"

	"github.com/zclconf/go-cty/cty"
	"github.com/zclconf/go-cty/cty/function"
)

const opSplit = "split"

var split = function.New(&function.Spec{
	VarParam: nil,
	Params: []function.Parameter{
		{Type: cty.String},
		{Type: cty.String},
	},
	Type: func(args []cty.Value) (cty.Type, error) {
		return cty.List(cty.String), checkArgumentType(opSplit, args, cty.String, cty.String)
	},
	Impl: func(args []cty.Value, retType cty.Type) (cty.Value, error) {
		s := args[0].AsString()
		sep := args[1].AsString()
		res := strings.Split(s, sep)
		items := make([]cty.Value, len(res))
		for index := range res {
			items[index] = cty.StringVal(res[index])
		}
		return cty.ListVal(items), nil
	},
})
