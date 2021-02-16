package executor

import (
	"fmt"
	"path/filepath"

	"github.com/wesovilabs/orion/dsl"

	"github.com/hashicorp/hcl/v2"
	log "github.com/sirupsen/logrus"
	"github.com/wesovilabs/orion/context"
	"github.com/wesovilabs/orion/helper"
	"github.com/wesovilabs/orion/internal/errors"
	"github.com/zclconf/go-cty/cty"
)

type Executor interface {
	SetUp(path string) errors.Error
	Run(variablesPath map[string]cty.Value) errors.Error
}

func New() Executor {
	return &executor{
		dec: new(decoder),
	}
}

type executor struct {
	dec     Decoder
	feature *dsl.Feature
}

func (e *executor) SetUp(path string) errors.Error {
	absPath, _ := filepath.Abs(path)
	log.Infof("[feat: %s]", absPath)
	feat, err := e.Parse(path)
	if err != nil {
		return errors.Unexpected(err.Error()).ThroughBy(err)
	}
	log.Debug("Input file is parsed successfully")
	log.Tracef("The feature contains %d scenarios", len(feat.Scenarios()))
	e.feature = feat
	return nil
}

func (e *executor) Run(envVariables map[string]cty.Value) errors.Error {
	includes := e.feature.Includes()
	for index := range includes {
		path := includes[index].Path()
		basePath := filepath.Dir(e.feature.Path())
		includeFeature, err := e.Parse(filepath.Join(basePath, path))
		if err != nil {
			return err
		}
		e.feature.Join(includeFeature)
	}
	if err := e.feature.Vars().To(envVariables); err != nil {
		return err
	}

	for index := range e.feature.Scenarios() {
		fmt.Println()
		ctx := context.New(envVariables)
		if err := e.feature.LoadVariables(ctx); err != nil {
			return err
		}
		scenario := e.feature.Scenarios()[index]
		log.Info(scenario)
		if scenario.IsIgnored(ctx.EvalContext()) {
			log.Warning("scenario is skipped!")
			continue
		}
		examples, err := scenario.Examples(ctx)
		if err != nil {
			return err
		}
		beforeHooks := e.feature.BeforeHooksByTag(scenario.Tags())
		afterHooks := e.feature.AfterHooksByTag(scenario.Tags())
		if examples != nil {
			log.Debug("It starts the execution wth examples")
			for index := range examples {
				example := examples[index]
				for k, v := range example {
					ctx.EvalContext().Variables[k] = v
				}
				if err := runScenario(ctx, scenario, beforeHooks, afterHooks); err != nil {
					return err
				}
				fmt.Println()
			}
			continue
		}
		if err := runScenario(ctx, scenario, beforeHooks, afterHooks); err != nil {
			return err
		}
	}
	return nil
}

func runScenario(ctx context.OrionContext, scenario *dsl.Scenario, beforeHooks, afterHooks []*dsl.Hook) errors.Error {
	log.Debug("It starts the execution")
	printVariables(ctx.EvalContext())
	if beforeHooks != nil {
		if err := runHooks(ctx, beforeHooks); err != nil {
			return err
		}
		printVariables(ctx.EvalContext())
	}
	if err := scenario.Execute(ctx); err != nil {
		if !scenario.ContinueOnError(ctx.EvalContext()) {
			return err
		}
		log.Warning(err.Error())
	}
	printVariables(ctx.EvalContext())
	if afterHooks != nil {
		if err := runHooks(ctx, afterHooks); err != nil {
			return err
		}
		printVariables(ctx.EvalContext())
	}
	return nil
}

func runHooks(ctx context.OrionContext, hooks []*dsl.Hook) errors.Error {
	for index := range hooks {
		hook := hooks[index]
		log.Info(hook)
		if err := hook.Execute(ctx); err != nil {
			return err
		}
	}
	return nil
}

func printVariables(ctx *hcl.EvalContext) {
	if log.IsLevelEnabled(log.TraceLevel) {
		msg := fmt.Sprintf("There are %d available variables", len(ctx.Variables))
		for name, value := range ctx.Variables {
			val, _ := helper.ToString(value)
			msg += fmt.Sprintf("\n - %s: %s", name, val)
		}
		log.Trace(msg)
	}
}

func runFunction(ctx context.OrionContext, f *dsl.Function, out string) errors.Error {
	if input := f.Input(); input != nil {
		if err := input.Execute(ctx); err != nil {
			return err
		}
	}
	actions := f.Body().Actions()
	for index := range actions {
		action := actions[index]
		if action.ShouldExecute(ctx.EvalContext()) {
			if err := action.Execute(ctx); err != nil {
				return err
			}
			continue
		}
		log.Debugf("action %s is skipped!", action)
	}
	if f.Return() != nil {
		result, d := f.Return().Value().Value(ctx.EvalContext())
		if err := errors.EvalDiagnostics(d); err != nil {
			return err
		}
		ctx.EvalContext().Variables[out] = result
	}
	return nil
}
