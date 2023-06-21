package aggregationPipes

import (
	"github.com/JavierFVasquez/truenorth-calculator-backend/libs/models"
	"github.com/JavierFVasquez/truenorth-calculator-backend/libs/pipes"
	aggregationModels "github.com/JavierFVasquez/truenorth-calculator-backend/libs/repositories/aggregations/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func PaginationToAggregation(options *models.PaginationParams) aggregationModels.AggregationPaginationOptions {
	pipes.InitPagination(options)

	aggregationOptions := aggregationModels.AggregationPaginationOptions{
		Skip:     (options.Page - 1) * options.PageSize,
		Limit:    options.PageSize,
		Sort:     options.SortBy,
		GetAll:   options.GetAll,
		Page:     options.Page,
		PageSize: options.PageSize,
		Search:   options.Search,
	}

	return aggregationOptions
}

func PaginationStep(options *aggregationModels.AggregationPaginationOptions) mongo.Pipeline {
	// Pagination Step
	paginationStep := bson.A{
		bson.D{{Key: "$skip", Value: options.Skip}},
	}

	if !options.GetAll {
		paginationStep = append(paginationStep, bson.D{{Key: "$limit", Value: options.Limit}})
	}

	if options.Sort != "" && !options.GetAll {
		paginationStep = append(paginationStep, bson.D{{Key: "$sort", Value: bson.D{{Key: options.Sort, Value: -1}}}})
	}

	// Metadata Step
	metadataStep := bson.A{
		bson.D{{Key: "$count", Value: "length"}},
		bson.D{
			{Key: "$addFields",
				Value: bson.D{
					{Key: "page", Value: options.Page},
					{Key: "pageSize", Value: options.PageSize},
					{Key: "hasNext",
						Value: bson.D{
							{Key: "$cond",
								Value: bson.D{
									{Key: "if", Value: bson.D{
										{Key: "$lt", Value: bson.A{
											options.Page,
											bson.D{{Key: "$divide", Value: bson.A{"$length", options.PageSize}}},
										}},
									}},
									{Key: "then", Value: true},
									{Key: "else", Value: false},
								},
							},
						},
					},
				},
			},
		},
	}

	// build facet
	facetStep := bson.D{{Key: "$facet", Value: bson.D{
		{Key: "data", Value: paginationStep},
		{Key: "metadata", Value: metadataStep},
	}}}

	// project step
	projectStep := bson.D{
		{Key: "$project",
			Value: bson.D{
				{Key: "_id", Value: 0},
				{Key: "data", Value: "$data"},
				{Key: "metadata",
					Value: bson.D{{Key: "$arrayElemAt", Value: bson.A{"$metadata", 0}}},
				},
			},
		},
	}

	return mongo.Pipeline{facetStep, projectStep}
}
