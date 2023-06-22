package services

import (
	"context"
	"errors"
	"fmt"

	"github.com/JavierFVasquez/truenorth-calculator-backend/libs/models"
	"github.com/JavierFVasquez/truenorth-calculator-backend/libs/repositories"

	"github.com/rs/zerolog"
)

type RecordService struct {
	recordRepo repositories.RecordRepositoryIF
	logger     *zerolog.Logger
}

func NewRecordService(
	recordRepo repositories.RecordRepositoryIF,
	log *zerolog.Logger,
) *RecordService {
	l := log.With().Str("resource", "create_random_string_service").Logger()

	return &RecordService{
		recordRepo: recordRepo,
		logger:     &l,
	}
}

func (service *RecordService) GetRecordList(
	ctx context.Context,
	options *models.PaginationParams,
) (*models.PaginatedResponse[models.Record], *error) {

	if user := ctx.Value("user"); user != nil {
		if userValue, ok := user.(models.User); ok {
			userIdValue := userValue.ID

			result, randomErr := service.recordRepo.GetRecordsByUserId(ctx, userIdValue, options)
			if randomErr != nil {
				fmt.Println("Get operation record error")
				return nil, &randomErr
			}

			return result, nil

		}
	}
	err := errors.New("NO_USER_ERROR")
	return nil, &err
}
