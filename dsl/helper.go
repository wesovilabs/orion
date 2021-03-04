package dsl

import (
	"github.com/hashicorp/hcl/v2"
	"github.com/wesovilabs/orion/actions/register"

	"github.com/wesovilabs/orion/helper"
	"github.com/wesovilabs/orion/internal/errors"
	"github.com/zclconf/go-cty/cty"
)

var handler = register.GetHandler()

func attributeToStringWithoutContext(attribute *hcl.Attribute) (string, errors.Error) {
	if len(attribute.Expr.Variables()) > 0 {
		return "", errors.IncorrectUsage("variables are not permitted in attribute %s", attribute.Name)
	}
	val, err := helper.EvalAttribute(nil, attribute)
	if err != nil {
		return "", err
	}
	return helper.ToStrictString(val)
}

func attributeToSliceWithoutContext(attribute *hcl.Attribute) ([]cty.Value, errors.Error) {
	if len(attribute.Expr.Variables()) > 0 {
		return nil, errors.IncorrectUsage("variables are not permitted in attribute %s", attribute.Name)
	}
	value, err := helper.EvalAttribute(nil, attribute)
	if err != nil {
		return nil, err
	}
	if value.Type().IsListType() || value.Type().IsTupleType() || value.Type().IsCollectionType() {
		return value.AsValueSlice(), nil
	}
	return nil, errors.InvalidArguments("expected a slice argument but it's not")
}
