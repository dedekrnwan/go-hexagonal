package orm

import (
	"go-boiler-clean/internal/model"

	"gorm.io/gorm"
)

type (
	Orm struct {
		User        model.User
		Transaction model.Transaction
	}
)

func New(connection *gorm.DB) *Orm {
	return &Orm{
		User:        NewUser(connection),
		Transaction: NewTransaction(connection),
	}
}
