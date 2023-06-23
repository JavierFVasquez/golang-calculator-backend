package aggregations

import (
	aggregationModels "github.com/JavierFVasquez/truenorth-calculator-backend/libs/repositories/aggregations/models"
	aggregationPipes "github.com/JavierFVasquez/truenorth-calculator-backend/libs/repositories/aggregations/pipes"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func BuildGetRecordsByUserIdAggregation(id string, options *aggregationModels.AggregationPaginationOptions) mongo.Pipeline {
	matchPipe := bson.D{
		{Key: "$match",
			Value: bson.D{
				{Key: "user.id", Value: id},
				{Key: "$or", Value: bson.A{
					bson.D{{Key: "deletedAt", Value: nil}},
					bson.D{{Key: "deletedAt", Value: bson.D{{Key: "$exists", Value: false}}}},
				}},
			},
		},
	}

	pagingPipe := aggregationPipes.PaginationStep(options)

	pipeline := mongo.Pipeline{matchPipe}

	if !options.GetAll {
		pipeline = append(pipeline, pagingPipe...)
	}

	return pipeline
}
