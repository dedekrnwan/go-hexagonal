package orm

import (
	"go-boiler-clean/internal/adapter/driven/orm/database"
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

func New() (o *Orm, err error) {
	connection, err := database.Connection[gorm.DB]("postgres")
	o = &Orm{
		User:        repository.NewUser(connection),
		Transaction: repository.NewTransaction(connection),
	}

	return
}
