package repositories

import (
	"context"

	"github.com/JavierFVasquez/truenorth-calculator-backend/libs/models"
)

//go:generate mockery --name OperationRepositoryIF --structname MockOperationRepositoryIF --filename record_repository_mock.go
type OperationRepositoryIF interface {
	GetOperation(ctx context.Context, id models.Operator) (*models.Operation, error)
}
