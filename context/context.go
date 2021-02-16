package context

import (
	"github.com/hashicorp/hcl/v2"
	"github.com/wesovilabs/orion/functions"
	"github.com/wesovilabs/orion/internal/errors"
	"github.com/zclconf/go-cty/cty"
)

type functionsType map[string]func(ctx OrionContext, as string) errors.Error

// OrionContext context used during the feature execution.
type OrionContext interface {
	StartScenario()
	FailScenario()
	StopScenario()
	EvalContext() *hcl.EvalContext
	Variables() Variables
	Functions() functionsType
}

// New returns a initialization of interface OrionContext.
func New(variables map[string]cty.Value, funcs functionsType) OrionContext {
	vars := make(map[string]cty.Value)
	for name, value := range variables {
		vars[name] = value
	}

	return &orionContext{
		ctx: &hcl.EvalContext{
			Functions: functions.Functions,
			Variables: vars,
		},
		variables: createVariables(),
		functions: funcs,
	}
}

type orionContext struct {
	ctx       *hcl.EvalContext
	variables Variables
	metrics   *scenarioMetrics
	functions map[string]func(ctx OrionContext, as string) errors.Error
}

// StartScenario starts the scenario.
func (c *orionContext) StartScenario() {
	c.metrics = newScenarioMetrics()
}

// CompleteScenario completes the scenario.
func (c *orionContext) CompleteScenario() {
	c.metrics.stopScenario()
}

// Variables return set of variables.
func (c *orionContext) Variables() Variables {
	return c.variables
}

// FailScenario scenario failed.
func (c *orionContext) FailScenario() {
	c.metrics.stopScenario()
}

// StopScenario scenario stopped.
func (c *orionContext) StopScenario() {
	c.metrics.stopScenario()
}

// EvalContext return the hcl eval context.
func (c *orionContext) EvalContext() *hcl.EvalContext {
	return c.ctx
}

// Functions returns the functions.
func (c *orionContext) Functions() functionsType {
	return c.functions
}
