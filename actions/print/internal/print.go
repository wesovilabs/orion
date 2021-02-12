// Package internal contains not exported types and functions
package internal

import (
	"encoding/json"
	"fmt"
	"time"

	ct "github.com/daviddengcn/go-colortext"
	log "github.com/sirupsen/logrus"
	"github.com/wesovilabs-tools/orion/internal/errors"
)

// DefTimestampFormat default timestamp format used by the plugin.
const DefTimestampFormat = "2006-01-02 15:04:05"

// Print contains the required information to print a message into the stdout.
type Print struct {
	Prefix          string
	Message         string
	ShowTimestamp   bool
	TimestampFormat string
	Format          string
}

func (action *Print) normalize() {
	if action.ShowTimestamp && action.TimestampFormat == "" {
		action.TimestampFormat = DefTimestampFormat
	}
}

// Execute function in charge of the plugin execution.
func (action *Print) Execute() errors.Error {
	action.normalize()
	ct.ResetColor()
	switch action.Format {
	case "json":
		return action.json()
	default:
		action.plain()
	}

	return nil
}

func (action *Print) plain() {
	prefix := ""
	if action.Prefix != "" {
		prefix += fmt.Sprintf("%s ", action.Prefix)
	}

	if action.ShowTimestamp {
		now := time.Now()
		fmt.Printf("%s%s %s\n", prefix, now.Format(action.TimestampFormat), action.Message)

		return
	}

	fmt.Printf("%s%s\n", prefix, action.Message)
}

func (action *Print) json() errors.Error {
	value := &jsonFormat{
		Prefix:    action.Prefix,
		Msg:       action.Message,
		Timestamp: "",
	}
	if action.ShowTimestamp {
		now := time.Now()
		value.Timestamp = now.Format(action.TimestampFormat)
	}
	bytes, err := json.Marshal(value)
	if err != nil {
		log.Warn(err.Error())

		return errors.Unexpected(err.Error())
	}
	fmt.Println(string(bytes))
	return nil
}

type jsonFormat struct {
	Prefix    string `json:"prefix,omitempty"`
	Msg       string `json:"msg"`
	Timestamp string `json:"timestamp,omitempty"`
}
