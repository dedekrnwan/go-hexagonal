package usecase

import (
	"go-boiler-clean/internal/model/sample"
)

type (
	Usecase struct {
		User User
	}
)

func New(
	mdlUser sample.User,
) *Usecase {
	return &Usecase{
		User: NewUser(mdlUser),
	}
}
