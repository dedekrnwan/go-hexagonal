package model

import (
	"context"
	"go-boiler-clean/dto"

	"gorm.io/gorm"
)

type (
	Base[T any] interface {
		GetDBConnector() *gorm.DB

		Find(ctx context.Context, search string, filters []dto.Filter, ascending []string, descending []string, pagination dto.Pagination) ([]T, *dto.PaginationInfo, error)
		FindOne(ctx context.Context, id int) (*T, error)
		CreateOne(ctx context.Context, data *T) (*T, error)
		CreateMany(ctx context.Context, data []T) ([]T, error)
		UpdateOne(ctx context.Context, id int, data *T) (*T, error)
		DeleteOne(ctx context.Context, id int) error
	}
)
