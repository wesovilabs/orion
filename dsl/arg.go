package dsl

import (
	"github.com/hashicorp/hcl/v2"
	"github.com/wesovilabs/orion/context"
	"github.com/wesovilabs/orion/internal/errors"
)

var schemaArg = &hcl.BodySchema{
	Attributes: []hcl.AttributeSchema{
		{Name: argDescription, Required: false},
		{Name: ArgDefault, Required: false},
	},
	Blocks: nil,
}

type Args []*Arg

type Arg struct {
	name        string
	description string
	def         hcl.Expression
}

func (a *Arg) Execute(ctx context.OrionContext) errors.Error {
	if _, ok := ctx.EvalContext().Variables[a.name]; !ok {
		if a.def == nil {
			return errors.IncorrectUsage("variable %s must be provided", a.name)
		}
		argDef, d := a.def.Value(ctx.EvalContext())
		if err := errors.EvalDiagnostics(d); err != nil {
			return err
		}
		ctx.EvalContext().Variables[a.name] = argDef
	}
	return nil
}

// DecodeArg decode bock arg.
func DecodeArg(b *hcl.Block) (*Arg, errors.Error) {
	arg := new(Arg)
	arg.name = b.Labels[0]
	bc, d := b.Body.Content(schemaArg)
	if err := errors.EvalDiagnostics(d); err != nil {
		return nil, err
	}
	if err := arg.populateArgAttributes(bc.Attributes); err != nil {
		return nil, err
	}
	return arg, nil
}

func (a *Arg) populateArgAttributes(attrs hcl.Attributes) errors.Error {
	var err errors.Error
	for name, attribute := range attrs {
		switch name {
		case ArgDefault:
			a.def = attribute.Expr
		case argDescription:
			a.description, err = attributeToStringWithoutContext(attribute)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
