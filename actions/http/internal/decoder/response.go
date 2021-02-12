package decoder

import "github.com/hashicorp/hcl/v2"

type Response struct {
	values map[string]hcl.Expression
}

func (res *Response) SetValues(attributes hcl.Attributes) {
	for index := range attributes {
		attribute := attributes[index]
		if res.values == nil {
			res.values = make(map[string]hcl.Expression)
		}
		res.values[attribute.Name] = attribute.Expr
	}
}

func (res *Response) Values() map[string]hcl.Expression {
	return res.values
}
