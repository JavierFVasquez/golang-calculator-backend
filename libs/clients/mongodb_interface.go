package clients

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//go:generate mockery --name MongoDatabaseIF  --structname MockMongoDatabaseIF --filename mongodb_database_mock.go
type MongoDatabaseIF interface {
	Collection(name string, opts ...*options.CollectionOptions) *mongo.Collection
}

//go:generate mockery --name MongoCollectionIF --structname MockMongoCollectionIF --filename mongodb_collection_mock.go
type MongoCollectionIF interface {
	InsertOne(ctx context.Context, document interface{}, opts ...*options.InsertOneOptions) (*mongo.InsertOneResult, error)
}
