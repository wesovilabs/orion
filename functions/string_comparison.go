package functions

import (
	"strings"

	"github.com/zclconf/go-cty/cty"
	"github.com/zclconf/go-cty/cty/function"
)

const opEqIgnoreCase = "eqIgnoreCase"

var eqIgnoreCase = function.New(&function.Spec{
	VarParam: nil,
	Params: []function.Parameter{
		{
			Type: cty.String,
		},
		{
			Type: cty.String,
		},
	},
	Type: func(args []cty.Value) (cty.Type, error) {
		return cty.Bool, checkArgumentType(opEqIgnoreCase, args, cty.String, cty.String)
	},
	Impl: func(args []cty.Value, retType cty.Type) (cty.Value, error) {
		if strings.EqualFold(args[0].AsString(), args[1].AsString()) {
			return cty.True, nil
		}
		return cty.False, nil
	},
})
