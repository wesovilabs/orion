package functions

import (
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

func invalidArgs(name string, args int) error {
	return fmt.Errorf("function `%s` must retrieve %d arguments", name, args)
}

func invalidArgType(name string, index int, t string) error {
	return fmt.Errorf("argument at position %d in function `%s` must be %s", index, name, t)
}
