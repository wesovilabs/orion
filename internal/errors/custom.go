package errors

// ThrowMissingRequiredLabel throw a missing required label error.
func ThrowMissingRequiredLabel(block string) Error {
	return IncorrectUsage("missing required label in block %s", block)
}

// ThrowUnsupportedArgument throw an unsupported argument error.
func ThrowUnsupportedArgument(block, argument string) Error {
	return IncorrectUsage("unsupported argument %s in block %s", argument, block)
}

// ThrowUnsupportedBlock throw an unsupported block error.
func ThrowUnsupportedBlock(parent, child string) Error {
	return IncorrectUsage("unsupported block %s in block %s", child, parent)
}

// ThrowBlocksAreNotPermitted throw an unsupported block error.
func ThrowBlocksAreNotPermitted(action string) Error {
	return IncorrectUsage("blocks are not permitted in action %s ", action)
}

func ThrowUnsupportedBlocks(block string) Error {
	return IncorrectUsage("blocks are not supported by  %s", block)
}

func ThrowUnsupportedArguments(block string) Error {
	return IncorrectUsage("arguments are not supported by  %s", block)
}

func ThrowsExceeddedNumberOfBlocks(block string, max int) Error {
	return IncorrectUsage("only %d block '%s' is permitted here", max, block)
}
