package pipes

import "github.com/JavierFVasquez/golang-calculator-backend/libs/models"

func InitPagination(params *models.PaginationParams) {
	if params.Page == 0 {
		params.Page = 1
	}
	if params.PageSize == 0 {
		params.PageSize = 10
	}
	if params.SortBy == "" {
		params.SortBy = "createdAt"
	}
}

func DataToPaginatedResponse[T any](data []T, options *models.PaginationParams) models.PaginatedResponse[T] {
	if options.GetAll {
		return models.PaginatedResponse[T]{
			Data: data,
			Metadata: models.ResponseMetadata{
				PageSize: len(data),
				Length:   len(data),
				Page:     1,
			},
		}
	}

	finalResponse := models.PaginatedResponse[T]{
		Data: data,
		Metadata: models.ResponseMetadata{
			Length:   len(data),
			Page:     options.Page,
			PageSize: options.PageSize,
		},
	}

	if len(data) < 1 {
		finalResponse.Data = []T{}
	}

	if len(data) > options.PageSize {
		finalResponse.Metadata.HasNext = true
		finalResponse.Metadata.Length = options.PageSize
		finalResponse.Data = data[0:options.PageSize]
	}

	return finalResponse
}
