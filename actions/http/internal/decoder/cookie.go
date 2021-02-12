package decoder

import (
	"github.com/hashicorp/hcl/v2"
	"github.com/wesovilabs-tools/orion/helper"
	"github.com/wesovilabs-tools/orion/internal/errors"
)

const (
	defCookieName   = ""
	defCookieValue  = ""
	defCookiePath   = ""
	defCookieDomain = ""
)

type Cookie struct {
	name   hcl.Expression
	value  hcl.Expression
	path   hcl.Expression
	domain hcl.Expression
}

func (c *Cookie) SetName(expr hcl.Expression) {
	c.name = expr
}

func (c *Cookie) Name(ctx *hcl.EvalContext) (string, errors.Error) {
	return helper.GetExpressionValueAsString(ctx, c.name, defCookieName)
}

func (c *Cookie) SetValue(expr hcl.Expression) {
	c.value = expr
}

func (c *Cookie) Value(ctx *hcl.EvalContext) (string, errors.Error) {
	return helper.GetExpressionValueAsString(ctx, c.value, defCookieValue)
}

func (c *Cookie) SetPath(expr hcl.Expression) {
	c.path = expr
}

func (c *Cookie) Path(ctx *hcl.EvalContext) (string, errors.Error) {
	return helper.GetExpressionValueAsString(ctx, c.path, defCookiePath)
}

func (c *Cookie) SetDomain(expr hcl.Expression) {
	c.domain = expr
}

func (c *Cookie) Domain(ctx *hcl.EvalContext) (string, errors.Error) {
	return helper.GetExpressionValueAsString(ctx, c.domain, defCookieDomain)
}
