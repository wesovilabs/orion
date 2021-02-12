package http

import (
	"github.com/hashicorp/hcl/v2"
	"github.com/wesovilabs-tools/orion/actions"
	"github.com/wesovilabs-tools/orion/actions/http/internal/decoder"
	"github.com/wesovilabs-tools/orion/internal/errors"
)

const (
	BlockHTTP     = "http"
	BlockRequest  = "request"
	BlockResponse = "response"
	LabelMethod   = "method"
)

var schemaHTTP = &hcl.BodySchema{
	Attributes: append(actions.BaseArguments, []hcl.AttributeSchema{}...),
	Blocks: []hcl.BlockHeaderSchema{
		{Type: BlockRequest},
		{Type: BlockResponse},
	},
}

type Decoder struct{}

// BlockHeaderSchema return the header schema for the plugin
func (dec *Decoder) BlockHeaderSchema() hcl.BlockHeaderSchema {
	return hcl.BlockHeaderSchema{
		Type:       BlockHTTP,
		LabelNames: []string{LabelMethod},
	}
}

func (dec *Decoder) DecodeBlock(block *hcl.Block) (actions.Action, errors.Error) {
	bodyContent, d := block.Body.Content(schemaHTTP)
	if err := errors.EvalDiagnostics(d); err != nil {
		return nil, err
	}
	http := &HTTP{Base: &actions.Base{}}
	http.SetMethod(block.Labels[0])
	if err := populateAttributes(http, bodyContent.Attributes); err != nil {
		return nil, err
	}
	if err := populateBlocks(http, bodyContent.Blocks); err != nil {
		return nil, err
	}
	return http, nil
}

func populateAttributes(http *HTTP, attrs hcl.Attributes) errors.Error {
	for name := range attrs {
		attribute := attrs[name]
		switch {
		case actions.IsPluginBaseArgument(name):
			if err := actions.SetBaseArgs(http, attribute); err != nil {
				return err
			}
		default:
			return errors.ThrowUnsupportedArgument(BlockHTTP, name)
		}
	}
	return nil
}

func populateBlocks(http *HTTP, blocks hcl.Blocks) errors.Error {
	blocksByType := blocks.ByType()
	for blockType, blockList := range blocksByType {
		switch {
		case blockType == BlockRequest:
			if len(blockList) > 1 {
				// throw error
			}
			req, err := decoder.DecodeRequest(blockList[0])
			if err != nil {
				return err
			}
			http.SetRequest(req)
		case blockType == BlockResponse:
			if len(blockList) > 1 {
				// throw error
			}
			res, err := decoder.DecodeResponse(blockList[0])
			if err != nil {
				return err
			}
			http.SetResponse(res)
		default:
			return errors.ThrowUnsupportedBlock(BlockHTTP, blockType)
		}
	}
	return nil
}
