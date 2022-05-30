package usecase

import (
	"go-boiler-clean/entity"
	"go-boiler-clean/internal/model"
)

type (
	User interface {
		Base[entity.User]
	}

	user struct {
		Base[entity.User]

		modelUser model.User
	}
)

func NewUser(
	modelUser model.User,
) User {
	return &user{
		Base:      NewBase[entity.User](modelUser),
		modelUser: modelUser,
	}
}
