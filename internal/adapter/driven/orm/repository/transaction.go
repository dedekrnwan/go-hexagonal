package repository

import (
	"go-boiler-clean/internal/adapter/driven/orm/entity"
	"go-boiler-clean/internal/model/sample"

	"gorm.io/gorm"
)

type (
	Transaction interface {
		Common[entity.Transaction, sample.TransactionEntity]
	}

	transaction struct {
		Common[entity.Transaction, sample.TransactionEntity]
	}
)

func NewTransaction(connectionGrom *gorm.DB) Transaction {
	common := NewCommon(connectionGrom, entity.Transaction{}, sample.TransactionEntity{})
	return &transaction{
		common,
	}
}
