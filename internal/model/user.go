package model

import (
	"context"
	"go-boiler-clean/entity"
)

type (
	User interface {
		Base[entity.User]
		Some(ctx context.Context, id int) error
	}
)
