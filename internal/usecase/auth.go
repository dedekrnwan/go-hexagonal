package usecase

import (
	"context"
	"encoding/json"
	"errors"
	"go-boiler-clean/dto"
	"go-boiler-clean/entity"
	"go-boiler-clean/internal/model"
	"go-boiler-clean/pkg/util"
)

type (
	Auth interface {
		Register(ctx context.Context, data *dto.AuthRegister) (result *dto.User, err error)
	}

	auth struct {
		modelUser model.User
		hasher    util.Hasher
	}
)

func NewAuth(
	modelUser model.User,
) Auth {
	return &auth{
		modelUser: modelUser,
		hasher:    util.NewHasher(),
	}
}

func (u *auth) Register(ctx context.Context, data *dto.AuthRegister) (result *dto.User, err error) {
	countByEmail, err := u.modelUser.CountByEmail(ctx, data.Email)
	if err != nil {
		return
	}

	if countByEmail > 0 {
		return nil, errors.New("email has already been registered")
	}

	user := new(entity.User)
	byteJson, err := json.Marshal(data)
	if err != nil {
		return
	}

	err = json.Unmarshal(byteJson, user)
	if err != nil {
		return
	}

	isActive := true
	user.IsActive = &isActive
	user.Password, err = u.hasher.HashPassword(user.Password)
	if err != nil {
		return
	}

	result, err = u.modelUser.CreateOne(ctx, user)
	if err != nil {
		return
	}

	return
}
