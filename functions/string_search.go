package functions

import (
	"strings"

	"github.com/zclconf/go-cty/cty"
	"github.com/zclconf/go-cty/cty/function"
)

const (
	opHasPrefix   = "hasPrefix"
	opHasSuffix   = "hasSuffix"
	opContains    = "contains"
	opCount       = "count"
	opIndexOf     = "indexOf"
	opLastIndexOf = "lastIndexOf"
)

var (
	hasPrefix   = function.New(textSearcher(opHasPrefix, strings.HasPrefix))
	hasSuffix   = function.New(textSearcher(opHasSuffix, strings.HasSuffix))
	containsStr = function.New(textSearcher(opContains, strings.Contains))
	count       = function.New(stringCount())
	indexOf     = function.New(stringIndexSearch(opIndexOf, strings.Index))
	lastIndexOf = function.New(stringIndexSearch(opLastIndexOf, strings.LastIndex))
)

func textSearcher(operation string, fn func(string, string) bool) *function.Spec {
	return &function.Spec{
		VarParam: nil,
		Params: []function.Parameter{
			{Type: cty.String},
			{Type: cty.String},
		},
		Type: func(args []cty.Value) (cty.Type, error) {
			return cty.Bool, checkArgumentType(operation, args, cty.String, cty.String)
		},
		Impl: func(args []cty.Value, retType cty.Type) (cty.Value, error) {
			s := args[0].AsString()
			text := args[1].AsString()
			res := fn(s, text)
			return cty.BoolVal(res), nil
		},
	}
}

func stringCount() *function.Spec {
	return &function.Spec{
		VarParam: nil,
		Params: []function.Parameter{
			{Type: cty.String},
			{Type: cty.String},
		},
		Type: func(args []cty.Value) (cty.Type, error) {
			return cty.Number, checkArgumentType(opCount, args, cty.String, cty.String)
		},
		Impl: func(args []cty.Value, retType cty.Type) (cty.Value, error) {
			s := args[0].AsString()
			text := args[1].AsString()
			res := strings.Count(s, text)
			return cty.NumberIntVal(int64(res)), nil
		},
	}
}

func stringIndexSearch(op string, fn func(string, string) int) *function.Spec {
	return &function.Spec{
		VarParam: nil,
		Params: []function.Parameter{
			{Type: cty.String},
			{Type: cty.String},
		},
		Type: func(args []cty.Value) (cty.Type, error) {
			return cty.Number, checkArgumentType(op, args, cty.String, cty.String)
		},
		Impl: func(args []cty.Value, retType cty.Type) (cty.Value, error) {
			s := args[0].AsString()
			text := args[1].AsString()
			res := fn(s, text)
			return cty.NumberIntVal(int64(res)), nil
		},
	}
}
