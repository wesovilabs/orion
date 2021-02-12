package errors

import (
	"encoding/json"
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type inputErrorTest struct {
	parent       error
	msgFormat    string
	args         []interface{}
	meta         map[string]interface{}
	expectedJSON string
}

var cases = []*inputErrorTest{
	{
		msgFormat:    "message with args %s %s",
		args:         []interface{}{"hello", "bye"},
		expectedJSON: `{"code":"%s","msg":"message with args hello bye"}`,
	},
	{
		msgFormat:    "message with args %s %v",
		args:         []interface{}{"hello", 20},
		expectedJSON: `{"code":"%s","msg":"message with args hello 20"}`,
	},
	{
		msgFormat:    "message with args %s %v and parent",
		args:         []interface{}{"hello", 20},
		expectedJSON: `{"code":"%s","msg":"message with args hello 20 and parent","parent":"parent error"}`,
		parent:       errors.New("parent error"),
	},
	{
		msgFormat: "message with args %s %v and meta",
		args:      []interface{}{"hello", 20},
		meta: map[string]interface{}{
			"a": 20,
			"b": true,
			"c": map[string]interface{}{
				"c1": 120.23,
				"c2": "bye",
			},
		},
		expectedJSON: `{"code":"%s","msg":"message with args hello 20 and meta","meta":{"a":20,"b":true,"c":{"c1":120.23,"c2":"bye"}}}`,
	},
}

func TestCommandNotFound(t *testing.T) {
	for _, c := range cases {
		err := CommandNotFound(c.msgFormat, c.args...)
		assertTest(t, commandNotFoundErrCode, 127, c, err)
	}
}

func TestIncorrectUsage(t *testing.T) {
	for _, c := range cases {
		err := IncorrectUsage(c.msgFormat, c.args...)
		assertTest(t, incorrectUsageErrCode, 2, c, err)
	}
}

func TestUnexpected(t *testing.T) {
	for _, c := range cases {
		err := Unexpected(c.msgFormat, c.args...)
		assertTest(t, unexpectedErrCode, 1, c, err)
	}
}

func TestInvalidArguments(t *testing.T) {
	for _, c := range cases {
		err := InvalidArguments(c.msgFormat, c.args...)
		assertTest(t, invalidArgumentsErrCode, 128, c, err)
	}
}

func assertTest(t *testing.T, errCode code, exitStatus int, input *inputErrorTest, internalErr Error) {
	assert.NotNil(t, internalErr)
	internalErr.ThroughBy(input.parent)
	for k, v := range input.meta {
		internalErr = internalErr.AddMeta(k, v)
	}
	assert.EqualValues(t, internalErr.getCode(), errCode)
	assert.Equal(t, internalErr.ExitStatus(), exitStatus)
	b, err := json.Marshal(internalErr)
	assert.Nil(t, err)
	json := fmt.Sprintf(input.expectedJSON, errCode)
	assert.EqualValues(t, string(b), json)
}
