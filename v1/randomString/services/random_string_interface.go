package services

import (
	"context"

	"github.com/JavierFVasquez/truenorth-calculator-backend/libs/models"
)

//go:generate mockery --name RandomStringServiceIF --structname MockRandomStringServiceIF --filename addition_service_mock.go
type RandomStringServiceIF interface {
	RandomString(ctx context.Context, operation models.Operation) (*models.Record, *error)
}
