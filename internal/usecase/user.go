package usecase

import (
	"go-boiler-clean/dto"
	"go-boiler-clean/entity"
	"go-boiler-clean/internal/model"
)

type (
	User interface {
		Base[entity.User, dto.User]
	}

	user struct {
		Base[entity.User, dto.User]

		modelUser model.User
	}
)

func NewUser(
	modelUser model.User,
) User {
	return &user{
		Base:      NewBase[entity.User, dto.User](modelUser),
		modelUser: modelUser,
	}
}
