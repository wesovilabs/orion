package dsl

import (
	"github.com/hashicorp/hcl/v2"

	"github.com/wesovilabs-tools/orion/internal/errors"
	"github.com/zclconf/go-cty/cty"
)

const blockVars = "vars"

type VarsList []Vars

type Vars hcl.Attributes

func (vars Vars) Append(newVars Vars) {
	for name, value := range newVars {
		vars[name] = value
	}
}

func decodeVars(block *hcl.Block) (Vars, errors.Error) {
	attributes, d := block.Body.JustAttributes()
	if err := errors.EvalDiagnostics(d); err != nil {
		return nil, err
	}
	return Vars(attributes), nil
}

func (vars Vars) To(current map[string]cty.Value) errors.Error {
	for name, value := range vars {
		v, d := value.Expr.Value(nil)
		if err := errors.EvalDiagnostics(d); err != nil {
			return err
		}
		current[name] = v
	}
	return nil
}
