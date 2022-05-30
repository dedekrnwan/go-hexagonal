package model

import (
	"context"
	"go-boiler-clean/dto"
	"go-boiler-clean/entity"
)

type (
	User interface {
		Base[entity.User, dto.User]
		CountByEmail(ctx context.Context, email string) (int64, error)
	}
)
