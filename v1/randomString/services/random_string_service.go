package services

import (
	"context"
	"errors"
	"fmt"

	"github.com/JavierFVasquez/truenorth-calculator-backend/libs/models"
	"github.com/JavierFVasquez/truenorth-calculator-backend/libs/repositories"
	random_string_repositories "github.com/JavierFVasquez/truenorth-calculator-backend/v1/randomString/repositories"

	"github.com/rs/zerolog"
)

type RandomStringService struct {
	recordRepo       repositories.RecordRepositoryIF
	operationRepo    repositories.OperationRepositoryIF
	randomStringRepo random_string_repositories.RandomStringRepositoryIF
	logger           *zerolog.Logger
}

func NewRandomStringService(
	recordRepo repositories.RecordRepositoryIF,
	operationRepo repositories.OperationRepositoryIF,
	randomStringRepo random_string_repositories.RandomStringRepositoryIF,
	log *zerolog.Logger,
) *RandomStringService {
	l := log.With().Str("resource", "create_random_string_service").Logger()

	return &RandomStringService{
		recordRepo:       recordRepo,
		operationRepo:    operationRepo,
		randomStringRepo: randomStringRepo,
		logger:           &l,
	}
}

func (service *RandomStringService) RandomString(
	ctx context.Context,
	operation models.Operation,
) (*models.Record, *error) {

	result, randomErr := service.randomStringRepo.GetRandomString(ctx)
	if randomErr != nil {
		fmt.Println("Get operation record error")
		return nil, &randomErr
	}

	operationFromDB, err := service.operationRepo.GetOperation(ctx, models.RANDOM_STRING)
	if err != nil {
		fmt.Println("Get operation record error")
		return nil, &err
	}

	defaultResponse := 0.0
	operation.Result = &defaultResponse
	operation.StringResult = result
	operation.Cost = operationFromDB.Cost

	if user := ctx.Value("user"); user != nil {
		if userValue, ok := user.(models.User); ok {
			userIdValue := userValue.ID
			lastUserRecord, err := service.recordRepo.GetRecordsByUserId(ctx, userIdValue, &models.PaginationParams{Page: 1, PageSize: 20, SortBy: "createdAt"})
			if err != nil {
				err := errors.New("GET_USER_LAST_RECORD_ERROR")
				return nil, &err
			}

			var userBalance float64
			if lastUserRecord.Data != nil {
				if ok && len(lastUserRecord.Data) > 0 {
					userBalance = float64(lastUserRecord.Data[0].UserBalance)
				}
			}
			newUserBalance := userBalance - operation.Cost
			if newUserBalance >= 0 {
				newRecord, err := service.recordRepo.CreateNewRecord(ctx, operation, newUserBalance)
				if err != nil {
					return nil, &err
				}

				service.logger.Info().Interface("record", newRecord).Msg("Record saved")

				return newRecord, nil
			} else {
				err := errors.New("INSUFFICENT_BALANCE")
				return nil, &err
			}
		}
	}
	err = errors.New("NO_USER_ERROR")
	return nil, &err
}
