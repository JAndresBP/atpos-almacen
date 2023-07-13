package repositorios

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDbCtx struct {
	Ctx    context.Context
	Client *mongo.Client
}

var MongoDbCtxInstance *MongoDbCtx

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

func GetMongoDbCtx(connectionString string) *MongoDbCtx {
	if MongoDbCtxInstance == nil {

		ctx := context.TODO()
		client, err := mongo.Connect(ctx, options.Client().ApplyURI(connectionString))
		failOnError(err, "Failed to connect to Mongo")
		MongoDbCtxInstance = &MongoDbCtx{
			Ctx:    ctx,
			Client: client,
		}
	}
	return MongoDbCtxInstance
}

func (dbctx *MongoDbCtx) Close() {
	dbctx.Client.Disconnect(dbctx.Ctx)
}
