package sleep

import (
	"path"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/wesovilabs/orion/actions/shared"
	"github.com/wesovilabs/orion/context"
	"github.com/zclconf/go-cty/cty"
)

const (
	featuresDir     = "testdata"
	featureScenario = "scenario1.hcl"
)

var decoder = new(Decoder)

func TestSleep(t *testing.T) {
	content, err := shared.GetBodyContent(path.Join(featuresDir, featureScenario), BlockSleep, []string{})
	assert.Nil(t, err)
	assert.NotNil(t, content)

	ctx := context.New(map[string]cty.Value{}, nil)

	for i := range content.Blocks {
		action, err := decoder.DecodeBlock(content.Blocks[i])
		assert.Nil(t, err)
		assert.NotNil(t, action)

		s, ok := action.(*Sleep)
		assert.True(t, ok)
		assert.NotNil(t, s)

		duration, err := s.Duration(ctx.EvalContext())
		assert.Nil(t, err)
		assert.Equal(t, 10*time.Millisecond, duration)

		start := time.Now()
		s.Execute(ctx)
		actualDuration := time.Now().Sub(start)
		assert.GreaterOrEqual(t, actualDuration.Milliseconds(), duration.Milliseconds())
		assert.LessOrEqual(t, actualDuration.Milliseconds(), duration.Milliseconds()+10)
	}
}
