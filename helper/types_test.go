package helper

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zclconf/go-cty/cty"
)

func TestIsSlice(t *testing.T) {
	assert.False(t, IsSlice(cty.StringVal("f")))
	assert.False(t, IsSlice(cty.NumberIntVal(1)))
	assert.False(t, IsSlice(cty.NumberFloatVal(12.23)))
	assert.False(t, IsSlice(cty.BoolVal(true)))
	assert.False(t, IsSlice(cty.ObjectVal(map[string]cty.Value{})))
	assert.False(t, IsSlice(cty.ObjectVal(map[string]cty.Value{
		"firstname": cty.StringVal("Jane"),
	})))
	assert.True(t, IsSlice(cty.TupleVal([]cty.Value{})))
	assert.True(t, IsSlice(cty.TupleVal([]cty.Value{cty.StringVal("a"), cty.NumberFloatVal(23.12)})))
	assert.True(t, IsSlice(cty.ListVal([]cty.Value{cty.StringVal("a"), cty.StringVal("b")})))
}

func TestToStrictString(t *testing.T) {
	value, err := ToStrictString(cty.StringVal("a"))
	assert.Nil(t, err)
	assert.Equal(t, "a", value)

	value, err = ToStrictString(cty.BoolVal(false))
	assert.NotNil(t, err)
	assert.Equal(t, "", value)

	value, err = ToStrictString(cty.NumberFloatVal(22.2))
	assert.NotNil(t, err)
	assert.Equal(t, "", value)
}

func TestIsMap(t *testing.T) {
	assert.False(t, IsMap(cty.StringVal("f")))
	assert.False(t, IsMap(cty.NumberIntVal(1)))
	assert.False(t, IsMap(cty.NumberFloatVal(12.23)))
	assert.False(t, IsMap(cty.BoolVal(true)))
	assert.True(t, IsMap(cty.ObjectVal(map[string]cty.Value{})))
	assert.True(t, IsMap(cty.ObjectVal(map[string]cty.Value{
		"firstname": cty.StringVal("Jane"),
	})))
	assert.True(t, IsMap(cty.MapVal(map[string]cty.Value{
		"firstname": cty.StringVal("Jane"),
	})))
	assert.False(t, IsMap(cty.TupleVal([]cty.Value{})))
	assert.False(t, IsMap(cty.TupleVal([]cty.Value{cty.StringVal("a"), cty.NumberFloatVal(23.12)})))
	assert.False(t, IsMap(cty.ListVal([]cty.Value{cty.StringVal("a"), cty.StringVal("b")})))
}

func TestToValueMap(t *testing.T) {
	result := ToValueMap(map[string]interface{}{
		"firstname": "John",
		"age":       20,
		"male":      true,
		"salary":    800.34,
		"children": []map[string]interface{}{
			{
				"firstname": "Tim",
				"lastname":  "Doe",
			},
			{
				"firstname": "Loe",
				"lastname":  "Doe",
			},
		},
	})
	assert.NotNil(t, result)
	valueMap := result.AsValueMap()
	assert.True(t, valueMap["firstname"].RawEquals(cty.StringVal("John")))
	assert.True(t, valueMap["age"].RawEquals(cty.NumberIntVal(20)))
	assert.True(t, valueMap["salary"].RawEquals(cty.NumberFloatVal(800.34)))
	assert.True(t, valueMap["male"].RawEquals(cty.BoolVal(true)))
	children := valueMap["children"].AsValueSlice()
	assert.True(t, children[0].RawEquals(cty.MapVal(map[string]cty.Value{
		"firstname": cty.StringVal("Tim"),
		"lastname":  cty.StringVal("Doe"),
	})))
	assert.True(t, children[1].RawEquals(cty.MapVal(map[string]cty.Value{
		"firstname": cty.StringVal("Loe"),
		"lastname":  cty.StringVal("Doe"),
	})))
}

func TestToValueList(t *testing.T) {
	result := ToValueList([]interface{}{
		map[string]interface{}{
			"firstname": "Tim",
			"lastname":  "Doe",
		},
		map[string]interface{}{
			"firstname": "Loe",
			"lastname":  "Doe",
		},
	})
	assert.NotNil(t, result)
	children := result.AsValueSlice()
	assert.True(t, children[0].Equals(cty.ObjectVal(map[string]cty.Value{
		"firstname": cty.StringVal("Tim"),
		"lastname":  cty.StringVal("Doe"),
	})).True())
	assert.True(t, children[1].Equals(cty.ObjectVal(map[string]cty.Value{
		"firstname": cty.StringVal("Loe"),
		"lastname":  cty.StringVal("Doe"),
	})).True())
}

func TestToValue(t *testing.T) {
	result := ToValue(12)
	assert.True(t, result.Equals(cty.NumberIntVal(12)).True())
	result = ToValue(true)
	assert.True(t, result.Equals(cty.BoolVal(true)).True())
	result = ToValue(false)
	assert.True(t, result.Equals(cty.BoolVal(false)).True())
	result = ToValue("Jane")
	assert.True(t, result.Equals(cty.StringVal("Jane")).True())
	result = ToValue(12.123123)
	assert.True(t, result.Equals(cty.NumberFloatVal(12.123123)).True())
	result = ToValue([]string{"a", "b"})
	assert.True(t, result.Equals(cty.TupleVal([]cty.Value{
		cty.StringVal("a"), cty.StringVal("b"),
	})).True())
	result = ToValue(map[string]string{"a": "a"})
	assert.True(t, result.Equals(cty.MapVal(map[string]cty.Value{
		"a": cty.StringVal("a"),
	})).True())
}
