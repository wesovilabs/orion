package decoder

import (
	"github.com/hashicorp/hcl/v2"
	"github.com/wesovilabs-tools/orion/internal/errors"
)

const (
	BlockConnection = "connection"
	BlockQuery      = "query"
	BlockResponse   = "response"
	argTimeout      = "timeout"
	argURI          = "uri"
	argDatabase     = "database"
	argCollection   = "collection"
	argLimit        = "limit"
	blockAuth       = "auth"
	labelMechanism  = "mechanism"
	argUsername     = "username"
	argPassword     = "password"
	blockSet        = "set"
	blockDocument   = "document"
	blockFilter     = "filter"
)

var schemaConnection = &hcl.BodySchema{
	Attributes: []hcl.AttributeSchema{
		{Name: argTimeout, Required: false},
		{Name: argURI, Required: false},
	},
	Blocks: []hcl.BlockHeaderSchema{
		{
			Type:       blockAuth,
			LabelNames: []string{labelMechanism},
		},
	},
}

var schemaAuth = &hcl.BodySchema{
	Attributes: []hcl.AttributeSchema{
		{Name: argUsername, Required: false},
		{Name: argPassword, Required: false},
	},
}

var schemaQuery = &hcl.BodySchema{
	Attributes: []hcl.AttributeSchema{
		{Name: argDatabase, Required: true},
		{Name: argCollection, Required: true},
		{Name: argLimit, Required: false},
	},
	Blocks: []hcl.BlockHeaderSchema{
		{Type: blockSet},
		{Type: blockFilter},
		{Type: blockDocument},
	},
}

func DecodeConnection(block *hcl.Block) (*Connection, errors.Error) {
	bodyContent, d := block.Body.Content(schemaConnection)
	if err := errors.EvalDiagnostics(d); err != nil {
		return nil, err
	}
	connection := new(Connection)
	for name := range bodyContent.Attributes {
		attribute := bodyContent.Attributes[name]
		switch attribute.Name {
		case argTimeout:
			connection.SetTimeout(attribute.Expr)
		case argURI:
			connection.SetURI(attribute.Expr)
		default:
			return nil, errors.ThrowUnsupportedArgument(BlockConnection, name)
		}
	}
	blocksByType := bodyContent.Blocks.ByType()
	for blockType, blocks := range blocksByType {
		if blockType != blockAuth {
			return nil, errors.ThrowUnsupportedBlock(BlockConnection, blockType)
		}
		if len(blocks) > 1 {
			return nil, errors.IncorrectUsage("only 1 auth can be defined for the connection")
		}
		auth, err := decodeAuth(blocks[0])
		if err != nil {
			return nil, err
		}
		connection.auth = auth
	}
	return connection, nil
}

func decodeAuth(block *hcl.Block) (*Auth, errors.Error) {
	bodyContent, d := block.Body.Content(schemaAuth)
	if err := errors.EvalDiagnostics(d); err != nil {
		return nil, err
	}
	authMechanism := block.Labels[0]
	if len(block.Labels) != 1 {
		return nil, errors.ThrowMissingRequiredLabel(blockAuth)
	}
	if _, ok := supportedAuthMechanisms[authMechanism]; !ok {
		return nil, errors.IncorrectUsage("mechanism '%s' is not valid. Supported authentication mechanism are: %v",
			authMechanism, supportedAuthMechanismsName)
	}
	auth := &Auth{
		mechanism: authMechanism,
	}
	for name := range bodyContent.Attributes {
		attribute := bodyContent.Attributes[name]
		switch attribute.Name {
		case argUsername:
			auth.SetUsername(attribute.Expr)
		case argPassword:
			auth.SetPassword(attribute.Expr)
		default:
			return nil, errors.ThrowUnsupportedArgument(blockAuth, name)
		}
	}
	return auth, nil
}

func DecodeQuery(block *hcl.Block) (*Query, errors.Error) {
	bodyContent, d := block.Body.Content(schemaQuery)
	if err := errors.EvalDiagnostics(d); err != nil {
		return nil, err
	}
	query := new(Query)
	for name := range bodyContent.Attributes {
		attribute := bodyContent.Attributes[name]
		switch attribute.Name {
		case argDatabase:
			query.SetDatabase(attribute.Expr)
		case argCollection:
			query.SetCollection(attribute.Expr)
		case argLimit:
			query.SetLimit(attribute.Expr)
		default:
			return nil, errors.ThrowUnsupportedArgument(BlockQuery, name)
		}
	}
	blocksByType := bodyContent.Blocks.ByType()
	for blockType, blocks := range blocksByType {

		switch blockType {
		case blockSet:
			if len(blocks) > 1 {
				return nil, errors.IncorrectUsage("only 1 block %s is permitted in block %s", blockSet, BlockQuery)
			}
			blockProps, err := NewBlockProperties(blocks[0])
			if err != nil {
				return nil, err
			}
			query.set = &Set{
				BlockProperties: blockProps,
			}
		case blockFilter:
			if len(blocks) > 1 {
				return nil, errors.IncorrectUsage("only 1 block %s is permitted in block %s", blockFilter, BlockQuery)
			}
			blockProps, err := NewBlockProperties(blocks[0])
			if err != nil {
				return nil, err
			}
			query.filter = &Filter{
				BlockProperties: blockProps,
			}
		case blockDocument:
			query.documents = make([]*Document, len(blocks))
			for index := range blocks {
				child := blocks[index]
				blockProps, err := NewBlockProperties(child)
				if err != nil {
					return nil, err
				}
				query.documents[index] = &Document{
					BlockProperties: blockProps,
				}
			}
		default:
			return nil, errors.ThrowUnsupportedBlock(BlockQuery, blockType)
		}

	}

	return query, nil
}
