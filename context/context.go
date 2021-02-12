package context

import (
	"github.com/hashicorp/hcl/v2"
	"github.com/wesovilabs/orion/functions"
	"github.com/zclconf/go-cty/cty"
)

// FeatureContext context used during the feature execution.
type FeatureContext interface {
	StartScenario()
	FailScenario()
	StopScenario()
	EvalContext() *hcl.EvalContext
	Variables() Variables
}

// NewFeatureContext returns a initialization of interface FeatureContext.
func NewFeatureContext(variables map[string]cty.Value) FeatureContext {
	vars := make(map[string]cty.Value)
	for name, value := range variables {
		vars[name] = value
	}

	return &featureContext{
		ctx: &hcl.EvalContext{
			Functions: functions.Functions,
			Variables: vars,
		},
		variables: createVariables(),
	}
}

type featureContext struct {
	ctx       *hcl.EvalContext
	variables Variables
	metrics   metric
}

// StartScenario starts the scenario.
func (c *featureContext) StartScenario() {
	c.metrics = newScenarioMetrics()
}

// CompleteScenario completes the scenario.
func (c *featureContext) CompleteScenario() {
	c.metrics.stopScenario()
}

// Variables return set of variables.
func (c *featureContext) Variables() Variables {
	return c.variables
}

// FailScenario scenario failed.
func (c *featureContext) FailScenario() {
	c.metrics.stopScenario()
}

// StopScenario scenario stopped.
func (c *featureContext) StopScenario() {
	c.metrics.stopScenario()
}

// EvalContext return the hcl eval context.
func (c *featureContext) EvalContext() *hcl.EvalContext {
	return c.ctx
}
