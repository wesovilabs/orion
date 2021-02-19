// Package value contains types and methods to deal with cty value
package helper

import (
	"fmt"
	"log"

	"github.com/wesovilabs/orion/internal/errors"
	"github.com/zclconf/go-cty/cty"
)

// IsSlice returns true if value is a slice.
func IsSlice(value cty.Value) bool {
	return value.Type().IsListType() || value.Type().IsTupleType() || value.Type().IsCollectionType()
}

// ToStrictString convert the value into a string.
func ToStrictString(value cty.Value) (string, errors.Error) {
	if value.Type() == cty.String {
		return value.AsString(), nil
	}

	return "", errors.InvalidArguments("expected a string field but it's not")
}

// IsMap return true if value is a map.
func IsMap(value cty.Value) bool {
	return value.Type().IsMapType() || value.Type().IsObjectType()
}

// ToValueMap convert the map into a value.
func ToValueMap(input map[string]interface{}) cty.Value {
	output := make(map[string]cty.Value)
	for name, value := range input {
		output[name] = ToValue(value)
	}

	return cty.ObjectVal(output)
}

// ToValueList convert the array into a value.
func ToValueList(input []interface{}) cty.Value {
	output := make([]cty.Value, len(input))
	for index := range input {
		value := input[index]
		output[index] = ToValue(value)
	}

	return cty.TupleVal(output)
}

// ToValue convert the interface into a value.
func ToValue(value interface{}) cty.Value {
	switch v := value.(type) {
	case string:
		return cty.StringVal(v)
	case int:
		return cty.NumberIntVal(int64(v))
	case int8, int16, int32, int64:
		return cty.NumberIntVal(v.(int64))
	case float32, float64:
		return cty.NumberFloatVal(v.(float64))
	case bool:
		return cty.BoolVal(v)
	case map[string]interface{}:
		return ToValueMap(v)
	case []interface{}:
		return ToValueList(v)
	case nil:
		return cty.NilVal
	default:
		log.Fatal(fmt.Sprintf("unsupported type %s\n", v))
	}

	return cty.NilVal
}
