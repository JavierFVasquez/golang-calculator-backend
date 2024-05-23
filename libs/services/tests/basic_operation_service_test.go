package services

import (
	"context"
	"testing"

	"github.com/JavierFVasquez/golang-calculator-backend/libs/models"
	"github.com/JavierFVasquez/golang-calculator-backend/libs/repositories/mocks"
	"github.com/JavierFVasquez/golang-calculator-backend/libs/services"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestBasicOperationServiceBasicOperation(t *testing.T) {

	recordRepo := new(mocks.MockRecordRepositoryIF)
	operationRepo := new(mocks.MockOperationRepositoryIF)

	operationInfo := models.Operation{
		Cost: 2000,
	}
	zeroResult := 0.0
	userMostRecentRecord := models.Record{
		ID: "6495a8166fa90092d1ed41ff",
		Amount: 0,
		DeletedAt: nil,
		Operation: models.Operation{
			Operation: "ADD_BALANCE", 
			Operand1: 0, 
			Operand2: 0, 
			Result: &zeroResult, 
			StringResult: nil,
			Cost: 30000,
		},
		User: models.User{
			ID: "04e88408-7041-708c-5a39-ff135e9ad400",
			Name: "Test User", 
			Email: "test@example.com", 
			Balance: 30000,
		},
		UserBalance: 30000,
	}
	result := 5.0
	expectedRecord := models.Record{
		ID:"",
		User: models.User{
				ID: "04e88408-7041-708c-5a39-ff135e9ad400",
				Name: "Test User", 
				Email: "test@example.com", 
				Balance: 30000,
			}, 
			Operation:models.Operation{
				Operation: models.ADDITION,
				Operand1:  2,
				Operand2:  3, 
				Result: &result, 
				StringResult: nil, 
				Cost: 2000,
			}, 
			Amount:0, 
			UserBalance:0, 
			CreatedAt: nil, 
			DeletedAt: nil,
		}
	
	recordRepo.On("GetUserMostRecentRecord", context.Background()).Return(&userMostRecentRecord, nil)
	recordRepo.On("CreateNewRecord", context.Background(), mock.Anything, mock.Anything).Return(&expectedRecord, nil)
	operationRepo.On("GetOperation", context.Background(), mock.Anything).Return(&operationInfo, nil).Once()
	
	service := services.NewBasicOperationService(recordRepo, operationRepo, &zerolog.Logger{})
	
	operation := models.Operation{
		Operation: models.ADDITION,
		Operand1:  2,
		Operand2:  3,
	}
	record, _ := service.BasicOperation(context.Background(), operation)
	assert.Equal(t, &expectedRecord, record)

	secondOperationInfo := models.Operation{
		Cost: 50000,
	}
	
	operationRepo.On("GetOperation", context.Background(), mock.Anything).Return(&secondOperationInfo, nil).Once()
	serviceFail := services.NewBasicOperationService(recordRepo, operationRepo, &zerolog.Logger{})
	_, err := serviceFail.BasicOperation(context.Background(), operation)

	assert.Error(t, *err)


	recordRepo.AssertExpectations(t)
	operationRepo.AssertExpectations(t)
}