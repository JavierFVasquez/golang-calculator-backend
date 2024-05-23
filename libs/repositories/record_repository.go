package repositories

import (
	"context"
	"errors"
	"time"

	"github.com/JavierFVasquez/golang-calculator-backend/libs/clients"
	"github.com/JavierFVasquez/golang-calculator-backend/libs/models"
	"github.com/JavierFVasquez/golang-calculator-backend/libs/repositories/aggregations"
	aggregationPipes "github.com/JavierFVasquez/golang-calculator-backend/libs/repositories/aggregations/pipes"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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

func (repo *RecordRepository) DeleteRecord(ctx context.Context, recordId *string) (*models.Record, error) {
	if user := ctx.Value("user"); user != nil {
		if userValue, ok := user.(models.User); ok {
			now := time.Now()
			filter := bson.D{
				{Key: "_id", Value: recordId},
				{Key: "user.id", Value: userValue.ID},
			}

			update := bson.D{
				{Key: "$set", Value: bson.D{{Key: "deletedAt", Value: now}}},
			}
			result := repo.collection.FindOneAndUpdate(ctx, filter, update)
			if err := result.Err(); err != nil {
				return nil, err
			}

			record := &models.Record{}
			err := result.Decode(record)
			if err != nil {
				return nil, err
			}

			return record, nil
		}
	}
	err := errors.New("no user found")
	return nil, err
}

func (repo *RecordRepository) GetRecordsByUserId(ctx context.Context, id string, options *models.PaginationParams) (*models.PaginatedResponse[models.Record], error) {
	aggOptions := aggregationPipes.PaginationToAggregation(options)
	aggregation := aggregations.BuildGetRecordsByUserIdAggregation(id, &aggOptions)
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

func (repo *RecordRepository) GetUserMostRecentRecord(ctx context.Context) (*models.Record, error) {
	if user := ctx.Value("user"); user != nil {
		if userValue, ok := user.(models.User); ok {
			mostRecentRecord := models.Record{}
			filter := bson.M{
				"user.id": userValue.ID,
			}
			options := options.FindOne().SetSort(bson.D{{Key: "createdAt", Value: -1}})

			err := repo.collection.FindOne(context.TODO(), filter, options).Decode(&mostRecentRecord)
			if err != nil {
				return nil, err
			}
			return &mostRecentRecord, nil
		}
	}
	err := errors.New("no user found")
	return nil, err
}

func closeCursor(ctx context.Context, cursor *mongo.Cursor) {
	err := cursor.Close(ctx)
	if err != nil {
		log.Err(err).Msg("Unable to close mongo cursor")
	}
}
