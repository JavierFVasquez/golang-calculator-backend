package services

import (
	"context"
	"errors"
	"fmt"

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
