package actions

import (
	"github.com/wesovilabs-tools/orion/context"
	"github.com/wesovilabs-tools/orion/helper"
	"github.com/wesovilabs-tools/orion/internal/errors"
)

// ExecuteFn defined type to be implemented by the plugins.
type ExecuteFn func(ctx context.FeatureContext) errors.Error

// Execute Given a plugin and the context is executed.
func Execute(ctx context.FeatureContext, action *Base, fn ExecuteFn) errors.Error {
	// ctx.StartAction()
	evalCtx := ctx.EvalContext()
	when, err := helper.GetExpressionValueAsBool(evalCtx, action.when, true)
	if err != nil {
		return err
	}
	if !when {
		return nil
	}
	if action.count != nil && !action.count.Range().Empty() {
		return doCount(ctx, action, fn)
	}
	if action.while != nil && !action.while.Range().Empty() {
		return doWhile(ctx, action, fn)
	}
	return fn(ctx)
}

func doWhile(ctx context.FeatureContext, plugin *Base, fn ExecuteFn) errors.Error {
	evalCtx := ctx.EvalContext()
	index := 0
	for {
		ctx.Variables().SetIndex(index)
		ctx.Variables().SetToContext(ctx.EvalContext())
		if while, err := helper.GetExpressionValueAsBool(evalCtx, plugin.while, true); err != nil || !while {
			return err
		}
		if err := fn(ctx); err != nil {
			return err
		}
		index++
	}
}

func doCount(ctx context.FeatureContext, plugin *Base, fn ExecuteFn) errors.Error {
	evalCtx := ctx.EvalContext()
	count, err := helper.GetExpressionValueAsInt(evalCtx, plugin.count, 1)
	if err != nil {
		return err
	}
	for index := 0; index < count; index++ {
		ctx.Variables().SetIndex(index)
		ctx.Variables().SetToContext(ctx.EvalContext())
		if while, err := helper.GetExpressionValueAsBool(evalCtx, plugin.while, true); err != nil || !while {
			return err
		}
		if err := fn(ctx); err != nil {
			return err
		}
	}
	return nil
}
