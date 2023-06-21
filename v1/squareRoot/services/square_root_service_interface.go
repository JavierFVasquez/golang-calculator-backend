package services

import (
	"context"

	"github.com/JavierFVasquez/truenorth-calculator-backend/libs/models"
)

//go:generate mockery --name NewSquareRootServiceIF --structname MockNewSquareRootServiceIF --filename addition_service_mock.go
type NewSquareRootServiceIF interface {
	SquareRoot(ctx context.Context, operation models.Operation) (*models.Record, *error)
}
