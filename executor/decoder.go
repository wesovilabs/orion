package executor

import (
	"github.com/hashicorp/hcl/v2"
	"github.com/wesovilabs/orion/dsl"
	"github.com/wesovilabs/orion/internal/errors"
)

type Decoder interface {
	Decode(body hcl.Body) (*dsl.Feature, errors.Error)
}

type decoder struct {
}

func (dec *decoder) Decode(body hcl.Body) (*dsl.Feature, errors.Error) {
	return dsl.DecodeFeature(body)
}
