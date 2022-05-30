package model

import (
	"context"
	"go-boiler-clean/dto"
	"go-boiler-clean/entity"
)

type (
	User interface {
		Base[entity.User, dto.User]
		Some(ctx context.Context, id int) error
	}
)
