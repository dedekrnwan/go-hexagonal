package common

import (
	"context"
	"go-boiler-clean/dto"
)

type (
	Common[T any, Y any] interface {
		Count(ctx context.Context, filters []dto.Filter) (int64, error)
		Find(ctx context.Context, search string, filters []dto.Filter, ascending []string, descending []string, pagination dto.Pagination, preloads []string, excludesOrder ...string) ([]Y, *dto.PaginationInfo, error)
		FindOne(ctx context.Context, id int, preloads []string) (*Y, error)
		CreateOne(ctx context.Context, data *Y) (*Y, error)
		CreateMany(ctx context.Context, data []Y) ([]Y, error)
		UpdateOne(ctx context.Context, id int, data *Y) (*Y, error)
		DeleteOne(ctx context.Context, id int) error
	}
)
