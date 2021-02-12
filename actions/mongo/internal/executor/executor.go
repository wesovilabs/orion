package executor

import (
	"context"

	log "github.com/sirupsen/logrus"
	"github.com/wesovilabs-tools/orion/internal/errors"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func New(ctx context.Context, operation string, database, collection string, opts *options.ClientOptions) (Executor, errors.Error) {

	return &executor{
		operation:  operation,
		ctx:        ctx,
		opts:       opts,
		dbName:     database,
		collection: collection,
	}, nil
}

type Executor interface {
	Run() (*MongoResponse, errors.Error)
	WithFilter(filter map[string]interface{}) Executor
	WithSet(data map[string]interface{}) Executor
	WithDocuments(documents []map[string]interface{}) Executor
	WithFindList(limit int) Executor
}

type executor struct {
	operation  string
	ctx        context.Context
	opts       *options.ClientOptions
	dbName     string
	collection string
	filter     map[string]interface{}
	set        map[string]interface{}
	documents  []map[string]interface{}
	findList   *findList
}

type findList struct {
	limit int
}

func (exec *executor) WithFindList(limit int) Executor {
	exec.findList = &findList{limit: limit}
	return exec
}
func (exec *executor) WithFilter(filter map[string]interface{}) Executor {
	exec.filter = filter
	return exec
}

func (exec *executor) WithSet(set map[string]interface{}) Executor {
	exec.set = set
	return exec
}

func (exec *executor) WithDocuments(documents []map[string]interface{}) Executor {
	exec.documents = documents
	return exec
}

type MongoResponse struct {
	Type     string
	Element  map[string]interface{}
	Elements []map[string]interface{}
}

func (exec *executor) Run() (*MongoResponse, errors.Error) {
	ctx := exec.ctx
	client, mngErr := mongo.Connect(ctx, exec.opts)
	if mngErr != nil {
		return nil, errors.Unexpected(mngErr.Error())
	}
	defer client.Disconnect(ctx)

	if mngErr := client.Ping(ctx, nil); mngErr != nil {
		return nil, errors.Unexpected(mngErr.Error())
	}

	db := client.Database(exec.dbName)
	log.Debugf("connected to database %s", exec.dbName)
	switch op := exec.operation; op {
	case OpFind:
		elements, err := exec.find(db)
		if err != nil {
			return nil, err
		}
		return &MongoResponse{
			Type:     "list",
			Elements: elements,
		}, nil
	case OpFindOne:
		element, err := exec.findOne(db)
		if err != nil {
			return nil, err
		}
		return &MongoResponse{
			Type:    "single",
			Element: element,
		}, nil
	case OpUpdateOne:
		return nil, exec.updateOne(db)
	case OpUpdateMany:
		return nil, exec.updateMany(db)
	case OpDeleteOne:
		return nil, exec.deleteOne(db)
	case OpDeleteMany:
		return nil, exec.deleteMany(db)
	case OpInsertOne:
		return nil, exec.insertOne(db)
	case OpInsertMany:
		return nil, exec.insertMany(db)
	case OpDrop:
		return nil, exec.drop(db)
	case OpCreate:
		return nil, exec.createCollection(db)
	default:
		return nil, errors.IncorrectUsage(
			"unsupported operation '%s'.", op)
	}
	return nil, nil
}
