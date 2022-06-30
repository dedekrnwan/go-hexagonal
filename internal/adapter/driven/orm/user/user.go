package user

import (
	"context"
	"go-boiler-clean/internal/adapter/driven/orm/base"
	"go-boiler-clean/internal/adapter/driven/orm/entity"
	"go-boiler-clean/internal/model/sample"

	"gorm.io/gorm"
)

type (
	User interface {
		base.Base[entity.User, sample.UserEntity]
		CountByEmail(ctx context.Context, email string) (int64, error)
	}

	user struct {
		base.Base[entity.User, sample.UserEntity]
	}
)

func NewUser(connectionGrom *gorm.DB) User {
	base := base.NewBase(connectionGrom, entity.User{}, sample.UserEntity{})
	return &user{
		base,
	}
}

func (m *user) CountByEmail(ctx context.Context, email string) (count int64, err error) {
	err = m.GetDBConnector().Model(entity.User{}).WithContext(ctx).Where("email", email).Count(&count).Error
	return
}
