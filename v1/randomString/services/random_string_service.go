package services

import (
	"context"
	"errors"
	"fmt"
	"strconv"

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

	lastUserRecord, err := service.recordRepo.GetUserMostRecentRecord(ctx)
	if err != nil {
		err := errors.New("GET_USER_LAST_RECORD_ERROR")
		return nil, &err
	}

	userBalance := float64(lastUserRecord.UserBalance)
	newUserBalance := userBalance - operation.Cost
	if newUserBalance >= 0 {
		newRecord, err := service.recordRepo.CreateNewRecord(ctx, operation, newUserBalance)
		if err != nil {
			return nil, &err
		}

		service.logger.Info().Interface("record", newRecord).Msg("Record saved")

		return newRecord, nil
	} else {
		err := errors.New("INSUFFICENT_BALANCE_OPERATION_COST_" + strconv.FormatFloat(operation.Cost, 'f', -1, 64))
		return nil, &err
	}

}
