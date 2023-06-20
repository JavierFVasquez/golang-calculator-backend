package repositories

import (
	"context"
	"errors"

	"github.com/JavierFVasquez/truenorth-calculator-backend/libs/clients"
	"github.com/JavierFVasquez/truenorth-calculator-backend/libs/models"

	"github.com/rs/zerolog"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type OperationRepository struct {
	collection *mongo.Collection
	logger     *zerolog.Logger
}

func NewOperationRepository(mongo *clients.MongoClient, tableName string, log *zerolog.Logger) *OperationRepository {
	l := log.With().Str("resource", "operation_repository").Logger()

	return &OperationRepository{
		collection: mongo.DB.Collection(tableName),
		logger:     &l,
	}
}


func (repo *OperationRepository) GetOperation(ctx context.Context, id models.Operator) (*models.Operation, error) {

	operation := &models.Operation{}
	result := repo.collection.FindOne(ctx, bson.D{{Key: "_id", Value: id}})
	err := result.Decode(operation)

	if errors.Is(err, mongo.ErrNoDocuments) {
		return nil, nil
	}

	if err != nil {
		repo.logger.Err(err).Msg("Get operation by id error")
		return nil, err
	}

	return operation, nil
}
