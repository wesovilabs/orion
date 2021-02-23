package helper

import (
	"github.com/stretchr/testify/assert"
	"github.com/wesovilabs/orion/testutil"
	"testing"
	"time"
)

func TestGetExpressionValueAsDuration(t *testing.T) {
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
		varDuration = "2s"
		varDuration2 = "2m"
		varDuration3 = "2h"
		varDuration4 = "2h2m 2s"
	`
	defDuration:=3*time.Second
	attrs, err := testutil.GetAttributesFromText(text)
	assert.Nil(t, err)
	value, err := GetExpressionValueAsDuration(testEvalCtx, attrs["varNumber"].Expr, &defDuration)
	assert.NotNil(t, err)
	assert.Nil(t, value)
	value, err = GetExpressionValueAsDuration(testEvalCtx, attrs["varString"].Expr, &defDuration)
	assert.NotNil(t, err)
	assert.Nil(t, value)
	value, err = GetExpressionValueAsDuration(testEvalCtx, attrs["varBool"].Expr, &defDuration)
	assert.NotNil(t, err)
	assert.Nil(t, value)
	value, err = GetExpressionValueAsDuration(testEvalCtx, nil, &defDuration)
	assert.Nil(t, err)
	assert.NotNil(t, value)
	assert.EqualValues(t, &defDuration, value)
	value, err = GetExpressionValueAsDuration(testEvalCtx, attrs["varSlice"].Expr, &defDuration)
	assert.NotNil(t, err)
	assert.Nil(t, value)
	value, err = GetExpressionValueAsDuration(testEvalCtx, attrs["varMap"].Expr, &defDuration)
	assert.NotNil(t, err)
	assert.Nil(t, value)
	value, err = GetExpressionValueAsDuration(testEvalCtx, attrs["varMap2"].Expr, &defDuration)
	assert.NotNil(t, err)
	assert.Nil(t, value)
	value, err = GetExpressionValueAsDuration(testEvalCtx, attrs["varDuration"].Expr, &defDuration)
	assert.Nil(t, err)
	assert.NotNil(t, value)
	assert.EqualValues(t, 2*time.Second, *value)

	value, err = GetExpressionValueAsDuration(testEvalCtx, attrs["varDuration2"].Expr, &defDuration)
	assert.Nil(t, err)
	assert.NotNil(t, value)
	assert.EqualValues(t, 2*time.Minute, *value)

	value, err = GetExpressionValueAsDuration(testEvalCtx, attrs["varDuration3"].Expr, &defDuration)
	assert.Nil(t, err)
	assert.NotNil(t, value)
	assert.EqualValues(t, 2*time.Hour, *value)

	value, err = GetExpressionValueAsDuration(testEvalCtx, attrs["varDuration4"].Expr, &defDuration)
	assert.Nil(t, err)
	assert.NotNil(t, value)
	assert.EqualValues(t,2*time.Hour + 2*time.Minute + 2*time.Second, *value)
}
