package errors

import (
	"encoding/json"
	"fmt"
)

const (
	// MetaLocation values used to save location information of an error.
	MetaLocation = "loc"
)

// Error encapsulates error info.
type Error interface {
	withMessage(msgFormat string, args []interface{}) Error
	AddMeta(key string, value interface{}) Error
	ThroughBy(err error) Error
	getCode() code
	ExitStatus() int
	Error() string
	Message() string
}

// Error encapsulates error info.
type internalError struct {
	code       code
	parent     error
	exitStatus int
	msg        string
	meta       map[string]interface{}
}

// Error converts the error into a readable message.
func (e *internalError) Error() string {
	out := fmt.Sprintf("[ERR] %s", e.msg)
	if e.meta != nil {
		for k, v := range e.meta {
			out += fmt.Sprintf("\n  - %s: %v", k, v)
		}
	}

	return out
}

func createError(code code, exitStatus int) Error {
	return &internalError{
		code:       code,
		exitStatus: exitStatus,
	}
}

func (e *internalError) getCode() code {
	return e.code
}

func (e *internalError) ExitStatus() int {
	return e.exitStatus
}

func (e *internalError) Message() string {
	return e.msg
}

// withMessage add the message to the error.
func (e *internalError) withMessage(msgFormat string, args []interface{}) Error {
	e.msg = fmt.Sprintf(msgFormat, args...)

	return e
}

// ThroughBy sets the parent error that through it.
func (e *internalError) ThroughBy(err error) Error {
	e.parent = err

	return e
}

// AddMeta permits add extra info to show displayed when printing the error.
func (e *internalError) AddMeta(key string, value interface{}) Error {
	if e.meta == nil {
		e.meta = make(map[string]interface{})
	}
	e.meta[key] = value

	return e
}

func (e *internalError) MarshalJSON() ([]byte, error) {
	var parentErr string
	if e.parent != nil {
		parentErr = e.parent.Error()
	}

	return json.Marshal(struct {
		Code    code                   `json:"code"`
		Message string                 `json:"msg,omitempty"`
		Meta    map[string]interface{} `json:"meta,omitempty"`
		Parent  string                 `json:"parent,omitempty"`
	}{
		Code:    e.code,
		Message: e.msg,
		Meta:    e.meta,
		Parent:  parentErr,
	})
}
