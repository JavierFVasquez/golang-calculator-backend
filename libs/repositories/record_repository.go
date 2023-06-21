package repositories

import (
	"context"
	"errors"

	"github.com/JavierFVasquez/truenorth-calculator-backend/libs/clients"
	"github.com/JavierFVasquez/truenorth-calculator-backend/libs/models"
	"github.com/JavierFVasquez/truenorth-calculator-backend/libs/repositories/aggregations"
	aggregationPipes "github.com/JavierFVasquez/truenorth-calculator-backend/libs/repositories/aggregations/pipes"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/mongo"
)

type key string

type RecordRepository struct {
	collection *mongo.Collection
	logger     *zerolog.Logger
}

func NewRecordRepository(mongo *clients.MongoClient, tableName string, log *zerolog.Logger) *RecordRepository {
	l := log.With().Str("resource", "record_repository").Logger()

	return &RecordRepository{
		collection: mongo.DB.Collection(tableName),
		logger:     &l,
	}
}

func (repo *RecordRepository) CreateNewRecord(ctx context.Context, operation models.Operation, newUserBalance float64) (*models.Record, error) {
	if user := ctx.Value("user"); user != nil {
		if userValue, ok := user.(models.User); ok {
			record := models.NewRecord(userValue, operation, operation.Cost, newUserBalance)
			if _, err := repo.collection.InsertOne(ctx, record); err != nil {
				repo.logger.Err(err).Msg("Store record error")
				return nil, err
			}

			return record, nil
		}
	}
	err := errors.New("no user found")
	return nil, err
}

func (repo *RecordRepository) GetRecordsByUserId(ctx context.Context, ids string, options *models.PaginationParams) (*models.PaginatedResponse[models.Record], error) {
	aggOptions := aggregationPipes.PaginationToAggregation(options)
	aggregation := aggregations.BuildGetRecordsByUserIdAggregation(ids, &aggOptions)
	pagedRecords := []models.PaginatedResponse[models.Record]{}
	allRecords := []models.Record{}

	cursor, err := repo.collection.Aggregate(ctx, aggregation)
	if err != nil {
		repo.logger.Err(err).Msg("Error getting records by user ids")
		return nil, err
	}

	defer closeCursor(ctx, cursor)
	if options.GetAll {
		if err := cursor.All(ctx, &allRecords); err != nil {
			repo.logger.Err(err).Msg("Error decoding cursor allRecords")
			return nil, err
		}
		paginatedResponse := models.PaginatedResponse[models.Record]{
			Metadata: models.ResponseMetadata{
				Page:     1,
				PageSize: len(allRecords),
				Length:   len(allRecords),
				HasNext:  false,
			},
			Data: allRecords,
		}
		return &paginatedResponse, nil
	} else {
		if err := cursor.All(ctx, &pagedRecords); err != nil {
			repo.logger.Err(err).Msg("Error decoding cursor")
			return nil, err
		}
		return &pagedRecords[0], nil
	}
}

func closeCursor(ctx context.Context, cursor *mongo.Cursor) {
	err := cursor.Close(ctx)
	if err != nil {
		log.Err(err).Msg("Unable to close mongo cursor")
	}
}
