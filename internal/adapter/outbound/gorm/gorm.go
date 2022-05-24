package adapterOutBoundGorm

import (
	"go-boiler-clean/internal/model"

	"gorm.io/gorm"
)

type (
	Gorm struct {
		User        model.User
		Transaction model.Transaction
	}
)

func New(connection *gorm.DB) *Gorm {
	return &Gorm{
		User:        NewUser(connection),
		Transaction: NewTransaction(connection),
	}
}
