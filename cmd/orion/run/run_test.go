package run

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func messageWithTimestamp(msg string) string {
	return fmt.Sprintf("[0-9]{2}:[0-9]{2}:[0-9]{2}.[0-9]{6} %s", msg)
}

var testRunData = []struct {
	input          string
	vars           string
	tags           []string
	expectMessages []string
}{
	{
		input: "feature001.hcl",
		expectMessages: []string{
			messageWithTimestamp(`\[feat: (.*)/feature001.hcl\]`),
			messageWithTimestamp(`\[tags: \[\]\]`),
			``,
			messageWithTimestamp(`\[scenario\] basic usage`),
			messageWithTimestamp(`Given a couple of numbers`),
			messageWithTimestamp(`When input values are sum`),
			messageWithTimestamp(`Then the result of variable c is correct`),
			messageWithTimestamp(`The scenario took .*.`),
			``,
		},
	},
	{
		input: "feature001.hcl",
		tags:  []string{"basic", "complex"},
		expectMessages: []string{
			messageWithTimestamp(`\[feat: (.*)/feature001.hcl\]`),
			messageWithTimestamp(`\[tags: \[basic complex\]\]`),
			``,
			messageWithTimestamp(`\[scenario\] basic usage`),
			messageWithTimestamp(`Given a couple of numbers`),
			messageWithTimestamp(`When input values are sum`),
			messageWithTimestamp(`Then the result of variable c is correct`),
			messageWithTimestamp(`The scenario took .*.`),
			``,
		},
	},
	{
		input: "feature001.hcl",
		tags:  []string{"easy", "complex"},
		expectMessages: []string{
			messageWithTimestamp(`\[feat: (.*)/feature001.hcl\]`),
			messageWithTimestamp(`\[tags: \[easy complex\]\]`),
			``,
			messageWithTimestamp(`\[scenario\] basic usage`),
			messageWithTimestamp(`scenario does not contain any of the specified tags`),
			``,
		},
	},
	{
		input: "feature002.hcl",
		expectMessages: []string{
			messageWithTimestamp(`\[feat: (.*)/feature002.hcl\]`),
			messageWithTimestamp(`\[tags: \[\]\]`),
			``,
			messageWithTimestamp(`\[scenario\] basic usage`),
			messageWithTimestamp(`Given a couple of numbers`),
			messageWithTimestamp(`When do addition with the numbers`),
			messageWithTimestamp(`Then the result of variable c is correct`),
			messageWithTimestamp(`When do subtraction with the numbers`),
			messageWithTimestamp(`Then the result of variable c is correct`),
			messageWithTimestamp(`The scenario took .*.`),
			``,
		},
	},
	{
		input: "feature003.hcl",
		expectMessages: []string{
			messageWithTimestamp(`\[feat: (.*)/feature003.hcl\]`),
			messageWithTimestamp(`\[tags: \[\]\]`),
			``,
			messageWithTimestamp(`\[scenario\] basic usage`),
			messageWithTimestamp(`When input values are sum`),
			messageWithTimestamp(`Then the result of variable c is correct`),
			messageWithTimestamp(`The scenario took .*.`),
			``,
		},
	},
	{
		input: "feature004.hcl",
		expectMessages: []string{
			messageWithTimestamp(`\[feat: (.*)/feature004.hcl\]`),
			messageWithTimestamp(`\[tags: \[\]\]`),
			``,
			messageWithTimestamp(`\[scenario\] scenario demo`),
			messageWithTimestamp(`When multiply x \* y`),
			messageWithTimestamp(`Then check the output`),
			messageWithTimestamp(`The scenario took .*.`),
			``,
		},
	},
	{
		input: "feature004.hcl",
		vars:  "variables004.hcl",
		expectMessages: []string{
			messageWithTimestamp(`\[feat: (.*)/feature004.hcl\]`),
			messageWithTimestamp(`\[tags: \[\]\]`),
			``,
			messageWithTimestamp(`\[scenario\] scenario demo`),
			messageWithTimestamp(`When multiply x \* y`),
			messageWithTimestamp(`Then check the output`),
			messageWithTimestamp(`The scenario took .*.`),
			``,
		},
	},
}

func TestRun(t *testing.T) {
	for index := range testRunData {
		old := os.Stdout
		r, w, _ := os.Pipe()
		os.Stdout = w

		outC := make(chan string)
		go func() {
			var buf bytes.Buffer
			io.Copy(&buf, r)
			outC <- buf.String()
		}()

		data := testRunData[index]
		cmd := New()
		b := bytes.NewBufferString("")
		cmd.SetOut(b)
		args := []string{"--input", fmt.Sprintf("testdata/%s", data.input)}
		if data.vars != "" {
			args = append(args, "--vars", fmt.Sprintf("testdata/%s", data.vars))
		}
		if len(data.tags) > 0 {
			args = append(args, "--tags", strings.Join(data.tags, ","))
		}
		cmd.SetArgs(args)
		err := cmd.Execute()
		assert.Nil(t, err)
		w.Close()
		os.Stdout = old
		lines := strings.Split(<-outC, "\n")
		assert.Len(t, lines, len(data.expectMessages))
		for index := range data.expectMessages {
			line := lines[index]
			re := regexp.MustCompile(data.expectMessages[index])
			assert.True(t, re.MatchString(line), "line:'%s' does not match regEx:'%s'", line, re)
		}
	}
}
