package transaction

import (
	"go-boiler-clean/internal/adapter/driven/orm/base"
	"go-boiler-clean/internal/adapter/driven/orm/entity"
	"go-boiler-clean/internal/model/sample"

	"gorm.io/gorm"
)

type (
	Transaction interface {
		base.Base[entity.Transaction, sample.TransactionEntity]
	}

	transaction struct {
		base.Base[entity.Transaction, sample.TransactionEntity]
	}
)

func NewTransaction(connectionGrom *gorm.DB) Transaction {
	base := base.NewBase(connectionGrom, entity.Transaction{}, sample.TransactionEntity{})
	return &transaction{
		base,
	}
}
