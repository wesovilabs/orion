package dsl

import (
	"fmt"

	actions2 "github.com/wesovilabs-tools/orion/actions"
	"github.com/wesovilabs-tools/orion/context"

	"github.com/hashicorp/hcl/v2"

	"github.com/wesovilabs-tools/orion/internal/errors"
)

var schemaHook = &hcl.BodySchema{
	Attributes: []hcl.AttributeSchema{
		{Name: argDescription, Required: false},
	},
	Blocks: handler.GetBlocksSpec(),
}

const (
	HookBefore        = "before"
	HookAfter         = "after"
	hookEach   string = "each"
)

type Hooks map[string]*Hook

type Hook struct {
	kind        string
	description string
	tag         string
	actions     actions2.Actions
}

func (h *Hook) TotalActions() int {
	return len(h.actions)
}

func (h *Hook) String() string {
	if h.description == "" {
		return fmt.Sprintf("[%s|%s]", h.tag, h.kind)
	}
	return fmt.Sprintf("[%s|%s]: %s", h.tag, h.kind, h.description)
}

func (h *Hook) Execute(ctx context.FeatureContext) errors.Error {
	for index := range h.actions {
		action := h.actions[index]
		if err := action.Execute(ctx); err != nil {
			return err
		}
	}
	return nil
}

func decodeHook(block *hcl.Block) (*Hook, errors.Error) {
	if len(block.Labels) != 1 {
		return nil, errors.ThrowMissingRequiredLabel(block.Type)
	}
	hook := &Hook{
		tag: block.Labels[0],
	}
	bc, d := block.Body.Content(schemaHook)
	if err := errors.EvalDiagnostics(d); err != nil {
		return nil, err
	}
	if err := hook.populateAttributes(bc.Attributes); err != nil {
		return nil, err
	}
	if err := hook.populateActions(bc.Blocks); err != nil {
		return nil, err
	}
	return hook, nil
}

func (h *Hook) populateAttributes(attrs hcl.Attributes) errors.Error {
	var err errors.Error
	for name, attribute := range attrs {
		switch name {
		case argDescription:
			if h.description, err = attributeToStringWithoutContext(attribute); err != nil {
				return err
			}
		default:
			return errors.ThrowUnsupportedArgument("hook", name)
		}
	}
	return nil
}

func (h *Hook) populateActions(blocks hcl.Blocks) (err errors.Error) {
	h.actions, err = handler.DecodePlugins(blocks)
	return
}
