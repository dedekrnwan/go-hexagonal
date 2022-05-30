package orm

import (
	"context"
	"go-boiler-clean/dto"
	"go-boiler-clean/entity"

	"gorm.io/gorm"
)

type (
	Transaction interface {
		Base[entity.Transaction, dto.Transaction]
		Some(ctx context.Context, id int) error
	}

	transaction struct {
		Base[entity.Transaction, dto.Transaction]
	}
)

func NewTransaction(connectionGrom *gorm.DB) Transaction {
	base := NewBase(connectionGrom, entity.Transaction{}, dto.Transaction{})
	return &transaction{
		base,
	}
}

func (m *transaction) Some(ctx context.Context, id int) error {
	return nil
}
