package decoder

import (
	"github.com/hashicorp/hcl/v2"

	"github.com/wesovilabs-tools/orion/internal/errors"
)

const (
	BlockHTTP        = "http"
	BlockHeaders     = "headers"
	BlockHeader      = "header"
	AttBaseURL       = "baseUrl"
	AttTimeout       = "timeout"
	LabelName        = "name"
	LabelType        = "type"
	AttributeValue   = "value"
	BlockAuthBasic   = "auth-basic"
	BlockAuthToken   = "auth-token"
	BlockCookie      = "cookie"
	AttributeName    = "name"
	AttributeBasURL  = "baseUrl"
	AttributePath    = "path"
	AttributeDomain  = "domain"
	BlockConnection  = "connection"
	AttributeTimeout = "timeout"
	AttributeProxy   = "Proxy"
	BlockPayload     = "payload"
	AttributeData    = "data"
	BlockQueryParams = "queryParams"
	BlockQueryParam  = "queryParam"
)

var schemaRequest = &hcl.BodySchema{
	Attributes: []hcl.AttributeSchema{
		{Name: AttributeBasURL, Required: true},
		{Name: AttributePath, Required: true},
	},
	Blocks: []hcl.BlockHeaderSchema{
		{Type: BlockHeaders},
		{Type: BlockHeader, LabelNames: []string{LabelName}},
		{Type: BlockQueryParams},
		{Type: BlockQueryParam, LabelNames: []string{LabelName}},
		{Type: BlockConnection},
		{Type: BlockCookie},
		{Type: BlockAuthBasic},
		{Type: BlockAuthToken},
		{Type: BlockPayload, LabelNames: []string{LabelType}},
	},
}

var schemaConnection = &hcl.BodySchema{
	Attributes: []hcl.AttributeSchema{
		{Name: AttributeTimeout, Required: false},
		{Name: AttributeProxy, Required: false},
	},
}

var schemaCookie = &hcl.BodySchema{
	Attributes: []hcl.AttributeSchema{
		{Name: AttributeName, Required: true},
		{Name: AttributeValue, Required: true},
		{Name: AttributeDomain, Required: false},
		{Name: AttributePath, Required: false},
	},
}

var schemaPayload = &hcl.BodySchema{
	Attributes: []hcl.AttributeSchema{
		{Name: AttributeData, Required: true},
	},
}

func DecodeRequest(block *hcl.Block) (*Request, errors.Error) {
	req := new(Request)
	bodyContent, d := block.Body.Content(schemaRequest)
	if err := errors.EvalDiagnostics(d); err != nil {
		return nil, err
	}

	if err := populateRequestAttributes(req, bodyContent.Attributes); err != nil {
		return nil, err
	}
	if err := populateRequestBlocks(req, bodyContent.Blocks); err != nil {
		return nil, err
	}
	return req, nil
}

func populateRequestAttributes(req *Request, attributes hcl.Attributes) errors.Error {
	for name := range attributes {
		attribute := attributes[name]
		switch name {
		case AttributePath:
			req.SetPath(attribute.Expr)
		case AttBaseURL:
			req.SetBaseURL(attribute.Expr)
		default:
			return errors.ThrowUnsupportedArgument(BlockHTTP, name)
		}
	}
	return nil
}

func populateRequestBlocks(req *Request, blocks hcl.Blocks) errors.Error {
	blocksByType := blocks.ByType()
	var err errors.Error
	for blockType, blockList := range blocksByType {
		switch blockType {
		case BlockConnection:
			if len(blockList) > 1 {
				// TODO Throw error
			}
			req.connection, err = decodeConnection(blockList[0])
			if err != nil {
				return err
			}
		case BlockCookie:
			req.cookies = make([]*Cookie, len(blockList))
			for index := range blockList {
				req.cookies[index], err = decodeCookie(blockList[index])
			}
		case BlockPayload:
			if len(blockList) > 1 {
				// TODO Throw error
			}
			for index := range blockList {
				req.payload, err = decodePayload(blockList[index])
			}
		case BlockHeaders:
			for index := range blockList {
				headers, err := decodeHeaders(blockList[index])
				if err != nil {
					return err
				}
				req.AddHeaders(headers)
			}
		case BlockQueryParams:
			for index := range blockList {
				queryParams, err := decodeQueryParams(blockList[index])
				if err != nil {
					return err
				}
				req.AddQueryParams(queryParams)
			}

		default:
			return errors.ThrowUnsupportedBlock(BlockHTTP, blockType)
		}
	}
	return err
}

func decodeConnection(block *hcl.Block) (*Connection, errors.Error) {
	bodyContent, d := block.Body.Content(schemaConnection)
	if err := errors.EvalDiagnostics(d); err != nil {
		return nil, err
	}
	connection := new(Connection)
	for name := range bodyContent.Attributes {
		attribute := bodyContent.Attributes[name]
		switch attribute.Name {
		case AttTimeout:
			connection.SetTimeout(attribute.Expr)
		case AttributeProxy:
			connection.SetProxy(attribute.Expr)
		default:
			return nil, errors.ThrowUnsupportedArgument(BlockConnection, name)
		}
	}
	return connection, nil
}

func decodeCookie(block *hcl.Block) (*Cookie, errors.Error) {
	cookie := new(Cookie)
	bodyContent, d := block.Body.Content(schemaCookie)
	if err := errors.EvalDiagnostics(d); err != nil {
		return nil, err
	}
	for name, attribute := range bodyContent.Attributes {
		switch name {
		case AttributeValue:
			cookie.SetValue(attribute.Expr)
		case AttributeName:
			cookie.SetName(attribute.Expr)
		case AttributePath:
			cookie.SetPath(attribute.Expr)
		case AttributeDomain:
			cookie.SetDomain(attribute.Expr)
		default:
			return nil, errors.ThrowUnsupportedArgument(BlockCookie, name)
		}
	}
	return cookie, nil
}

func decodeHeaders(block *hcl.Block) (Headers, errors.Error) {
	attributes, d := block.Body.JustAttributes()
	if err := errors.EvalDiagnostics(d); err != nil {
		return nil, err
	}
	headers := make(map[string]hcl.Expression)
	for name := range attributes {
		headers[name] = attributes[name].Expr
	}
	return headers, nil
}

func decodeQueryParams(block *hcl.Block) (QueryParams, errors.Error) {
	attributes, d := block.Body.JustAttributes()
	if err := errors.EvalDiagnostics(d); err != nil {
		return nil, err
	}
	queryParams := make(map[string]hcl.Expression)
	for name := range attributes {
		queryParams[name] = attributes[name].Expr
	}
	return queryParams, nil
}

func decodePayload(block *hcl.Block) (*Payload, errors.Error) {
	bodyContent, d := block.Body.Content(schemaPayload)
	if err := errors.EvalDiagnostics(d); err != nil {
		return nil, err
	}
	payload := new(Payload)
	payload.SetType(block.Labels[0])
	for name := range bodyContent.Attributes {
		switch name {
		case AttributeData:
			payload.SetData(bodyContent.Attributes[name].Expr)
		default:
			return nil, errors.ThrowUnsupportedArgument(BlockPayload, name)
		}
	}
	return payload, nil
}

func DecodeResponse(block *hcl.Block) (*Response, errors.Error) {
	res := new(Response)
	attributes, d := block.Body.JustAttributes()
	if err := errors.EvalDiagnostics(d); err != nil {
		return nil, err
	}
	res.SetValues(attributes)
	return res, nil
}
