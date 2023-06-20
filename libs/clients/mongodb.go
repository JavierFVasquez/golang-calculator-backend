package clients

import (
	"context"

	"github.com/rs/zerolog"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoClient struct {
	DB     MongoDatabaseIF
	logger *zerolog.Logger
}

func NewMongoClient(ctx context.Context, uri string, dbName string, logger *zerolog.Logger) (*MongoClient, func()) {
	log := logger.With().Str("resource", "mongodb_client").Logger()

	// Use the SetServerAPIOptions() method to set the Stable API version to 1
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI)

	// Create a new client and connect to the server
	connection, err := mongo.Connect(ctx, opts)
	if err != nil {
		log.Panic().Err(err).Msg("Unable to establish mongodb connection")
	}

	mongoDb := connection.Database(dbName)
	log.Info().Msg("Successfully connected to MongoDB!")

	mongoClient := &MongoClient{
		DB:     mongoDb,
		logger: &log,
	}

	closeFunc := func() {
		if err = connection.Disconnect(ctx); err != nil {
			log.Panic().Err(err).Msg("Unable to close mongodb connection")
		}
	}

	return mongoClient, closeFunc
}
