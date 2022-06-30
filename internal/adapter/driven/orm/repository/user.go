package repository

import (
	"context"
	"go-boiler-clean/internal/adapter/driven/orm/entity"
	"go-boiler-clean/internal/model/sample"

	"gorm.io/gorm"
)

type (
	User interface {
		Common[entity.User, sample.UserEntity]
		CountByEmail(ctx context.Context, email string) (int64, error)
	}

	user struct {
		Common[entity.User, sample.UserEntity]
	}
)

func NewUser(connectionGrom *gorm.DB) User {
	common := NewCommon(connectionGrom, entity.User{}, sample.UserEntity{})
	return &user{
		common,
	}
}

func (m *user) CountByEmail(ctx context.Context, email string) (count int64, err error) {
	err = m.GetDBConnector().Model(entity.User{}).WithContext(ctx).Where("email", email).Count(&count).Error
	return
}
