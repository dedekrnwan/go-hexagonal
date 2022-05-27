package usecase

import "go-boiler-clean/internal/model"

type (
	Usecase struct {
		UsecaseUser User
	}
)

func New(
	modelUser model.User,
) *Usecase {
	return &Usecase{
		UsecaseUser: NewUser(modelUser),
	}
}
