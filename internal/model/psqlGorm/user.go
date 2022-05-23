package modelPsqlGorm

import (
	"context"
	"go-boiler-clean/entity"

	"gorm.io/gorm"
)

type (
	User interface {
		Base[entity.User]
		Some(ctx context.Context, id int) error
	}

	user struct {
		Base[entity.User]
	}
)

func NewUser(connectionGrom *gorm.DB) User {
	base := NewBase(connectionGrom, entity.User{})
	return &user{
		base,
	}
}

func (m *user) Some(ctx context.Context, id int) error {
	return nil
}
