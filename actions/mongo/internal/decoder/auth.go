package decoder

import (
	"strings"

	"github.com/hashicorp/hcl/v2"
	mngHelper "github.com/wesovilabs/orion/actions/mongo/internal/helper"
	"github.com/wesovilabs/orion/helper"
	"github.com/wesovilabs/orion/internal/errors"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	defUsername        = "test"
	defPassword        = ""
	mechanismSha1      = "scram-sha-1"
	mechanismSha256    = "scram-sha-256"
	mechanismMongoDBCr = "mongodb-cr"
	mechanismPlain     = "plain"
	mechanismGssapi    = "gssapi"
	mechanismX509      = "mongodb-x509"
	mechanismAWS       = "mongodb-aws"
)

var (
	supportedAuthMechanisms = map[string]struct{}{
		mechanismSha1:      {},
		mechanismSha256:    {},
		mechanismMongoDBCr: {},
		mechanismPlain:     {},
		mechanismGssapi:    {},
		mechanismX509:      {},
		mechanismAWS:       {},
	}
	supportedAuthMechanismsName = mngHelper.MapStructToArray(supportedAuthMechanisms)
)

type Auth struct {
	mechanism string
	username  hcl.Expression
	password  hcl.Expression
}

func (auth *Auth) SetUsername(expr hcl.Expression) {
	auth.username = expr
}

func (auth *Auth) SetPassword(expr hcl.Expression) {
	auth.password = expr
}

func (auth *Auth) Mechanism() string {
	return strings.ToUpper(auth.mechanism)
}

func (auth *Auth) Username(ctx *hcl.EvalContext) (string, errors.Error) {
	return helper.GetExpressionValueAsString(ctx, auth.username, defUsername)
}

func (auth *Auth) Password(ctx *hcl.EvalContext) (string, errors.Error) {
	return helper.GetExpressionValueAsString(ctx, auth.password, defPassword)
}

func (auth *Auth) mongoCredentials(ctx *hcl.EvalContext) (options.Credential, errors.Error) {
	mechanism := auth.Mechanism()
	username, err := auth.Username(ctx)
	if err != nil {
		return options.Credential{}, err
	}
	password, err := auth.Password(ctx)
	if err != nil {
		return options.Credential{}, err
	}
	return options.Credential{
		AuthMechanism: mechanism,
		Username:      username,
		Password:      password,
	}, nil
}
