package context

import (
	"github.com/hashicorp/hcl/v2"
	"github.com/stretchr/testify/assert"
	"github.com/zclconf/go-cty/cty"
	"math/big"
	"testing"
)

func TestVariables_SetIndex(t *testing.T) {
	vars := createVariables().(*variables)
	vars.SetIndex(-1)
	assert.Equal(t, vars.index, -1)
	vars.SetIndex(2)
	assert.Equal(t, vars.index, 2)
}

func TestVariables_SetToContext(t *testing.T) {
	vars := createVariables().(*variables)
	vars.SetIndex(2)
	evalCtx:=&hcl.EvalContext{}
	rootVar:=evalCtx.Variables["_"]
	assert.Equal(t,rootVar,cty.NilVal)
	vars.SetToContext(evalCtx)
	assert.Len(t,evalCtx.Variables,1)
	rootVar=evalCtx.Variables["_"]
	index,acc:=rootVar.AsValueMap()["index"].AsBigFloat().Int64()
	assert.Equal(t,index,int64(2))
	assert.Equal(t,acc,big.Accuracy(0))
}
