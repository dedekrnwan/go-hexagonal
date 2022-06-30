package orm

import (
	"go-boiler-clean/internal/adapter/driven/orm/repository"
	sample "go-boiler-clean/internal/model/sample"

	"gorm.io/gorm"
)

type (
	Orm struct {
		User        sample.User
		Transaction sample.Transaction
	}
)

func New(connection *gorm.DB) *Orm {
	return &Orm{
		User:        repository.NewUser(connection),
		Transaction: repository.NewTransaction(connection),
	}
}
