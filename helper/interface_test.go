package helper

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wesovilabs/orion/testutil"
)

func TestGetExpressionValueAsInterface(t *testing.T) {
	text := `
		varMap = {
			fullName = "Jane Doe"
		}
		varMap2 = {
			fullName = unknown
		}
		invalid = unknown
		varNumber = 1+1
		varNumber2 = 3.00
		varBool =  1==1
		varSlice = [1,2,3,4,5]
	`
	attrs, err := testutil.GetAttributesFromText(text)
	assert.Nil(t, err)

	var value interface{}
	value, err = GetExpressionValueAsInterface(testEvalCtx, attrs["varMap"].Expr, nil)
	assert.Nil(t, err)
	assert.NotNil(t, value)
	valMap, ok := value.(map[string]interface{})
	assert.True(t, ok)
	assert.Equal(t, "Jane Doe", valMap["fullName"])

	var valSlice []interface{}
	value, err = GetExpressionValueAsInterface(testEvalCtx, attrs["varSlice"].Expr, nil)
	assert.Nil(t, err)
	assert.NotNil(t, value)
	valSlice, ok = value.([]interface{})
	assert.True(t, ok)
	assert.Len(t, valSlice, 5)

	var valFloat float64
	value, err = GetExpressionValueAsInterface(testEvalCtx, attrs["varNumber"].Expr, nil)
	assert.Nil(t, err)
	assert.NotNil(t, value)
	valFloat, ok = value.(float64)
	assert.True(t, ok)
	assert.EqualValues(t, 2, valFloat)
	value, err = GetExpressionValueAsInterface(testEvalCtx, attrs["varNumber2"].Expr, nil)
	assert.Nil(t, err)
	assert.NotNil(t, value)
	valFloat, ok = value.(float64)
	assert.True(t, ok)
	assert.EqualValues(t, 3.00, valFloat)

	var valBool bool
	value, err = GetExpressionValueAsInterface(testEvalCtx, attrs["varBool"].Expr, nil)
	assert.Nil(t, err)
	assert.NotNil(t, value)
	valBool, ok = value.(bool)
	assert.True(t, ok)
	assert.EqualValues(t, true, valBool)

	value, err = GetExpressionValueAsInterface(testEvalCtx, attrs["invalid"].Expr, nil)
	assert.NotNil(t, err)
	assert.Nil(t, value)

	value, err = GetExpressionValueAsInterface(testEvalCtx, nil, nil)
	assert.Nil(t, err)
	assert.Nil(t, value)

	value, err = GetExpressionValueAsInterface(testEvalCtx, attrs["varMap2"].Expr, nil)
	assert.NotNil(t, err)
	assert.Nil(t, value)
}
