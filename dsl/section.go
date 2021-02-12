package dsl

import (
	"fmt"
	"strings"

	actions2 "github.com/wesovilabs-tools/orion/actions"
	"github.com/wesovilabs-tools/orion/context"

	log "github.com/sirupsen/logrus"

	"github.com/hashicorp/hcl/v2"
	"github.com/wesovilabs-tools/orion/internal/errors"
)

type Sections []Section

func (s section) TotalActions() int {
	return len(s.actions)
}

type Section interface {
	Name() string
	String() string
	SetName(name string)
	executeActions(ctx context.FeatureContext) errors.Error
	TotalActions() int
}

type section struct {
	name        string
	description string
	actions     actions2.Actions
}

func (s *section) SetName(name string) {
	s.name = name
}

func (s *section) Name() string {
	return s.name
}

func (s *section) SetDescription(description string) {
	s.description = description
}

func (s *section) Description() string {
	return s.description
}

func (s *section) String() string {
	if s.description != "" {
		return fmt.Sprintf("%s %s", capitalize(s.name), s.description)
	}
	return strings.ToTitle(s.name)
}

func newSection(name string, block *hcl.Block) *section {
	s := &section{
		name:    name,
		actions: make(actions2.Actions, 0),
	}
	if len(block.Labels) == 1 {
		s.description = block.Labels[0]
	}
	return s
}

func capitalize(value string) string {
	return strings.ToUpper(string(value[0])) + value[1:]
}

func (s *section) populateActions(blocks hcl.Blocks) errors.Error {
	for index := range blocks {
		block := blocks[index]
		action, err := handler.DecodePlugin(block)
		if err != nil {
			return err
		}
		action.SetKind(block.Type)
		s.actions = append(s.actions, action)
	}
	return nil
}

func (s *section) executeActions(ctx context.FeatureContext) errors.Error {
	if len(s.actions) == 0 {
		return errors.IncorrectUsage("block '%s' cannot be empty. It must contain one action at least", s.name)
	}
	for index := range s.actions {
		action := s.actions[index]
		if action.ShouldExecute(ctx.EvalContext()) {
			if err := action.Execute(ctx); err != nil {
				return err
			}
			continue
		}
		log.Debugf("action %s is skipped!", action)
	}
	return nil
}
