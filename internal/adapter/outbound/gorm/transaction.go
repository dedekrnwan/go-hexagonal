package adapterOutBoundGorm

import (
	"context"
	"go-boiler-clean/entity"

	"gorm.io/gorm"
)

type (
	Transaction interface {
		Base[entity.Transaction]
		Some(ctx context.Context, id int) error
	}

	transaction struct {
		Base[entity.Transaction]
	}
)

func NewTransaction(connectionGrom *gorm.DB) Transaction {
	base := NewBase(connectionGrom, entity.Transaction{})
	return &transaction{
		base,
	}
}

func (m *transaction) Some(ctx context.Context, id int) error {
	return nil
}
