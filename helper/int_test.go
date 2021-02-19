package helper

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wesovilabs/orion/testutil"
)

func TestGetExpressionValueAsInt(t *testing.T) {
	text := `
		people = {
			fullName = "Jane Doe"
		}
		invalid = unknown
		var1 = 1
		var2 = 1+1
		var3 = 3.00
	`
	attrs, err := testutil.GetAttributesFromText(text)
	assert.Nil(t, err)
	_, err = GetExpressionValueAsInt(testEvalCtx, attrs["invalid"].Expr, 100)
	assert.NotNil(t, err)
	v, err := GetExpressionValueAsInt(testEvalCtx, attrs["people"].Expr, 100)
	assert.Nil(t, err)
	assert.Equal(t, v, 100)
	v, err = GetExpressionValueAsInt(testEvalCtx, attrs["var1"].Expr, 100)
	assert.Nil(t, err)
	assert.Equal(t, v, 1)
	v, err = GetExpressionValueAsInt(testEvalCtx, attrs["var2"].Expr, 100)
	assert.Nil(t, err)
	assert.Equal(t, v, 2)
	v, err = GetExpressionValueAsInt(testEvalCtx, attrs["var3"].Expr, 100)
	assert.Nil(t, err)
	assert.Equal(t, v, 3)
}
