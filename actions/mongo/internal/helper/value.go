package helper

import (
	"fmt"

	"github.com/zclconf/go-cty/cty"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// ToValue convert the interface into a value.
func ToValue(v interface{}) cty.Value {
	switch v := v.(type) {
	case string:
		return cty.StringVal(v)
	case int, int8, int16, int32, int64:
		return cty.NumberIntVal(v.(int64))
	case float32, float64:
		return cty.NumberFloatVal(v.(float64))
	case bool:
		return cty.BoolVal(v)
	case map[string]interface{}:
		return toValueMap(v)
	case []interface{}:
		return toValueList(v)
	case []map[string]interface{}:
		items := make([]cty.Value, len(v))
		for i := range v {
			items[i] = toValueMap(v[i])
		}
		return cty.ListVal(items)
	case nil:
		return cty.NilVal
	case primitive.ObjectID:
		return cty.StringVal(v.Hex())
	default:
		return cty.StringVal(fmt.Sprintf("%v", v))
	}

	return cty.NilVal
}

func toValueMap(input map[string]interface{}) cty.Value {
	output := make(map[string]cty.Value)
	for name, value := range input {
		output[name] = ToValue(value)
	}

	return cty.ObjectVal(output)
}

func toValueList(input []interface{}) cty.Value {
	output := make([]cty.Value, len(input))
	for index := range input {
		value := input[index]
		output[index] = ToValue(value)
	}
	return cty.TupleVal(output)
}
