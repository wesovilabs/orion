package mongo

import (
	"context"

	"github.com/hashicorp/hcl/v2"
	log "github.com/sirupsen/logrus"
	"github.com/wesovilabs/orion/actions/mongo/internal/helper"

	"github.com/wesovilabs/orion/actions"
	"github.com/wesovilabs/orion/actions/mongo/internal/decoder"
	"github.com/wesovilabs/orion/actions/mongo/internal/executor"
	orionContext "github.com/wesovilabs/orion/context"
	"github.com/wesovilabs/orion/internal/errors"
	"github.com/zclconf/go-cty/cty"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var supportedOperations = map[string]struct{}{
	executor.OpDrop:       {},
	executor.OpCreate:     {},
	executor.OpFindOne:    {},
	executor.OpFind:       {},
	executor.OpDeleteOne:  {},
	executor.OpDeleteMany: {},
	executor.OpUpdateOne:  {},
	executor.OpUpdateMany: {},
	executor.OpInsertOne:  {},
	executor.OpInsertMany: {},
}

type Mongo struct {
	*actions.Base
	operation string
	query     *decoder.Query
	conn      *decoder.Connection
	response  *decoder.Response
}

func (m *Mongo) SetQuery(q *decoder.Query) {
	m.query = q
}

func (m *Mongo) SetOperation(operation string) {
	m.operation = operation
}

func (m *Mongo) CreteContext(evalCtx *hcl.EvalContext) context.Context {
	if m.conn == nil {
		return context.Background()
	}
	timeout, err := m.conn.Timeout(evalCtx)
	if err != nil {
		log.Warn("it fails silently obtaining the timeout.")
	}
	if timeout != nil {
		// nolint
		ctx, _ := context.WithTimeout(context.Background(), *timeout)
		// cancelFn()
		return ctx
	}
	return context.Background()
}

func (m *Mongo) Filter(ctx orionContext.OrionContext) (map[string]interface{}, errors.Error) {
	filter, err := m.query.Filter(ctx.EvalContext())
	out := bson.M{}
	for k, v := range filter {
		if k == "_id" {
			id, _ := primitive.ObjectIDFromHex(v.(string))
			out[k] = id
			continue
		}
		out[k] = v
	}
	return out, err
}

func (m *Mongo) Operation() string {
	return m.operation
}

func (m *Mongo) Execute(ctx orionContext.OrionContext) errors.Error {
	mngCtx := m.CreteContext(ctx.EvalContext())
	clientOptions, err := m.conn.ClientOpts(ctx.EvalContext())
	if err != nil {
		return err
	}
	database, err := m.query.Database(ctx.EvalContext())
	if err != nil {
		return err
	}
	collection, err := m.query.Collection(ctx.EvalContext())
	if err != nil {
		return err
	}
	executor, err := executor.New(mngCtx, m.operation, database, collection, clientOptions)
	if err != nil {
		return err
	}
	filter, mngErr := m.query.Filter(ctx.EvalContext())
	if mngErr != nil {
		return errors.Unexpected(mngErr.Error())
	}
	set, mngErr := m.query.Set(ctx.EvalContext())
	if mngErr != nil {
		return errors.Unexpected(mngErr.Error())
	}
	documents, mngErr := m.query.Documents(ctx.EvalContext())
	if mngErr != nil {
		return errors.Unexpected(mngErr.Error())
	}
	executor = executor.WithSet(set).WithFilter(filter).WithDocuments(documents)
	limitVal, mngErr := m.query.Limit(ctx.EvalContext())
	if mngErr != nil {
		return errors.Unexpected(mngErr.Error())
	}
	if limitVal != 0 {
		executor.WithFindList(limitVal)
	}
	response, err := executor.Run()
	if err != nil {
		return err
	}
	if response != nil {
		switch response.Type {
		case "list":
			return m.processElementsResponse(ctx.EvalContext(), response.Elements)
		case "single":
			return m.processElementResponse(ctx.EvalContext(), response.Element)
		}
	}
	cleanVariables(ctx.EvalContext())
	return nil
}

func (m *Mongo) processElementResponse(ctx *hcl.EvalContext, v interface{}) errors.Error {
	if m.response == nil {
		return nil
	}
	setDocumentAsVar(ctx, v)
	defer cleanVariables(ctx)
	return m.response.Evaluate(ctx)
}

func (m *Mongo) processElementsResponse(ctx *hcl.EvalContext, v []map[string]interface{}) errors.Error {
	if m.response == nil {
		return nil
	}
	setDocumentsAsVar(ctx, v)
	defer cleanVariables(ctx)
	return m.response.Evaluate(ctx)
}

func setDocumentAsVar(ctx *hcl.EvalContext, document interface{}) {
	mongoVars := cty.ObjectVal(map[string]cty.Value{
		"document": helper.ToValue(document),
	})
	setMongoVars(ctx, mongoVars)
}

func setDocumentsAsVar(ctx *hcl.EvalContext, documents []map[string]interface{}) {
	mongoVars := cty.ObjectVal(map[string]cty.Value{
		"documents": helper.ToValue(documents),
	})
	setMongoVars(ctx, mongoVars)
}

func setMongoVars(ctx *hcl.EvalContext, mongoVars cty.Value) {
	if rootVars, ok := ctx.Variables["_"]; ok {
		rootValueMap := rootVars.AsValueMap()
		if rootValueMap != nil {
			rootValueMap[BlockMongo] = mongoVars
			ctx.Variables["_"] = cty.ObjectVal(rootValueMap)
			return
		}
	}
	ctx.Variables["_"] = cty.ObjectVal(map[string]cty.Value{
		BlockMongo: mongoVars,
	})
}

func cleanVariables(evalCtx *hcl.EvalContext) {
	variables := evalCtx.Variables
	if _, ok := variables["_"]; ok {
		rootVars := variables["_"].AsValueMap()
		delete(rootVars, BlockMongo)
		variables["_"] = cty.ObjectVal(rootVars)
		evalCtx.Variables = variables
	}
}
