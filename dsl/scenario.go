package dsl

import (
	"fmt"

	"github.com/wesovilabs/orion/context"

	"github.com/zclconf/go-cty/cty"

	"github.com/hashicorp/hcl/v2"

	"github.com/wesovilabs/orion/helper"

	log "github.com/sirupsen/logrus"
	"github.com/wesovilabs/orion/internal/errors"
)

const minNumberOfSections = 2

var schemaScenario = &hcl.BodySchema{
	Blocks: []hcl.BlockHeaderSchema{
		{Type: blockGiven, LabelNames: []string{labelDescription}},
		{Type: blockWhen, LabelNames: []string{labelDescription}},
		{Type: blockThen, LabelNames: []string{labelDescription}},
	},
	Attributes: []hcl.AttributeSchema{
		{Name: argTags, Required: false},
		{Name: argContinueOnError, Required: false},
		{Name: argIgnore, Required: false},
		{Name: argExamples, Required: false},
	},
}

// Scenario represents a single schemaScenario to be tested.
// A schemaScenario is composed by the Given-Block-Then blocks.
type Scenario struct {
	description     string
	sections        Sections
	tags            []string
	examples        hcl.Expression
	continueOnError hcl.Expression
	ignore          hcl.Expression
}

func (s *Scenario) ContinueOnError(ctx *hcl.EvalContext) bool {
	if s.continueOnError == nil {
		return false
	}
	v, d := s.continueOnError.Value(ctx)
	if err := errors.EvalDiagnostics(d); err != nil {
		log.Warning(err)
		return false
	}
	return v.True()
}

func (s *Scenario) Validate() errors.Error {
	log.Debugf("It validates the %d sections in the scenario:", len(s.sections))
	if len(s.sections) < minNumberOfSections {
		return errors.IncorrectUsage("A scenario must contains 2 sections at least")
	}
	if s.sections[0].Name() == blockThen {
		return errors.IncorrectUsage("Initial section in a scenario must be 'given' or 'then'.")
	}
	var lastSection string
	for index := range s.sections {
		section := s.sections[index]
		if lastSection != "" {
			if section.Name() == blockGiven {
				return errors.IncorrectUsage("section 'given' can only be the initial section ina  scenario.")
			}
			if section.Name() == blockWhen && lastSection == blockWhen {
				return errors.IncorrectUsage("section 'when' can only proceed sections 'given' or 'then'.")
			}
			if section.Name() == blockThen && lastSection != blockWhen {
				return errors.IncorrectUsage("section 'then' can only proceed section 'when'.")
			}
		}
		lastSection = section.Name()
	}
	return nil
}

func (s *Scenario) String() string {
	return fmt.Sprintf("[scenario] %s ", s.description)
}

func (s *Scenario) Tags() []string {
	return s.tags
}

func (s *Scenario) Sections() Sections {
	return s.sections
}

func (s *Scenario) Examples(ctx context.FeatureContext) ([]map[string]cty.Value, errors.Error) {
	if s.examples == nil {
		return nil, nil
	}
	val, d := s.examples.Value(ctx.EvalContext())
	if err := errors.EvalDiagnostics(d); err != nil {
		return nil, err
	}
	if !helper.IsSlice(val) {
		return nil, errors.IncorrectUsage("examples must contain an array of entries")
	}
	slice := val.AsValueSlice()
	output := make([]map[string]cty.Value, len(slice))
	for index := range slice {
		row := slice[index]
		if helper.IsMap(row) {
			output[index] = row.AsValueMap()
			continue
		}
		return nil, errors.IncorrectUsage("unsupported example record")
	}
	return output, nil
}

func (s *Scenario) Execute(ctx context.FeatureContext) errors.Error {
	ctx.StartScenario()
	if err := s.Validate(); err != nil {
		return err
	}
	log.Debug("It starts the execution")
	for index := range s.Sections() {
		section := s.Sections()[index]
		log.Info(section)
		if err := section.executeActions(ctx); err != nil {
			ctx.FailScenario()
			return err
		}
	}
	ctx.StopScenario()
	log.Debug("The s execution was completed successfully.")
	return nil
}

func (s *Scenario) IsIgnored(ctx *hcl.EvalContext) bool {
	if s.ignore == nil {
		return false
	}
	value, d := s.ignore.Value(ctx)
	if err := errors.EvalDiagnostics(d); err != nil {
		log.Warning(err)
		return false
	}
	return value.True()
}

func decodeScenario(b *hcl.Block) (*Scenario, errors.Error) {
	scenario := &Scenario{
		description: b.Labels[0],
		sections:    make(Sections, 0),
		tags:        make([]string, 0),
	}
	bc, d := b.Body.Content(schemaScenario)
	if err := errors.EvalDiagnostics(d); err != nil {
		return nil, err
	}
	if err := scenario.populateAttributes(bc.Attributes); err != nil {
		return nil, err
	}
	if err := scenario.populateSections(bc.Blocks); err != nil {
		return nil, err
	}
	return scenario, nil
}

func (s *Scenario) populateAttributes(attrs hcl.Attributes) errors.Error {
	for name, attribute := range attrs {
		switch name {
		case argTags:
			tagsValue, err := attributeToSliceWithoutContext(attribute)
			if err != nil {
				return err
			}
			for index := range tagsValue {
				tag, err := helper.ToStrictString(tagsValue[index])
				if err != nil {
					return err
				}
				s.tags = append(s.tags, tag)
			}
		case argContinueOnError:
			s.continueOnError = attribute.Expr
		case argIgnore:
			s.ignore = attribute.Expr
		case argExamples:
			s.examples = attribute.Expr
		default:
			return errors.ThrowUnsupportedArgument(blockScenario, name)
		}
	}

	return nil
}

func (s *Scenario) populateSections(blocks hcl.Blocks) errors.Error {
	for index := range blocks {
		block := blocks[index]
		var section Section
		var err errors.Error
		switch block.Type {
		case blockGiven:
			section, err = decodeGiven(block)
		case blockWhen:
			section, err = decodeWhen(block)
		case blockThen:
			section, err = decodeThen(block)
		default:
			return errors.ThrowUnsupportedBlock(blockScenario, block.Type)
		}
		if err != nil {
			return err
		}
		s.sections = append(s.sections, section)
	}
	return nil
}
