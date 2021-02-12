package context

import (
	"github.com/hashicorp/hcl/v2"
	"github.com/zclconf/go-cty/cty"
)

// Variables interface definition.
type Variables interface {
	SetIndex(int)
	SetToContext(ctx *hcl.EvalContext)
}

type variables struct {
	index int
}

func createVariables() Variables {
	v := new(variables)
	return v
}

// SetIndex establish index variable.
func (v *variables) SetIndex(index int) {
	v.index = index
}

// SetToContext set the variables into the context.
func (v *variables) SetToContext(ctx *hcl.EvalContext) {
	ctx.Variables["_"] = cty.MapVal(map[string]cty.Value{
		"index": cty.NumberIntVal(int64(v.index)),
	})
}
