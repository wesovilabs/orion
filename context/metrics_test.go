package context

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestScenarioMetric(t *testing.T) {
	m := newScenarioMetrics()
	assert.NotNil(t, m.startTime)
	assert.Nil(t, m.endTime)
	m.stopScenario()
	assert.NotNil(t, m.startTime)
	assert.NotNil(t, m.endTime)
	assert.Equal(t, m.endTime.Sub(m.startTime), m.duration())
}
