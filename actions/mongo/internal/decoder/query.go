package decoder

import (
	"github.com/hashicorp/hcl/v2"
	"github.com/wesovilabs/orion/helper"
	"github.com/wesovilabs/orion/internal/errors"
)

const (
	defDatabase   = "test"
	defCollection = "test"
	defLimit      = 0
)

type Filter struct {
	*BlockProperties
}

type Set struct {
	*BlockProperties
}

type Document struct {
	*BlockProperties
}

type Query struct {
	database   hcl.Expression
	collection hcl.Expression
	limit      hcl.Expression
	filter     *Filter
	set        *Set
	documents  []*Document
}

func (q *Query) HasSet() bool {
	return q.set != nil && len(q.set.values) > 0
}

func (q *Query) SetDatabase(expr hcl.Expression) {
	q.database = expr
}

func (q *Query) SetCollection(expr hcl.Expression) {
	q.collection = expr
}

func (q *Query) SetLimit(expr hcl.Expression) {
	q.limit = expr
}

func (q *Query) Database(ctx *hcl.EvalContext) (string, errors.Error) {
	return helper.GetExpressionValueAsString(ctx, q.database, defDatabase)
}

func (q *Query) Collection(ctx *hcl.EvalContext) (string, errors.Error) {
	return helper.GetExpressionValueAsString(ctx, q.collection, defCollection)
}

func (q *Query) Limit(ctx *hcl.EvalContext) (int, errors.Error) {
	return helper.GetExpressionValueAsInt(ctx, q.limit, defLimit)
}

func (q *Query) Filter(ctx *hcl.EvalContext) (map[string]interface{}, errors.Error) {
	if q.filter == nil {
		return make(map[string]interface{}), nil
	}
	return q.filter.Values(ctx)
}

func (q *Query) Set(ctx *hcl.EvalContext) (map[string]interface{}, errors.Error) {
	if q.set == nil {
		return make(map[string]interface{}), nil
	}
	return q.set.Values(ctx)
}

func (q *Query) Documents(ctx *hcl.EvalContext) ([]map[string]interface{}, errors.Error) {
	documents := make([]map[string]interface{}, len(q.documents))
	for index := range q.documents {
		document, err := q.documents[index].Values(ctx)
		if err != nil {
			return nil, err
		}
		documents[index] = document
	}
	return documents, nil
}
