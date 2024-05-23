package services

import (
	"context"

	"github.com/JavierFVasquez/golang-calculator-backend/libs/models"
)

//go:generate mockery --name BasicOperationServiceIF --structname MockBasicOperationServiceIF --filename basic_operation_service_mock.go
type BasicOperationServiceIF interface {
	BasicOperation(ctx context.Context, operation models.Operation) (*models.Record, *error)
}
