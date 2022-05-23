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
		Base:              NewBase[entity.User](f.Model.PsqlGorm.User),
		modelPsqlGormUser: f.Model.PsqlGorm.User,
	}
}

func (u *user) Custom() (string, error) {
	return "", nil
}
