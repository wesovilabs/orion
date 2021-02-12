package context

import (
	"time"

	log "github.com/sirupsen/logrus"
)

type metric interface {
	stopScenario()
}

type scenarioMetrics struct {
	startTime time.Time
	endTime   time.Time
}

func (m *scenarioMetrics) duration() time.Duration {
	return m.endTime.Sub(m.startTime)
}

func newScenarioMetrics() *scenarioMetrics {
	return &scenarioMetrics{
		startTime: time.Now(),
	}
}

func (m *scenarioMetrics) stopScenario() {
	m.endTime = time.Now()
	log.Infof("The scenario took %s.", m.duration().String())
}
