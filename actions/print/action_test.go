package pprint

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"path"
	"strings"
	"testing"
	"time"

	"github.com/wesovilabs/orion/actions/shared"
	"github.com/wesovilabs/orion/context"

	"github.com/wesovilabs/orion/actions/print/internal"

	"github.com/hashicorp/hcl/v2"
	"github.com/stretchr/testify/assert"
	"github.com/zclconf/go-cty/cty"
)

const (
	featureScenario1 = "scenario1.hcl"
	featuresDir      = "testdata"
	featureShowcases = "scenario2.hcl"
)

var decoder = new(Decoder)

var prints = []internal.Print{
	{
		Prefix:          "DEBUG",
		Message:         "Hello Mr Robot",
		ShowTimestamp:   defTimestamp,
		TimestampFormat: internal.DefTimestampFormat,
		Format:          defFormat,
	},
	{
		Prefix:          "DEBUG",
		Message:         "Hello Mr Robot",
		ShowTimestamp:   defTimestamp,
		TimestampFormat: internal.DefTimestampFormat,
		Format:          defFormat,
	},
	{
		Prefix:          "DEBUG",
		Message:         "Hello Mr Robot",
		ShowTimestamp:   true,
		TimestampFormat: internal.DefTimestampFormat,
		Format:          defFormat,
	},
	{
		Prefix:          "DEBUG",
		Message:         "Hello Mr Robot",
		ShowTimestamp:   true,
		TimestampFormat: internal.DefTimestampFormat,
		Format:          defFormat,
	},
	{
		Prefix:          "DEBUG",
		Message:         "Hello Mr Robot",
		ShowTimestamp:   true,
		TimestampFormat: internal.DefTimestampFormat,
		Format:          defFormat,
	},
}

func TestPrint(t *testing.T) {
	content, err := shared.GetBodyContent(path.Join(featuresDir, featureScenario1), BlockPrint, []string{})
	assert.Nil(t, err)
	assert.NotNil(t, content)
	ctx := &hcl.EvalContext{
		Variables: map[string]cty.Value{
			"lastname":      cty.StringVal("Robot"),
			"showTimestamp": cty.BoolVal(true),
		},
		Functions: nil,
	}
	for index := range content.Blocks {
		expected := prints[index]
		action, err := decoder.DecodeBlock(content.Blocks[index])
		assert.Nil(t, err)
		assert.NotNil(t, action)
		print, ok := action.(*Print)
		assert.True(t, ok)
		assert.NotNil(t, print)
		prefix, err := print.Prefix(ctx)
		assert.Nil(t, err)
		assert.Equal(t, expected.Prefix, prefix)
		msg, err := print.Msg(ctx)
		assert.Nil(t, err)
		assert.Equal(t, expected.Message, msg)
		timestampFormat, err := print.TimestampFormat(ctx)
		assert.Nil(t, err)
		assert.Equal(t, expected.TimestampFormat, timestampFormat)
		showTimestamp, err := print.MustShowTimestamp(ctx)
		assert.Nil(t, err)
		assert.Equal(t, expected.ShowTimestamp, showTimestamp)
		format, err := print.Format(ctx)
		assert.Nil(t, err)
		assert.Equal(t, expected.Format, format)
	}
}

func TestPrint_Execute(t *testing.T) {
	messages := []string{
		"Hello Mr Robot",
		"DEBUG Hello Mr Robot",
		"DEBUG Hello Mr Robot",
		"Robot",
		"true",
		"{\"msg\":\"Robot_Robot\"}",
		"Robot_Robot",
		fmt.Sprintf("{\"msg\":\"Hello Mr Robot\",\"timestamp\":\"%s\"}", time.Now().Format("2006-01-02 3:04")),
	}
	content, err := shared.GetBodyContent(path.Join(featuresDir, featureShowcases), BlockPrint, []string{})
	assert.Nil(t, err)
	assert.NotNil(t, content)
	variables := map[string]cty.Value{
		"lastname":      cty.StringVal("Robot"),
		"showTimestamp": cty.BoolVal(true),
	}
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
	for index := range content.Blocks {
		action, err := decoder.DecodeBlock(content.Blocks[index])
		assert.Nil(t, err)
		assert.NotNil(t, action)
		print, ok := action.(*Print)
		assert.True(t, ok)
		assert.NotNil(t, print)
		err = print.Execute(context.New(variables,nil))
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
