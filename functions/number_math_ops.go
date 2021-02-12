package functions

import (
	"math"

	"github.com/zclconf/go-cty/cty"
	"github.com/zclconf/go-cty/cty/function"
)

const (
	opSqrt  = "sqrt"
	opCos   = "cos"
	opSin   = "sin"
	opRound = "round"
	opPow   = "pow"
	opMod   = "mod"
	opMax   = "max"
	opMin   = "min"
)

var (
	sqrt  = function.New(numberConverter(opSqrt, math.Sqrt))
	cos   = function.New(numberConverter(opCos, math.Cos))
	sin   = function.New(numberConverter(opSin, math.Sin))
	round = function.New(numberConverter(opRound, math.Round))
	pow   = function.New(numberOperationWithTwoFloat64(opPow, math.Pow))
	mod   = function.New(numberOperationWithTwoFloat64(opMod, math.Mod))
	max   = function.New(numberOperationWithTwoFloat64(opMod, math.Max))
	min   = function.New(numberOperationWithTwoFloat64(opMod, math.Min))
)

func numberConverter(operation string, fn func(float64) float64) *function.Spec {
	return &function.Spec{
		VarParam: nil,
		Params: []function.Parameter{
			{Type: cty.Number},
		},
		Type: func(args []cty.Value) (cty.Type, error) {
			return cty.Number, checkArgumentType(operation, args, cty.Number)
		},
		Impl: func(args []cty.Value, retType cty.Type) (cty.Value, error) {
			val := args[0].AsBigFloat()
			var valFloat float64
			if val.IsInt() {
				valInt, _ := val.Int64()
				valFloat = float64(valInt)
			} else {
				valFloat, _ = val.Float64()
			}
			res := fn(valFloat)
			return cty.NumberFloatVal(res), nil
		},
	}
}

func numberOperationWithTwoFloat64(operation string, fn func(float64, float64) float64) *function.Spec {
	return &function.Spec{
		VarParam: nil,
		Params: []function.Parameter{
			{Type: cty.Number},
			{Type: cty.Number},
		},
		Type: func(args []cty.Value) (cty.Type, error) {
			return cty.Number, checkArgumentType(operation, args, cty.Number, cty.Number)
		},
		Impl: func(args []cty.Value, retType cty.Type) (cty.Value, error) {
			val1 := getValueAsFloat64(args[0])
			val2 := getValueAsFloat64(args[1])
			res := fn(val1, val2)
			return cty.NumberFloatVal(res), nil
		},
	}
}

func getValueAsFloat64(value cty.Value) float64 {
	val := value.AsBigFloat()
	if val.IsInt() {
		valInt, _ := val.Int64()
		return float64(valInt)
	}
	valFloat, _ := val.Float64()
	return valFloat
}
