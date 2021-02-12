package decoder

import (
	"github.com/hashicorp/hcl/v2"
	"github.com/wesovilabs-tools/orion/helper"
	"github.com/wesovilabs-tools/orion/internal/errors"
)

type BlockProperties struct {
	values map[string]hcl.Expression
}

func (b *BlockProperties) SetValues(attributes hcl.Attributes) {
	for index := range attributes {
		attribute := attributes[index]
		if b.values == nil {
			b.values = make(map[string]hcl.Expression)
		}
		b.values[attribute.Name] = attribute.Expr
	}
}

func (b *BlockProperties) Values(ctx *hcl.EvalContext) (map[string]interface{}, errors.Error) {
	out := make(map[string]interface{})
	for name, val := range b.values {
		value, err := helper.GetExpressionValueAsInterface(ctx, val, nil)
		if err != nil {
			return nil, err
		}
		out[name] = value
	}
	return out, nil
}

func (b BlockProperties) Evaluate(ctx *hcl.EvalContext) errors.Error {
	return helper.EvaluateExpressions(ctx, b.values)
}

func NewBlockProperties(block *hcl.Block) (*BlockProperties, errors.Error) {
	b := new(BlockProperties)
	attributes, d := block.Body.JustAttributes()
	if err := errors.EvalDiagnostics(d); err != nil {
		return nil, err
	}
	b.SetValues(attributes)
	return b, nil
}
