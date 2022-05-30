package model

import (
	"context"
	"go-boiler-clean/dto"
)

type (
	Base[T any, Y any] interface {
		Find(ctx context.Context, search string, filters []dto.Filter, ascending []string, descending []string, pagination dto.Pagination) ([]Y, *dto.PaginationInfo, error)
		FindOne(ctx context.Context, id int) (*Y, error)
		CreateOne(ctx context.Context, data *T) (*Y, error)
		CreateMany(ctx context.Context, data []T) ([]Y, error)
		UpdateOne(ctx context.Context, id int, data *T) (*Y, error)
		DeleteOne(ctx context.Context, id int) error
	}
)
