package repositories

import (
	"context"

	"github.com/JavierFVasquez/golang-calculator-backend/libs/models"
)

//go:generate mockery --name OperationRepositoryIF --structname MockOperationRepositoryIF --filename operation_repository_mock.go
type OperationRepositoryIF interface {
	GetOperation(ctx context.Context, id models.Operator) (*models.Operation, error)
}
