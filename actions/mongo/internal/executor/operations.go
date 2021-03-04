package executor

import (
	log "github.com/sirupsen/logrus"
	"github.com/wesovilabs/orion/internal/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	OpFindOne    = "findOne"
	OpFind       = "find"
	OpUpdateOne  = "updateOne"
	OpUpdateMany = "updateMany"
	OpDeleteOne  = "deleteOne"
	OpDeleteMany = "deleteMany"
	OpInsertOne  = "insertOne"
	OpInsertMany = "insertMany"
	OpCreate     = "create"
	OpDrop       = "drop"
	OpCount      = "count"
	attrSet      = "$set"
)

func (exec *executor) find(db *mongo.Database) ([]map[string]interface{}, errors.Error) {
	ctx := exec.ctx
	collection := db.Collection(exec.collection)
	log.Debug("mongo find operation")
	opts := &options.FindOptions{}
	if exec.findList != nil {
		limit64 := int64(exec.findList.limit)
		opts.Limit = &limit64
	}
	cursor, err := collection.Find(ctx, exec.filter, opts)
	if err != nil {
		return nil, errors.Unexpected(err.Error())
	}
	if cursor.Err() != nil {
		return nil, errors.Unexpected(cursor.Err().Error())
	}
	output := make([]map[string]interface{}, 0)
	if err := cursor.All(ctx, &output); err != nil {
		return nil, errors.Unexpected(err.Error())
	}
	return output, nil
}

func (exec *executor) findOne(db *mongo.Database) (map[string]interface{}, errors.Error) {
	ctx := exec.ctx
	collection := db.Collection(exec.collection)
	log.Debug("mongo findOne operation")
	result := collection.FindOne(ctx, exec.filter)
	if result.Err() != nil {
		return nil, errors.Unexpected(result.Err().Error())
	}
	out := make(map[string]interface{})
	if err := result.Decode(&out); err != nil {
		return nil, errors.Unexpected(err.Error())
	}
	return out, nil
}

func (exec *executor) updateMany(db *mongo.Database) errors.Error {
	ctx := exec.ctx
	collection := db.Collection(exec.collection)
	log.Debug("mongo updateMany operation")
	update := bson.D{
		bson.E{Key: attrSet, Value: exec.set},
	}
	result, err := collection.UpdateMany(ctx, exec.filter, update)
	if err != nil {
		return errors.Unexpected(err.Error())
	}
	log.Debugf("records updateMany: %d", result.UpsertedCount)
	return nil
}

func (exec *executor) deleteMany(db *mongo.Database) errors.Error {
	ctx := exec.ctx
	collection := db.Collection(exec.collection)
	log.Debug("mongo deleteMany operation")
	_, err := collection.DeleteMany(ctx, exec.filter)
	if err != nil {
		return errors.Unexpected(err.Error())
	}
	return nil
}

func (exec *executor) insertMany(db *mongo.Database) errors.Error {
	ctx := exec.ctx
	collection := db.Collection(exec.collection)
	log.Debug("mongo insertMany operation")
	documents := make([]interface{}, len(exec.documents))
	for index := range exec.documents {
		documents[index] = exec.documents[index]
	}
	_, err := collection.InsertMany(ctx, documents)
	if err != nil {
		return errors.Unexpected(err.Error())
	}
	return nil
}

func (exec *executor) updateOne(db *mongo.Database) errors.Error {
	ctx := exec.ctx
	collection := db.Collection(exec.collection)
	log.Debug("mongo updateOne operation")
	updateMany := bson.D{
		bson.E{Key: attrSet, Value: exec.set},
	}
	result, err := collection.UpdateOne(ctx, exec.filter, updateMany)
	if err != nil {
		return errors.Unexpected(err.Error())
	}
	log.Debugf("records updateMany: %d", result.UpsertedCount)
	return nil
}

func (exec *executor) deleteOne(db *mongo.Database) errors.Error {
	ctx := exec.ctx
	collection := db.Collection(exec.collection)
	log.Debug("mongo deleteOne operation")
	_, err := collection.DeleteOne(ctx, exec.filter)
	if err != nil {
		return errors.Unexpected(err.Error())
	}
	return nil
}

func (exec *executor) insertOne(db *mongo.Database) errors.Error {
	ctx := exec.ctx
	collection := db.Collection(exec.collection)
	log.Debug("mongo insertOne operation")
	log.Infof("Document %v", exec.documents[0])
	_, err := collection.InsertOne(ctx, exec.documents[0])
	if err != nil {
		return errors.Unexpected(err.Error())
	}
	return nil
}

func (exec *executor) drop(db *mongo.Database) errors.Error {
	ctx := exec.ctx
	collection := db.Collection(exec.collection)
	log.Debug("mongo drop operation")
	err := collection.Drop(ctx)
	if err != nil {
		return errors.Unexpected(err.Error())
	}
	return nil
}

func (exec *executor) createCollection(db *mongo.Database) errors.Error {
	ctx := exec.ctx
	log.Debug("mongo createCollection operation")
	err := db.CreateCollection(ctx, exec.collection)
	if err != nil {
		return errors.Unexpected(err.Error())
	}
	return nil
}
