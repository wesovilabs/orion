package actions

import (
	"github.com/hashicorp/hcl/v2"
)

const (
	// ArgDesc contains the name of the argument used in the plugins block.
	ArgDesc = "description"
	// ArgWhen contains the name of the argument used in the plugins block.
	ArgWhen = "when"
	// ArgWhile contains the name of the argument used in the plugins block.
	ArgWhile = "while"
	// ArgCount contains the name of the argument used in the plugins block.
	ArgCount = "count"
	// ArgItems contains the name of the slice to be iterated over
	ArgItems = "items"
)

// BaseArguments contain the list of common attributes.
var BaseArguments = []hcl.AttributeSchema{
	{Name: ArgDesc, Required: false},
	{Name: ArgWhen, Required: false},
	{Name: ArgWhile, Required: false},
	{Name: ArgCount, Required: false},
	{Name: ArgItems, Required: false},
}

var baseArgumentsNames = map[string]struct{}{
	ArgDesc:  {},
	ArgWhen:  {},
	ArgWhile: {},
	ArgCount: {},
	ArgItems: {},
}

// IsPluginBaseArgument returns true if the name is common for all the plugins.
func IsPluginBaseArgument(name string) bool {
	_, ok := baseArgumentsNames[name]

	return ok
}
