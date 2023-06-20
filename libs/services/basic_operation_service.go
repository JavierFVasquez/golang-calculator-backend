package services

import (
	"context"
	"errors"
	"fmt"

	"github.com/JavierFVasquez/truenorth-calculator-backend/libs/models"
	"github.com/JavierFVasquez/truenorth-calculator-backend/libs/repositories"

	"github.com/rs/zerolog"
)
type key string

type BasicOperationService struct {
	recordRepo        repositories.RecordRepositoryIF
	operationRepo        repositories.OperationRepositoryIF
	logger          *zerolog.Logger
}

func NewBasicOperationService(
	recordRepo repositories.RecordRepositoryIF,
	operationRepo repositories.OperationRepositoryIF,
	log *zerolog.Logger,
) *BasicOperationService {
	l := log.With().Str("resource", "create_basic_operation_service").Logger()

	return &BasicOperationService{
		recordRepo:        recordRepo,
		operationRepo:        operationRepo,
		logger:          &l,
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

	user:=ctx.Value(key("user"))
	fmt.Println(user)
	if user != nil {
		if userValue, ok := user.(models.User); ok {
			userIdValue:= userValue.ID
			lastUserRecord, err := service.recordRepo.GetRecordsByUserId(ctx, userIdValue, &models.PaginationParams{Page: 1, PageSize: 1, SortBy: "createdAt"})
			if err != nil {
				fmt.Println("Get user last record error")
				err := errors.New("get user last record error")
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
					fmt.Println("Create record error")
					err := errors.New("create record error")
					return nil, &err
				}
				
				service.logger.Info().Interface("record", newRecord).Msg("Record saved")
				
				return newRecord, nil
				}else{
					err := errors.New("insufficent balance")
					return nil, &err
				}
		}
	}
	err = errors.New("no user")
	return nil, &err
}

