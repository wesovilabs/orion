package functions

import (
	"github.com/zclconf/go-cty/cty"
	"github.com/zclconf/go-cty/cty/function"

	"github.com/wesovilabs/orion/internal/errors"
)

const opEval = "eval"

var eval = function.New(&function.Spec{
	VarParam: nil,
	Params: []function.Parameter{
		{
			Type: cty.Bool,
		},
	},
	Type: func(args []cty.Value) (t cty.Type, err error) {
		t = cty.Bool
		if len(args) != 1 {
			err = errors.IncorrectUsage("`eval` function must retrieve 2 arguments", nil)
			return
		}
		return
	},
	Impl: func(args []cty.Value, retType cty.Type) (cty.Value, error) {
		return args[0], nil
	},
})
