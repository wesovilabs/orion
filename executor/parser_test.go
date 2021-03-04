package executor

import (
	"github.com/stretchr/testify/assert"
	"github.com/zclconf/go-cty/cty"
	"testing"
)

func TestParseVariables(t *testing.T) {
	vars,err:=ParseVariables("testdata/variables.hcl")
	assert.Nil(t,err)
	assert.NotNil(t,vars)
	assert.True(t,vars["firstname"].RawEquals(cty.StringVal("John")))
	assert.True(t,vars["age"].RawEquals(cty.NumberIntVal(25)))
	assert.True(t,vars["person"].AsValueMap()["firstname"].RawEquals(cty.StringVal("Louis")))
	assert.True(t,vars["person"].AsValueMap()["age"].RawEquals(cty.NumberIntVal(20)))
	vars,err=ParseVariables("testdata/variables.error.hcl")
	assert.Nil(t,vars)
	assert.NotNil(t,err)
	assert.Equal(t,"testdata/variables.error.hcl:4,15-25: Variables not allowed; Variables may not be used here.",err.Message())
	vars,err=ParseVariables("unknown-file")
	assert.Nil(t,vars)
	assert.NotNil(t,err)
	assert.Equal(t,"file cannot be read",err.Message())
	vars,err=ParseVariables("testdata/variables.error2.hcl")
	assert.Nil(t,vars)
	assert.NotNil(t,err)
	assert.Equal(t,"testdata/variables.error2.hcl:1,1-8: Argument or block definition required; An argument or block definition is required here. To set an argument, use the equals sign \"=\" to introduce the argument value.",err.Message())
	vars,err=ParseVariables("testdata/variables.error3.hcl")
	assert.Nil(t,vars)
	assert.NotNil(t,err)
	assert.Equal(t,"testdata/variables.error3.hcl:2,1-7: Unexpected \"person\" block; Blocks are not allowed here.",err.Message())
}

func TestExecutor_Parse(t *testing.T) {
	exec:=New().(*executor)
	feature,err:=exec.Parse("testdata/feature.hcl")
	assert.Nil(t,err)
	assert.NotNil(t,feature)
}
