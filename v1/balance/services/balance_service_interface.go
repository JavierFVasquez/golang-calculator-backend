package services

import (
	"context"

	"github.com/JavierFVasquez/golang-calculator-backend/libs/models"
)

//go:generate mockery --name NewBalanceServiceIF --structname MockNewBalanceServiceIF --filename addition_service_mock.go
type NewBalanceServiceIF interface {
	GetBalance(ctx context.Context) (*models.User, *error)
}
