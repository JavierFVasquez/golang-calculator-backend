package services

import (
	"context"

	"github.com/JavierFVasquez/truenorth-calculator-backend/libs/models"
)

//go:generate mockery --name RecordServiceIF --structname MockRecordServiceIF --filename addition_service_mock.go
type RecordServiceIF interface {
	GetRecordList(ctx context.Context, options *models.PaginationParams) (*models.PaginatedResponse[models.Record], *error)
}
