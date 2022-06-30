package usecase

import (
	"go-boiler-clean/internal/model/sample"
)

type (
	User interface {
		Base[sample.UserEntity, sample.UserEntity]
	}

	user struct {
		Base[sample.UserEntity, sample.UserEntity]

		mdlUser sample.User
	}
)

func NewUser(
	mdlUser sample.User,
) User {
	return &user{
		Base:    NewBase[sample.UserEntity, sample.UserEntity](mdlUser),
		mdlUser: mdlUser,
	}
}
