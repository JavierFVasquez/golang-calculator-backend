package services

import (
	"context"
	"errors"

	"github.com/JavierFVasquez/golang-calculator-backend/libs/models"
	"github.com/JavierFVasquez/golang-calculator-backend/libs/repositories"

	"github.com/rs/zerolog"
)

type BalanceService struct {
	recordRepo repositories.RecordRepositoryIF
	logger     *zerolog.Logger
}

func NewBalanceService(
	recordRepo repositories.RecordRepositoryIF,
	log *zerolog.Logger,
) *BalanceService {
	l := log.With().Str("resource", "create_square_root_service").Logger()

	return &BalanceService{
		recordRepo: recordRepo,
		logger:     &l,
	}
}

func (service *BalanceService) GetBalance(
	ctx context.Context,
) (*models.User, *error) {

	lastUserRecord, err := service.recordRepo.GetUserMostRecentRecord(ctx)
	if err != nil {
		err := errors.New("GET_USER_BALANCE_ERROR")
		return nil, &err
	}
	lastUserRecord.User.Balance = lastUserRecord.UserBalance
	return &lastUserRecord.User, nil

}
