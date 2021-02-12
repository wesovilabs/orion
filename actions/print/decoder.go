// Package pprint contains the root elements required by the plugin
package pprint

import (
	"github.com/hashicorp/hcl/v2"
	"github.com/wesovilabs/orion/actions"
	"github.com/wesovilabs/orion/internal/errors"
)

const (
	// BlockPrint plugin block identifier.
	BlockPrint = "print"
	// AttributePrefix name of the argument used in the block print.
	AttributePrefix = "prefix"
	// AttributeMsg name of the argument used in the block print.
	AttributeMsg = "msg"
	// AttributeTimestamp name of the argument used in the block print.
	AttributeTimestamp = "showTimestamp"
	// AttributeTimestampFormat name of the argument used in the block print.
	AttributeTimestampFormat = "timestampFormat"
	// AttributeFormat name of the argument used in the block print.
	AttributeFormat = "format"
)

// nolint: gochecknoglobals
var schemaPrint = &hcl.BodySchema{
	Attributes: append(actions.BaseArguments, []hcl.AttributeSchema{
		{Name: AttributeMsg, Required: true},
		{Name: AttributePrefix, Required: false},
		{Name: AttributeFormat, Required: false},
		{Name: AttributeTimestampFormat, Required: false},
		{Name: AttributeTimestamp, Required: false},
	}...),
}

// Decoder struct used to implement required interfaces.
type Decoder struct{}

// BlockHeaderSchema return the header schema for the plugin.
func (dec *Decoder) BlockHeaderSchema() hcl.BlockHeaderSchema {
	return hcl.BlockHeaderSchema{
		Type:       BlockPrint,
		LabelNames: nil,
	}
}

// DecodeBlock required to implement the plugin interface.
func (dec *Decoder) DecodeBlock(block *hcl.Block) (actions.Action, errors.Error) {
	bodyContent, body, d := block.Body.PartialContent(schemaPrint)

	if err := errors.EvalDiagnostics(d); err != nil {
		return nil, err
	}

	attributes, d := body.JustAttributes()

	if err := errors.EvalDiagnostics(d); err != nil {
		return nil, err
	}
	p := &Print{Base: &actions.Base{}}

	if err := populateAttributes(p, attributes); err != nil {
		return nil, err
	}

	if err := populateAttributes(p, bodyContent.Attributes); err != nil {
		return nil, err
	}

	if len(block.Labels) != 1 {
		return nil, errors.ThrowMissingRequiredLabel(BlockPrint)
	}

	return p, nil
}

func populateAttributes(p *Print, attrs hcl.Attributes) errors.Error {
	for name := range attrs {
		attribute := attrs[name]

		switch {
		case actions.IsPluginBaseArgument(name):
			if err := actions.SetBaseArgs(p, attribute); err != nil {
				return err
			}
		case name == AttributePrefix:
			p.SetPrefix(attribute.Expr)
		case name == AttributeMsg:
			p.SetMessage(attribute.Expr)
		case name == AttributeTimestamp:
			p.SetTimestamp(attribute.Expr)
		case name == AttributeFormat:
			p.SetFormat(attribute.Expr)
		case name == AttributeTimestampFormat:
			p.SetTimestampFormat(attribute.Expr)
		default:
			return errors.ThrowUnsupportedArgument(BlockPrint, name)
		}
	}

	return nil
}
