package pprint

import (
	"github.com/hashicorp/hcl/v2"
	log "github.com/sirupsen/logrus"
	"github.com/wesovilabs-tools/orion/actions"
	"github.com/wesovilabs-tools/orion/actions/print/internal"
	"github.com/wesovilabs-tools/orion/context"
	"github.com/wesovilabs-tools/orion/helper"
	"github.com/wesovilabs-tools/orion/internal/errors"
)

const (
	defPrefix    = ""
	defMessage   = ""
	defTimestamp = false

	defFormat = "plain"
)

// Print plugin structure used to marshal the block definition.
type Print struct {
	*actions.Base
	prefix          hcl.Expression
	msg             hcl.Expression
	timestamp       hcl.Expression
	timestampFormat hcl.Expression
	format          hcl.Expression
}

// SetMessage used to populate the attribute msg in the struct.
func (p *Print) SetMessage(expr hcl.Expression) {
	p.msg = expr
}

// Msg returns the message after evaluating the expression.
func (p *Print) Msg(ctx *hcl.EvalContext) (string, errors.Error) {
	return helper.GetExpressionValueAsString(ctx, p.msg, defMessage)
}

// SetPrefix used to populate the attribute prefix in the struct.
func (p *Print) SetPrefix(expr hcl.Expression) {
	p.prefix = expr
}

// Prefix returns the prefix after evaluating the expression.
func (p *Print) Prefix(ctx *hcl.EvalContext) (string, errors.Error) {
	return helper.GetExpressionValueAsString(ctx, p.prefix, defPrefix)
}

// SetTimestamp used to populate the attribute timestamp in the struct.
func (p *Print) SetTimestamp(expr hcl.Expression) {
	p.timestamp = expr
}

// MustShowTimestamp returns if timestamp must be shown after evaluating the expression.
func (p *Print) MustShowTimestamp(ctx *hcl.EvalContext) (bool, errors.Error) {
	return helper.GetExpressionValueAsBool(ctx, p.timestamp, defTimestamp)
}

// SetTimestampFormat used to populate the attribute timestampFormat in the struct.
func (p *Print) SetTimestampFormat(expr hcl.Expression) {
	p.timestampFormat = expr
}

// TimestampFormat returns the format in which the timestamp must be shown after evaluating the expression.
func (p *Print) TimestampFormat(ctx *hcl.EvalContext) (string, errors.Error) {
	timestampFmt, err := helper.GetExpressionValueAsString(ctx, p.timestampFormat, internal.DefTimestampFormat)
	if err != nil {
		return "", err
	}

	return internal.TimeExpression(timestampFmt), nil
}

// SetFormat used to populate the attribute format in the struct.
func (p *Print) SetFormat(expr hcl.Expression) {
	p.format = expr
}

// Format returns the format in which the format in which the message must be printed out.
func (p *Print) Format(ctx *hcl.EvalContext) (string, errors.Error) {
	return helper.GetExpressionValueAsString(ctx, p.format, defFormat)
}

// Execute functin in charge of executing the plugin.
func (p *Print) Execute(ctx context.FeatureContext) errors.Error {
	return actions.Execute(ctx, p.Base, func(ctx context.FeatureContext) errors.Error {
		evalCtx := ctx.EvalContext()
		var err errors.Error
		action := &internal.Print{}
		if action.Message, err = p.Msg(evalCtx); err != nil {
			return err
		}
		if action.Prefix, err = p.Prefix(evalCtx); err != nil {
			return err
		}
		if action.ShowTimestamp, err = p.MustShowTimestamp(evalCtx); err != nil {
			return err
		}
		if action.TimestampFormat, err = p.TimestampFormat(evalCtx); err != nil {
			return err
		}
		if action.Format, err = p.Format(evalCtx); err != nil {
			return err
		}
		log.Debugf("[%s] It starts the execution", BlockPrint)
		log.Tracef("printing message in %s format", action.Format)

		return action.Execute()
	})
}
