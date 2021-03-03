package helper

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wesovilabs/orion/testutil"
)

func TestGetExpressionValueAsString2(t *testing.T) {
	text := `
		varMap = {
			fullName = "Jane Doe"
		}
		varMap2 = {
			fullName = unknown
		}
		varString ="hi"
		invalid = unknown
		varNumber = 1+1
		varNumber2 = 3.00
		varBool =  1==1
		varSlice = [1,2,3,4,5]
	`
	attrs, err := testutil.GetAttributesFromText(text)
	assert.Nil(t, err)
	value, err := GetExpressionValueAsString(testEvalCtx, attrs["varNumber"].Expr, "")
	assert.Nil(t, err)
	assert.NotNil(t, value)
	assert.EqualValues(t, "2", value)
	value, err = GetExpressionValueAsString(testEvalCtx, attrs["varString"].Expr, "")
	assert.Nil(t, err)
	assert.NotNil(t, value)
	assert.EqualValues(t, "hi", value)
	value, err = GetExpressionValueAsString(testEvalCtx, attrs["varBool"].Expr, "")
	assert.Nil(t, err)
	assert.NotNil(t, value)
	assert.EqualValues(t, "true", value)
	value, err = GetExpressionValueAsString(testEvalCtx, nil, "def")
	assert.Nil(t, err)
	assert.NotNil(t, value)
	assert.EqualValues(t, "def", value)
	value, err = GetExpressionValueAsString(testEvalCtx, attrs["varSlice"].Expr, "")
	assert.Nil(t, err)
	assert.NotNil(t, value)
	assert.EqualValues(t, "1, 2, 3, 4, 5", value)
	value, err = GetExpressionValueAsString(testEvalCtx, attrs["varMap"].Expr, "")
	assert.Nil(t, err)
	assert.NotNil(t, value)
	assert.EqualValues(t, "fullName:[JaneDoe]", strings.ReplaceAll(value, " ", ""))
	value, err = GetExpressionValueAsString(testEvalCtx, attrs["varMap2"].Expr, "def")
	assert.NotNil(t, err)
	assert.NotNil(t, value)
	assert.EqualValues(t, "", value)
}

func TestAttributeToStringWithoutContext(t *testing.T) {
	text := `
		varMap = {
			fullName = "Jane Doe"
		}
		varMap2 = {
			fullName = unknown
		}
		varString ="hi"
		invalid = unknown
		varNumber = 1+1
		varNumber2 = 3.00
		varBool =  1==1
		varSlice = [1,2,3,4,5]
	`
	attrs, err := testutil.GetAttributesFromText(text)
	assert.Nil(t, err)
	value, err := AttributeToStringWithoutContext(attrs["varMap"])
	assert.NotNil(t, err)
	assert.NotNil(t, value)
	assert.EqualValues(t, "", value)
	value, err = AttributeToStringWithoutContext(attrs["varNumber2"])
	assert.NotNil(t, err)
	assert.NotNil(t, value)
	assert.EqualValues(t, "", value)
	value, err = AttributeToStringWithoutContext(attrs["invalid"])
	assert.NotNil(t, err)
	assert.NotNil(t, value)
	assert.EqualValues(t, "", value)
	value, err = AttributeToStringWithoutContext(attrs["varString"])
	assert.Nil(t, err)
	assert.NotNil(t, value)
	assert.EqualValues(t, "hi", value)
}
