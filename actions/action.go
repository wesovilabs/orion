package actions

import (
	"github.com/hashicorp/hcl/v2"
	log "github.com/sirupsen/logrus"
	"github.com/wesovilabs/orion/helper"
	"github.com/wesovilabs/orion/internal/errors"

	"github.com/wesovilabs/orion/context"
	"github.com/zclconf/go-cty/cty"
)

type Actions []Action

// Action interface to be implemented by any action.
type Action interface {
	SetDesc(string)
	SetWhen(hcl.Expression)
	SetCount(hcl.Expression)
	SetWhile(hcl.Expression)
	SetItems(hcl.Expression)
	SetKind(string)
	Description() string
	String() string
	ShouldExecute(ctx *hcl.EvalContext) bool
	Execute(ctx context.OrionContext) errors.Error
}

// Base contain the common attributes of all the plugins.
type Base struct {
	kind  string
	desc  string
	when  hcl.Expression
	while hcl.Expression
	count hcl.Expression
	items hcl.Expression
}

// SetDesc set description for the block.
func (s *Base) SetDesc(desc string) {
	s.desc = desc
}

// Description return block description.
func (s *Base) Description() string {
	return s.desc
}

// SetKind set attribute kind.
func (s *Base) SetKind(kind string) {
	s.kind = kind
}

// SetWhen set attribute when.
func (s *Base) SetWhen(when hcl.Expression) {
	s.when = when
}

// SetWhile set attribute while.
func (s *Base) SetWhile(while hcl.Expression) {
	s.while = while
}

// SetCount set attribute count.
func (s *Base) SetCount(count hcl.Expression) {
	s.count = count
}

// SetCount set attribute count.
func (s *Base) SetItems(items hcl.Expression) {
	s.items = items
}

// String default method String.
func (s *Base) String() string {
	return s.kind
}

// SetBaseArgs set the common arguments.
func SetBaseArgs(action Action, attribute *hcl.Attribute) errors.Error {
	switch attribute.Name {
	case ArgWhen:
		action.SetWhen(attribute.Expr)
	case ArgWhile:
		action.SetWhile(attribute.Expr)
	case ArgCount:
		action.SetCount(attribute.Expr)
	case ArgDesc:
		desc, err := helper.AttributeToStringWithoutContext(attribute)
		if err != nil {
			return err
		}
		action.SetDesc(desc)
	}
	return nil
}

func (s *Base) ShouldExecute(ctx *hcl.EvalContext) bool {
	if s.when == nil {
		return true
	}
	v, d := s.when.Value(ctx)
	if err := errors.EvalDiagnostics(d); err != nil {
		log.Warn(err)
		return false
	}
	if v.Type() == cty.Bool {
		return v.True()
	}
	log.Warningf("when attribute must be boolean")
	return false
}
