package usecase

import (
	"go-boiler-clean/entity"
	"go-boiler-clean/internal/factory"
	modelPsqlGorm "go-boiler-clean/internal/model/psqlGorm"
)

type (
	User interface {
		Base[entity.User]
	}

	user struct {
		Base[entity.User]

		modelPsqlGormUser modelPsqlGorm.User
	}
)

func NewUser(f *factory.Factory) User {
	return &user{
		Base:              f.Model.ModelPsqlGormUser,
		modelPsqlGormUser: f.Model.ModelPsqlGormUser,
	}
}

func (u *user) Custom() (string, error) {
	return "", nil
}
