package dsl

import (
	"github.com/hashicorp/hcl/v2"
	"github.com/wesovilabs/orion/context"
	"github.com/wesovilabs/orion/internal/errors"
)

var schemaFeature = &hcl.BodySchema{
	Blocks: []hcl.BlockHeaderSchema{
		{Type: blockInput},
		{Type: blockScenario, LabelNames: []string{labelName}},
		{Type: blockHookBefore, LabelNames: []string{labelTag}},
		{Type: blockHookAfter, LabelNames: []string{labelTag}},
		{Type: blockVars},
		{Type: blockFunc, LabelNames: []string{labelName}},
	},
	Attributes: []hcl.AttributeSchema{
		{Name: argDescription, Required: false},
		{Name: argIncludes, Required: false},
	},
}

// Feature  list of scenarios.
type Feature struct {
	path        string
	description string
	includes    Includes
	input       *Input
	scenarios   []*Scenario
	beforeEach  *Hook
	afterEach   *Hook
	beforeTag   map[string]*Hook
	afterTag    map[string]*Hook
	functions   Functions
	vars        Vars
}

func (feature *Feature) Path() string {
	return feature.path
}

func (feature *Feature) SetPath(path string) {
	feature.path = path
}

func (feature *Feature) Join(include *Feature) {
	if feature.afterEach == nil && include.afterEach != nil {
		feature.afterEach = include.afterEach
	}
	if feature.beforeEach == nil && include.beforeEach != nil {
		feature.beforeEach = include.beforeEach
	}
	feature.scenarios = append(include.scenarios, feature.scenarios...)
	if include.input != nil {
		if feature.input != nil {
			feature.input.args = append(feature.input.args, include.input.args...)
		} else {
			feature.input = include.input
		}
	}
	for tag, h := range include.afterTag {
		if _, ok := feature.afterTag[tag]; !ok {
			feature.afterTag[tag] = h
		}
	}
	for tag, h := range include.beforeTag {
		if _, ok := feature.beforeTag[tag]; !ok {
			feature.beforeTag[tag] = h
		}
	}
	feature.vars.Append(include.vars)
	feature.functions.Append(include.functions)
}

func (feature *Feature) Description() string {
	return feature.description
}

func (feature *Feature) Vars() Vars {
	return feature.vars
}

func (feature *Feature) LoadVariables(ctx context.OrionContext) errors.Error {
	if feature.input != nil && feature.input.args != nil {
		return feature.input.Execute(ctx)
	}
	return nil
}

func (feature *Feature) AddScenario(scenario *Scenario) {
	if feature.scenarios == nil {
		feature.scenarios = make([]*Scenario, 0)
	}
	feature.scenarios = append(feature.scenarios, scenario)
}

func (feature *Feature) Scenarios() []*Scenario {
	return feature.scenarios
}

func (feature *Feature) Includes() Includes {
	return feature.includes
}

func (feature *Feature) Input() *Input {
	return feature.input
}

func (feature *Feature) BeforeHooksByTag(tags []string) []*Hook {
	tagHooks := feature.hooksByTag(feature.beforeTag, tags)
	if feature.beforeEach != nil {
		return append([]*Hook{feature.beforeEach}, tagHooks...)
	}
	return tagHooks
}

func (feature *Feature) AfterHooksByTag(tags []string) []*Hook {
	tagHooks := feature.hooksByTag(feature.afterTag, tags)
	if feature.afterEach != nil {
		return append([]*Hook{feature.afterEach}, tagHooks...)
	}
	return tagHooks
}

func (feature *Feature) hooksByTag(hooks map[string]*Hook, tags []string) []*Hook {
	output := make([]*Hook, 0)
	for index := range tags {
		tag := tags[index]
		if h := hooks[tag]; h != nil {
			output = append(output, h)
		}
	}
	return output
}

func (feature *Feature) Functions() map[string]func(context.OrionContext, string) errors.Error {
	output := make(map[string]func(context.OrionContext, string) errors.Error)
	for index := range feature.functions {
		function := feature.functions[index]
		output[function.name] = function.runFunction
	}
	return output
}

func DecodeFeature(body hcl.Body) (*Feature, errors.Error) {
	feature := &Feature{
		scenarios: make([]*Scenario, 0),
		includes:  make(Includes, 0),
		beforeTag: make(Hooks),
		afterTag:  make(Hooks),
		functions: make(Functions),
		vars:      make(Vars),
	}
	bc, d := body.Content(schemaFeature)
	if err := errors.EvalDiagnostics(d); err != nil {
		return nil, err
	}

	if err := feature.populateAttributes(bc.Attributes); err != nil {
		return nil, err
	}
	if err := feature.populateBlocks(bc.Blocks.ByType()); err != nil {
		return nil, err
	}
	return feature, nil
}

func (feature *Feature) populateAttributes(attributes hcl.Attributes) errors.Error {
	var err errors.Error
	for name, attr := range attributes {
		switch name {
		case argDescription:
			if feature.description, err = attributeToStringWithoutContext(attr); err != nil {
				return err
			}
		case argIncludes:
			if feature.includes, err = includesFromValue(attr); err != nil {
				return err
			}
		default:
			return errors.ThrowUnsupportedArgument("", name)
		}
	}
	return nil
}

func (feature *Feature) populateBlocks(blocksByType map[string]hcl.Blocks) errors.Error {
	for name, blocks := range blocksByType {
		switch name {
		case blockInput:
			if len(blocks) > 1 {
				return errors.IncorrectUsage("only one block 'input' is permitted")
			}
			input, err := decodeInput(blocks[0])
			if err != nil {
				return err
			}
			feature.input = input
		case blockScenario:
			for i := range blocks {
				scenario, err := decodeScenario(blocks[i])
				if err != nil {
					return err
				}
				feature.scenarios = append(feature.scenarios, scenario)
			}
		case blockHookBefore:
			for i := range blocks {
				hook, err := decodeHook(blocks[i])
				if err != nil {
					return err
				}
				hook.kind = HookBefore
				if hook.tag == hookEach {
					feature.beforeEach = hook
					continue
				}
				feature.beforeTag[hook.tag] = hook
			}
		case blockHookAfter:
			for i := range blocks {
				hook, err := decodeHook(blocks[i])
				if err != nil {
					return err
				}
				hook.kind = HookAfter
				if hook.tag == hookEach {
					feature.afterEach = hook
					continue
				}
				feature.afterTag[hook.tag] = hook
			}
		case blockFunc:
			feature.functions = make(Functions)
			for i := range blocks {
				function, err := decodeFunc(blocks[i])
				if err != nil {
					return err
				}
				feature.functions[function.name] = function
			}
		case blockVars:
			if len(blocks) > 1 {
				return errors.IncorrectUsage("only one block '%s' is permitted", blockVars)
			}
			vars, err := decodeVars(blocks[0])
			if err != nil {
				return err
			}
			feature.vars = vars
		default:
			return errors.ThrowUnsupportedBlock("", name)
		}
	}
	return nil
}
