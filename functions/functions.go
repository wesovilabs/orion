package functions

import "github.com/zclconf/go-cty/cty/function"

// Functions contain the list of provided functions.
var Functions = map[string]function.Function{

	opEqIgnoreCase: eqIgnoreCase,
	opLen:          lenOp,
	opEval:         eval,
	opJSON:         toJSON,

	// string converter
	opToUppercase: toUppercase,
	opToLowercase: toLowercase,
	opReplaceAll:  replaceAll,
	opReplaceOne:  replaceOne,
	opTrimPrefix:  trimPrefix,
	opTrimSuffix:  trimSuffix,
	// string search
	opContains:    containsStr,
	opHasPrefix:   hasPrefix,
	opHasSuffix:   hasSuffix,
	opCount:       count,
	opIndexOf:     indexOf,
	opLastIndexOf: lastIndexOf,
	// string others
	opSplit: split,

	// Numbers - converter
	opSqrt:  sqrt,
	opCos:   cos,
	opSin:   sin,
	opRound: round,
	opPow:   pow,
	opMod:   mod,
	opMax:   max,
	opMin:   min,

	// collection operations
	opFirst: first,
	opLast:  last,
}
