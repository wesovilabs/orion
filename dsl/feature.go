package dsl

import (
	"github.com/hashicorp/hcl/v2"
	"github.com/wesovilabs-tools/orion/context"
	"github.com/wesovilabs-tools/orion/internal/errors"
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

// Feature  list of scenarios
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

func (f *Feature) Path() string {
	return f.path
}

func (f *Feature) AllHooks() []*Hook {
	hooks := make([]*Hook, len(f.beforeTag)+len(f.afterTag))
	index := 0
	for _, hook := range f.beforeTag {
		hooks[index] = hook
	}
	for _, hook := range f.afterTag {
		hooks[index] = hook
	}
	if f.beforeEach != nil {
		hooks = append(hooks, f.beforeEach)
	}
	if f.afterEach != nil {
		hooks = append(hooks, f.afterEach)
	}
	return hooks
}

func (f *Feature) SetPath(path string) {
	f.path = path
}

func (f *Feature) Join(include *Feature) {
	if f.afterEach == nil && include.afterEach != nil {
		f.afterEach = include.afterEach
	}
	if f.beforeEach == nil && include.beforeEach != nil {
		f.beforeEach = include.beforeEach
	}
	f.scenarios = append(include.scenarios, f.scenarios...)
	if include.input != nil {
		if f.input != nil {
			f.input.args = append(f.input.args, include.input.args...)
		} else {
			f.input = include.input
		}
	}
	for tag, h := range include.afterTag {
		if _, ok := f.afterTag[tag]; !ok {
			f.afterTag[tag] = h
		}
	}
	for tag, h := range include.beforeTag {
		if _, ok := f.beforeTag[tag]; !ok {
			f.beforeTag[tag] = h
		}
	}
	f.vars.Append(include.vars)

}

func (f *Feature) Description() string {
	return f.description
}

func (f *Feature) Vars() Vars {
	return f.vars
}

func (f *Feature) LoadVariables(ctx context.FeatureContext) errors.Error {
	if f.input != nil && f.input.args != nil {
		return f.input.Execute(ctx)
	}
	return nil
}

func (f *Feature) AddScenario(scenario *Scenario) {
	if f.scenarios == nil {
		f.scenarios = make([]*Scenario, 0)
	}
	f.scenarios = append(f.scenarios, scenario)
}

func (f *Feature) Scenarios() []*Scenario {
	return f.scenarios
}

func (f *Feature) Includes() Includes {
	return f.includes
}

func (f *Feature) Input() *Input {
	return f.input
}

func (f *Feature) BeforeHooksByTag(tags []string) []*Hook {
	tagHooks := f.hooksByTag(f.beforeTag, tags)
	if f.beforeEach != nil {
		return append([]*Hook{f.beforeEach}, tagHooks...)
	}
	return tagHooks
}

func (f *Feature) AfterHooksByTag(tags []string) []*Hook {
	tagHooks := f.hooksByTag(f.afterTag, tags)
	if f.afterEach != nil {
		return append([]*Hook{f.afterEach}, tagHooks...)
	}
	return tagHooks
}

func (f *Feature) hooksByTag(hooks map[string]*Hook, tags []string) []*Hook {
	output := make([]*Hook, 0)
	for index := range tags {
		tag := tags[index]
		if h := hooks[tag]; h != nil {
			output = append(output, h)
		}
	}
	return output
}

func DecodeFeature(body hcl.Body) (*Feature, errors.Error) {
	feature := &Feature{
		scenarios: make([]*Scenario, 0),
		includes:  make(Includes, 0),
		beforeTag: make(Hooks),
		afterTag:  make(Hooks),
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
	var err errors.Error
	for name, blocks := range blocksByType {
		switch name {
		case blockInput:
			if len(blocks) > 1 {
				return errors.IncorrectUsage("only one block 'input' is permitted")
			}
			if feature.input, err = decodeInput(blocks[0]); err != nil {
				return err
			}
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
			if feature.vars, err = decodeVars(blocks[0]); err != nil {
				return err
			}
		default:
			return errors.ThrowUnsupportedBlock("", name)
		}

	}
	return nil
}
