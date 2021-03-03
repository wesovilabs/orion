package helper

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wesovilabs/orion/testutil"
)

func TestGetExpressionValueAsBool(t *testing.T) {
	text := `
		people = {
			fullName = "Jane Doe"
		}
		invalid = unknown
		varFalse = false
		varFalse2 = 1==2
		varTrue = true
		varTrue2 = 1==1
	`
	attrs, err := testutil.GetAttributesFromText(text)
	assert.Nil(t, err)
	_, err = GetExpressionValueAsBool(testEvalCtx, attrs["invalid"].Expr, true)
	assert.NotNil(t, err)
	v, err := GetExpressionValueAsBool(testEvalCtx, attrs["people"].Expr, true)
	assert.Nil(t, err)
	assert.True(t, v)
	v, err = GetExpressionValueAsBool(testEvalCtx, attrs["varFalse"].Expr, true)
	assert.Nil(t, err)
	assert.False(t, v)
	v, err = GetExpressionValueAsBool(testEvalCtx, attrs["varFalse2"].Expr, true)
	assert.Nil(t, err)
	assert.False(t, v)
	v, err = GetExpressionValueAsBool(testEvalCtx, attrs["varTrue"].Expr, true)
	assert.Nil(t, err)
	assert.True(t, v)
	v, err = GetExpressionValueAsBool(testEvalCtx, attrs["varTrue2"].Expr, true)
	assert.Nil(t, err)
	assert.True(t, v)
}
