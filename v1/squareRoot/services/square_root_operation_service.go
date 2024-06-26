package services

import (
	"context"
	"errors"
	"fmt"
	"math"
	"strconv"

	"github.com/JavierFVasquez/golang-calculator-backend/libs/models"
	"github.com/JavierFVasquez/golang-calculator-backend/libs/repositories"

	"github.com/rs/zerolog"
)

type SquareRootService struct {
	recordRepo    repositories.RecordRepositoryIF
	operationRepo repositories.OperationRepositoryIF
	logger        *zerolog.Logger
}

func NewSquareRootService(
	recordRepo repositories.RecordRepositoryIF,
	operationRepo repositories.OperationRepositoryIF,
	log *zerolog.Logger,
) *SquareRootService {
	l := log.With().Str("resource", "create_square_root_service").Logger()

	return &SquareRootService{
		recordRepo:    recordRepo,
		operationRepo: operationRepo,
		logger:        &l,
	}
}

func (service *SquareRootService) SquareRoot(
	ctx context.Context,
	operation models.Operation,
) (*models.Record, *error) {

	result := math.Sqrt(operation.Operand1)

	operationFromDB, err := service.operationRepo.GetOperation(ctx, models.SQUARE_ROOT)
	if err != nil {
		fmt.Println("Get operation record error")
		return nil, &err
	}
	operation.Result = &result
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
