package orm

import (
	"context"
	"go-boiler-clean/dto"
	"go-boiler-clean/entity"

	"gorm.io/gorm"
)

type (
	User interface {
		Base[entity.User, dto.User]
		Some(ctx context.Context, id int) error
	}

	user struct {
		Base[entity.User, dto.User]
	}
)

func NewUser(connectionGrom *gorm.DB) User {
	base := NewBase(connectionGrom, entity.User{}, dto.User{})
	return &user{
		base,
	}
}

func (m *user) Some(ctx context.Context, id int) error {
	return nil
}
