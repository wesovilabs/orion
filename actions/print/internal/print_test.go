package internal

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestPrint_Execute(t *testing.T) {
	AssertStdout(t, []Print{
		{
			Message: "Hey Ms Robinson",
		},
		{
			Prefix:  "[DEBUG]",
			Message: "Hey Ms Robinson",
		},
		{
			Prefix:  "[INFO]",
			Message: "Hey Ms Robinson",
		},
		{
			Message:         "Hey Ms Robinson",
			ShowTimestamp:   true,
			TimestampFormat: DefTimestampFormat,
		},
		{
			Message: "Hey Ms Robinson",
			Format:  "json",
		},
		{
			Prefix:        "A",
			Message:       "Hey Ms Robinson",
			ShowTimestamp: true,
			Format:        "json",
		},
	},
		[]string{
			"Hey Ms Robinson",
			"[DEBUG] Hey Ms Robinson",
			"[INFO] Hey Ms Robinson",
			fmt.Sprintf("%s Hey Ms Robinson", time.Now().Format(DefTimestampFormat)),
			"{\"msg\":\"Hey Ms Robinson\"}",
			fmt.Sprintf("{\"prefix\":\"A\",\"msg\":\"Hey Ms Robinson\",\"timestamp\":\"%s\"}", time.Now().Format(DefTimestampFormat)),
		},
	)
}

func AssertStdout(t *testing.T, prints []Print, messages []string) {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	print()
	outC := make(chan string)
	go func() {
		var buf bytes.Buffer
		io.Copy(&buf, r)
		outC <- buf.String()
	}()
	for index := range prints {
		err := prints[index].Execute()
		assert.Nil(t, err)
	}
	w.Close()
	os.Stdout = old

	lines := strings.Split(<-outC, "\n")
	assert.Len(t, lines, len(messages)+1)
	for index := range messages {
		line := lines[index]
		assert.Equal(t, line, messages[index])
	}
}
