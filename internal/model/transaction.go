package model

import (
	"context"
	"go-boiler-clean/entity"
)

type (
	Transaction interface {
		Base[entity.Transaction]
		Some(ctx context.Context, id int) error
	}
)
