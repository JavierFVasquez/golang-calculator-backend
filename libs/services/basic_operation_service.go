package services

import (
	"context"
	"errors"
	"fmt"
	"strconv"

	"github.com/JavierFVasquez/truenorth-calculator-backend/libs/models"
	"github.com/JavierFVasquez/truenorth-calculator-backend/libs/repositories"

	"github.com/rs/zerolog"
)

type BasicOperationService struct {
	recordRepo    repositories.RecordRepositoryIF
	operationRepo repositories.OperationRepositoryIF
	logger        *zerolog.Logger
}

func NewBasicOperationService(
	recordRepo repositories.RecordRepositoryIF,
	operationRepo repositories.OperationRepositoryIF,
	log *zerolog.Logger,
) *BasicOperationService {
	l := log.With().Str("resource", "create_basic_operation_service").Logger()

	return &BasicOperationService{
		recordRepo:    recordRepo,
		operationRepo: operationRepo,
		logger:        &l,
	}
}

func (service *BasicOperationService) BasicOperation(
	ctx context.Context,
	operation models.Operation,
) (*models.Record, *error) {

	var result float64
	switch operation.Operation {
	case models.ADDITION:
		result = operation.Operand1 + operation.Operand2
	case models.SUBSTRACTION:
		result = operation.Operand1 - operation.Operand2
	case models.MULTIPLICATION:
		result = operation.Operand1 * operation.Operand2
	case models.DIVISION:
		result = operation.Operand1 / operation.Operand2
	default:
		result = 0
	}

	operationFromDB, err := service.operationRepo.GetOperation(ctx, operation.Operation)
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
