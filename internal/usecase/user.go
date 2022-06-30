package usecase

import (
	"go-boiler-clean/internal/model/sample"
)

type (
	User interface {
		Common[sample.UserEntity, sample.UserEntity]
	}

	user struct {
		Common[sample.UserEntity, sample.UserEntity]

		mdlUser sample.User
	}
)

func NewUser(
	mdlUser sample.User,
) User {
	return &user{
		Common:  NewCommon[sample.UserEntity, sample.UserEntity](mdlUser),
		mdlUser: mdlUser,
	}
}
