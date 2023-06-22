package repositories

import (
	"context"

	"github.com/JavierFVasquez/truenorth-calculator-backend/libs/models"
)

//go:generate mockery --name RecordRepositoryIF --structname MockRecordRepositoryIF --filename record_repository_mock.go
type RecordRepositoryIF interface {
	CreateNewRecord(ctx context.Context, operation models.Operation, newUserBalance float64) (*models.Record, error)
	GetRecordsByUserId(ctx context.Context, id string, options *models.PaginationParams) (*models.PaginatedResponse[models.Record], error)
}
