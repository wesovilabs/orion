package functions

import (
	"strings"

	"github.com/zclconf/go-cty/cty"
	"github.com/zclconf/go-cty/cty/function"
)

const (
	opToLowercase = "toLowercase"
	opToUppercase = "toUppercase"
	opTrimPrefix  = "trimPrefix"
	opTrimSuffix  = "trimSuffix"
	opReplaceOne  = "replaceOne"
	opReplaceAll  = "replaceAll"
)

var (
	toLowercase = function.New(stringCaseConverter(opToLowercase, strings.ToLower))
	toUppercase = function.New(stringCaseConverter(opToUppercase, strings.ToUpper))
	trimPrefix  = function.New(stringTrim(opTrimPrefix, strings.TrimPrefix))
	trimSuffix  = function.New(stringTrim(opTrimSuffix, strings.TrimSuffix))
	replaceOne  = function.New(stringReplace(opReplaceOne, func(s, old, new string) string {
		return strings.Replace(s, old, new, 1)
	}))
	replaceAll = function.New(stringReplace(opReplaceAll, strings.ReplaceAll))
)

func stringCaseConverter(operation string, fn func(string) string) *function.Spec {
	return &function.Spec{
		VarParam: nil,
		Params: []function.Parameter{
			{Type: cty.String},
		},
		Type: func(args []cty.Value) (cty.Type, error) {
			return cty.String, checkArgumentType(operation, args, cty.String)
		},
		Impl: func(args []cty.Value, retType cty.Type) (cty.Value, error) {
			s := args[0].AsString()
			res := fn(s)
			return cty.StringVal(res), nil
		},
	}
}

func stringTrim(operation string, fn func(string, string) string) *function.Spec {
	return &function.Spec{
		VarParam: nil,
		Params: []function.Parameter{
			{Type: cty.String},
			{Type: cty.String},
		},
		Type: func(args []cty.Value) (cty.Type, error) {
			return cty.String, checkArgumentType(operation, args, cty.String, cty.String)
		},
		Impl: func(args []cty.Value, retType cty.Type) (cty.Value, error) {
			s := args[0].AsString()
			text := args[1].AsString()
			res := fn(s, text)
			return cty.StringVal(res), nil
		},
	}
}

func stringReplace(operation string, fn func(string, string, string) string) *function.Spec {
	return &function.Spec{
		VarParam: nil,
		Params: []function.Parameter{
			{Type: cty.String},
			{Type: cty.String},
			{Type: cty.String},
		},
		Type: func(args []cty.Value) (cty.Type, error) {
			return cty.String, checkArgumentType(operation, args, cty.String, cty.String, cty.String)
		},
		Impl: func(args []cty.Value, retType cty.Type) (cty.Value, error) {
			s := args[0].AsString()
			text := args[1].AsString()
			rep := args[2].AsString()
			res := fn(s, text, rep)
			return cty.StringVal(res), nil
		},
	}
}
