package mongo

import (
	"github.com/hashicorp/hcl/v2"
	log "github.com/sirupsen/logrus"
	"github.com/wesovilabs/orion/actions/mongo/internal/executor"

	"github.com/wesovilabs/orion/actions"
	"github.com/wesovilabs/orion/actions/mongo/internal/decoder"
	"github.com/wesovilabs/orion/actions/mongo/internal/helper"
	"github.com/wesovilabs/orion/internal/errors"
)

const (
	BlockMongo     = "mongo"
	labelOperation = "operation"
)

var supportedOperationsName = helper.MapStructToArray(supportedOperations)

var schemaMongo = &hcl.BodySchema{
	Attributes: actions.BaseArguments,
	Blocks: []hcl.BlockHeaderSchema{
		{Type: decoder.BlockConnection},
		{Type: decoder.BlockQuery},
		{Type: decoder.BlockResponse},
	},
}

type Decoder struct{}

// BlockHeaderSchema return the header schema for the plugin.
func (dec *Decoder) BlockHeaderSchema() hcl.BlockHeaderSchema {
	return hcl.BlockHeaderSchema{
		Type:       BlockMongo,
		LabelNames: []string{labelOperation},
	}
}

func (dec *Decoder) DecodeBlock(block *hcl.Block) (actions.Action, errors.Error) {
	log.Tracef("it starts decoding of block %s", BlockMongo)
	bodyContent, d := block.Body.Content(schemaMongo)
	if err := errors.EvalDiagnostics(d); err != nil {
		return nil, err
	}
	mongo := &Mongo{Base: &actions.Base{}}
	if len(block.Labels) != 1 {
		return nil, errors.ThrowMissingRequiredLabel(BlockMongo)
	}
	operation := block.Labels[0]
	if _, ok := supportedOperations[operation]; !ok {
		return nil, errors.IncorrectUsage("operation '%s' is not valid. Supported operations are: %v",
			operation, supportedOperationsName)
	}
	mongo.SetOperation(block.Labels[0])
	if err := populateMongoAttributes(mongo, bodyContent.Attributes); err != nil {
		return nil, err
	}
	if err := populateMongoBlocks(mongo, bodyContent.Blocks); err != nil {
		return nil, err
	}
	log.Tracef("block %s is decoded successfully", BlockMongo)
	return mongo, nil
}

func populateMongoAttributes(mongo *Mongo, attrs hcl.Attributes) errors.Error {
	for name := range attrs {
		attribute := attrs[name]
		switch {
		case actions.IsPluginBaseArgument(name):
			if err := actions.SetBaseArgs(mongo, attribute); err != nil {
				return err
			}
		default:
			return errors.ThrowUnsupportedArgument(BlockMongo, name)
		}
	}
	return nil
}

func populateMongoBlocks(mongo *Mongo, blocks hcl.Blocks) errors.Error {
	blocksByType := blocks.ByType()
	for blockType, blockList := range blocksByType {
		switch {
		case blockType == decoder.BlockQuery:
			if len(blockList) > 1 {
				return errors.IncorrectUsage("only 1 block %s is permitted in block %s", decoder.BlockQuery, BlockMongo)
			}
			query, err := decoder.DecodeQuery(blockList[0])
			if err != nil {
				return err
			}
			if mongo.operation == executor.OpUpdateMany || mongo.operation == executor.OpUpdateOne {
				if !query.HasSet() {
					return errors.IncorrectUsage("missing required block set in query for operation %s", mongo.operation)
				}
			}
			mongo.SetQuery(query)
		case blockType == decoder.BlockConnection:
			if len(blockList) > 1 {
				return errors.IncorrectUsage("only 1 block %s is permitted in block %s", decoder.BlockConnection, BlockMongo)
			}
			conn, err := decoder.DecodeConnection(blockList[0])
			if err != nil {
				return err
			}
			mongo.conn = conn
		case blockType == decoder.BlockResponse:
			if len(blockList) > 1 {
				return errors.IncorrectUsage("only 1 block %s is permitted in block %s", decoder.BlockResponse, BlockMongo)
			}
			properties, err := decoder.NewBlockProperties(blockList[0])
			if err != nil {
				return err
			}
			mongo.response = &decoder.Response{
				BlockProperties: properties,
			}
		default:
			return errors.ThrowUnsupportedBlock(BlockMongo, blockType)
		}
	}
	return nil
}
