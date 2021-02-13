package functions

import (
	"errors"
	"fmt"

	"github.com/zclconf/go-cty/cty"
)

func checkArgumentType(operation string, args []cty.Value, types ...cty.Type) error {
	if len(args) != len(types) {
		return invalidArgs(operation, 1)
	}
	for index := range args {
		if args[index].Type() != types[index] {
			return invalidArgType(operation, index, cty.String.GoString())
		}
	}
	return nil
}

// nolint:goerr113
func invalidArgs(name string, args int) error {
	errMsg := fmt.Sprintf("function `%s` must retrieve %d arguments", name, args)
	return errors.New(errMsg)
}

// nolint:goerr113
func invalidArgType(name string, index int, t string) error {
	errMsg := fmt.Sprintf("argument at position %d in function `%s` must be %s", index, name, t)
	return errors.New(errMsg)
}
