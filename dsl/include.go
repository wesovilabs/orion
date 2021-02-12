package dsl

import (
	"strings"

	"github.com/hashicorp/hcl/v2"

	"github.com/wesovilabs/orion/helper"
	"github.com/wesovilabs/orion/internal/errors"
)

const (
	filePath = "file://"
	fileType = "file"
)

type Includes []*Include

type Include struct {
	kind string
	path string
}

func (i *Include) IsFile() bool {
	return i.kind == fileType
}

func (i *Include) Path() string {
	return i.path
}

func includesFromValue(attr *hcl.Attribute) (Includes, errors.Error) {
	includesValue, d := attr.Expr.Value(nil)
	if err := errors.EvalDiagnostics(d); err != nil {
		return nil, err
	}
	if !helper.IsSlice(includesValue) {
		return nil, errors.IncorrectUsage("attribute '$s' must be a slice", argIncludes)
	}

	slice := includesValue.AsValueSlice()
	includes := make(Includes, len(slice))
	for index := range slice {
		itemValue := slice[index]
		path, err := helper.ToStrictString(itemValue)
		if err != nil {
			return nil, err
		}
		var include *Include
		switch {
		case strings.HasPrefix(path, filePath):
			include = &Include{
				kind: fileType,
				path: strings.TrimLeft(path, filePath),
			}
		default:
			include = &Include{
				kind: fileType,
				path: path,
			}
		}
		includes[index] = include
	}
	return includes, nil
}
