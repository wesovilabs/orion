package decoder

import (
	"time"

	"github.com/wesovilabs/orion/helper"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/hashicorp/hcl/v2"

	"github.com/wesovilabs/orion/internal/errors"
)

const ()

var (
	defConnectionTimeout = 10 * time.Second
	defConnectionURI     = "mongodb://localhost:27017"
)

type Connection struct {
	uri     hcl.Expression
	timeout hcl.Expression
	auth    *Auth
}

func (c *Connection) SetURI(expr hcl.Expression) {
	c.uri = expr
}

func (c *Connection) SetTimeout(expr hcl.Expression) {
	c.timeout = expr
}

func (c *Connection) SetAuth(auth *Auth) {
	c.auth = auth
}

func (c *Connection) Timeout(ctx *hcl.EvalContext) (*time.Duration, errors.Error) {
	return helper.GetExpressionValueAsDuration(ctx, c.timeout, &defConnectionTimeout)
}

func (c *Connection) ClientOpts(ctx *hcl.EvalContext) (*options.ClientOptions, errors.Error) {
	credential, err := c.auth.mongoCredentials(ctx)
	if err != nil {
		return nil, err
	}
	uri, err := helper.GetExpressionValueAsString(ctx, c.uri, defConnectionURI)
	if err != nil {
		return nil, err
	}
	return options.Client().ApplyURI(uri).SetAuth(credential), nil
}
