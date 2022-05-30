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
		CountByEmail(ctx context.Context, email string) (int64, error)
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

func (m *user) CountByEmail(ctx context.Context, email string) (count int64, err error) {
	err = m.GetDBConnector().Model(entity.User{}).WithContext(ctx).Where("email", email).Count(&count).Error
	return
}
