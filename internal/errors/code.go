// Package errors content relative to the package
package errors

// code type used to define the error codes in koazee.
type code string

const (
	unexpectedErrCode       = "unexpectedErrCode"
	invalidArgumentsErrCode = "invalidArgumentsErrCode"
	commandNotFoundErrCode  = "commandNotFoundErrCode"
	incorrectUsageErrCode   = "incorrectUsageErrCode"
	limitsExceedErrCode     = "limitsExceedErrCode"
)

// Unexpected return an unexpectedErrCode error.
func Unexpected(msg string, args ...interface{}) Error {
	return createError(unexpectedErrCode, 1).
		withMessage(msg, args)
}

// InvalidArguments return an invalid arguments error.
func InvalidArguments(msg string, args ...interface{}) Error {
	return createError(invalidArgumentsErrCode, 128).
		withMessage(msg, args)
}

// CommandNotFound return a command not found error.
func CommandNotFound(msg string, args ...interface{}) Error {
	return createError(commandNotFoundErrCode, 127).
		withMessage(msg, args)
}

// IncorrectUsage return an incorrect usage error.
func IncorrectUsage(msg string, args ...interface{}) Error {
	return createError(incorrectUsageErrCode, 2).
		withMessage(msg, args)
}

// LimitsExceed return an limit exceed error.
func LimitsExceed(msg string, args ...interface{}) Error {
	return createError(limitsExceedErrCode, 2).
		withMessage(msg, args)
}

// String print string of error code.
func (c code) String() string {
	return string(c)
}
