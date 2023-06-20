package services

import (
	"context"

	"github.com/JavierFVasquez/truenorth-calculator-backend/libs/models"
)

//go:generate mockery --name BasicOperationServiceIF --structname MockBasicOperationServiceIF --filename addition_service_mock.go
type BasicOperationServiceIF interface {
	BasicOperation(ctx context.Context, operation models.Operation) (*models.Record, *error)
}
