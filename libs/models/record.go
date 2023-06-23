package models

import (
	"time"

	"github.com/google/uuid"
)

type Record struct {
	ID          string     `faker:"uuid_hyphenated" json:"_id" bson:"_id"`
	User        User       `faker:"user" json:"user" bson:"user"`
	Operation   Operation  `faker:"operation" json:"operation" bson:"operation"`
	Amount      float64    `json:"amount" bson:"amount"`
	UserBalance float64    `json:"userBalance" bson:"userBalance"`
	CreatedAt   *time.Time `json:"createdAt" bson:"createdAt"`
	DeletedAt   *time.Time `json:"deletedAt" bson:"deletedAt"`
}

func NewRecord(user User, operation Operation, amount, userBalance float64) *Record {
	now := time.Now()

	record := Record{
		ID:          uuid.New().String(),
		User:        user,
		Operation:   operation,
		Amount:      amount,
		UserBalance: userBalance,
		CreatedAt:   &now,
		DeletedAt:   nil,
	}
	return &record
}
