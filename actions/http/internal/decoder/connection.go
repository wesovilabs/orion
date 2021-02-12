package decoder

import (
	"time"

	"github.com/hashicorp/hcl/v2"
	"github.com/wesovilabs/orion/helper"
	"github.com/wesovilabs/orion/internal/errors"
)

var (
	defConnectionTimeout = 10 * time.Second
	defConnectionProxy   = ""
)

type Connection struct {
	timeout hcl.Expression
	proxy   hcl.Expression
}

func (c *Connection) SetTimeout(expr hcl.Expression) {
	c.timeout = expr
}

func (c *Connection) Timeout(ctx *hcl.EvalContext) (time.Duration, errors.Error) {
	timeDuration, err := helper.GetExpressionValueAsDuration(ctx, c.timeout, &defConnectionTimeout)
	return *timeDuration, err
}

func (c *Connection) SetProxy(expr hcl.Expression) {
	c.proxy = expr
}

func (c *Connection) Proxy(ctx *hcl.EvalContext) (string, errors.Error) {
	return helper.GetExpressionValueAsString(ctx, c.proxy, defConnectionProxy)
}
