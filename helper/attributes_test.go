package helper

import (
	"strings"
	"testing"

	"github.com/hashicorp/hcl/v2"
	"github.com/stretchr/testify/assert"
	"github.com/wesovilabs/orion/testutil"
	"github.com/zclconf/go-cty/cty"
)

var testEvalCtx = &hcl.EvalContext{
	Variables: map[string]cty.Value{
		"firstname": cty.StringVal("Jane"),
		"lastname":  cty.StringVal("Doe"),
		"myNumber":  cty.NumberIntVal(56),
		"people": cty.TupleVal([]cty.Value{
			cty.ObjectVal(map[string]cty.Value{
				"firstname": cty.StringVal("John"),
				"age":       cty.NumberIntVal(35),
			}),
			cty.ObjectVal(map[string]cty.Value{
				"firstname": cty.StringVal("Jane"),
				"age":       cty.NumberIntVal(34),
			}),
			cty.ObjectVal(map[string]cty.Value{
				"firstname": cty.StringVal("Time"),
				"age":       cty.NumberIntVal(2),
			}),
		}),
	},
}

const testAttributesPath = "testdata/attributes.hcl"

var values = map[string]cty.Value{
	"attribute_text":    cty.StringVal("John"),
	"attribute_text2":   cty.StringVal("Jane"),
	"attribute_text3":   cty.StringVal("Jane Doe"),
	"attribute_number":  cty.NumberIntVal(2),
	"attribute_number2": cty.NumberIntVal(8),
	"attribute_number3": cty.NumberIntVal(56),
	"attribute_map": cty.ObjectVal(map[string]cty.Value{
		"firstname": cty.StringVal("Jane"),
		"age":       cty.NumberIntVal(56),
	}),
	"attribute_array": cty.TupleVal([]cty.Value{
		cty.StringVal("Jane"),
		cty.StringVal("a"),
		cty.NumberIntVal(5),
		cty.NumberIntVal(56),
	}),
}

func TestEvalAttribute(t *testing.T) {
	attrs, err := testutil.GetAttributes(testAttributesPath)
	assert.Nil(t, err)
	assert.NotNil(t, attrs)
	for index := range attrs {
		attr := attrs[index]
		val, err := EvalAttribute(testEvalCtx, attr)
		if strings.HasPrefix(attr.Name, "error") {
			assert.NotNil(t, err)
			continue
		}
		assert.Nil(t, err)
		assert.True(t, values[attr.Name].Equals(val).True())
	}
}

func TestEvalUnorderedExpression(t *testing.T) {
	text := `
		attr = 4
		attr2 = attr3
		attr3 = 4 * attr4
		attr4 = 5
	`
	expectations := map[string]cty.Value{
		"attr":  cty.NumberIntVal(4),
		"attr2": cty.NumberIntVal(20),
		"attr3": cty.NumberIntVal(20),
		"attr4": cty.NumberIntVal(5),
	}
	attrs, err := testutil.GetAttributesFromText(text)
	assert.Nil(t, err)
	expressions := testutil.MapStringAttributeToStringExpression(attrs)
	err = EvalUnorderedExpression(testEvalCtx, expressions)
	assert.Nil(t, err)
	for name, value := range expectations {
		assert.True(t, testEvalCtx.Variables[name].Equals(value).True())
	}
}

func TestEvalUnorderedExpressionError(t *testing.T) {
	text := `
		attr = unknown
	`
	attrs, err := testutil.GetAttributesFromText(text)
	assert.Nil(t, err)
	expressions := testutil.MapStringAttributeToStringExpression(attrs)
	err = EvalUnorderedExpression(testEvalCtx, expressions)
	assert.NotNil(t, err)
}

func TestEvaluateArrayItemExpression(t *testing.T) {
	text := `
		people = {
			fullName = "Jane Doe"
		}
	`
	attrs, err := testutil.GetAttributesFromText(text)
	assert.Nil(t, err)

	err = EvaluateArrayItemExpression(testEvalCtx, "people", 1, attrs["people"].Expr)
	element := testEvalCtx.Variables["people"].AsValueSlice()[1].AsValueMap()
	assert.True(t, element["fullName"].Equals(cty.StringVal("Jane Doe")).True())

	err = EvaluateArrayItemExpression(testEvalCtx, "people", 3, attrs["people"].Expr)
	element = testEvalCtx.Variables["people"].AsValueSlice()[3].AsValueMap()
	assert.True(t, element["fullName"].Equals(cty.StringVal("Jane Doe")).True())
}

func TestEvaluateArrayItemExpressionErrror(t *testing.T) {
	text := `
		people = {
			fullName = "Jane Doe"
		}
	`
	attrs, err := testutil.GetAttributesFromText(text)
	assert.Nil(t, err)
	err = EvaluateArrayItemExpression(testEvalCtx, "nilVal", 1, attrs["people"].Expr)
	assert.NotNil(t, err)
}
