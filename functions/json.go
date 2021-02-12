package functions

import (
	"encoding/json"
	"errors"
	"reflect"

	"github.com/zclconf/go-cty/cty"
	"github.com/zclconf/go-cty/cty/function"

	"github.com/wesovilabs/orion/helper"
)

const opJSON = "json"

var errUnmarshalJSON = errors.New("error un-marshaling response")

var toJSON = function.New(&function.Spec{
	VarParam: nil,
	Params: []function.Parameter{
		{
			Type: cty.String,
		},
	},
	Type: func(args []cty.Value) (cty.Type, error) {
		return cty.DynamicPseudoType, nil
	},
	Impl: func(args []cty.Value, retType cty.Type) (cty.Value, error) {
		content := []byte(args[0].AsString())
		var output interface{}
		if err := json.Unmarshal(content, &output); err != nil {
			return cty.NilVal, errUnmarshalJSON
		}
		if reflect.TypeOf(output).Kind() == reflect.Map {
			return helper.ToValueMap(output.(map[string]interface{})), nil
		}
		return helper.ToValueList(output.([]interface{})), nil
	},
})
