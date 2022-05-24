package usecase

import (
	"context"
	"go-boiler-clean/dto"
	"go-boiler-clean/internal/model"
)

type (
	Base[T any] interface {
		Find(ctx context.Context, search string, filters []dto.Filter, ascending []string, descending []string, pagination dto.Pagination) ([]T, *dto.PaginationInfo, error)
		FindOne(ctx context.Context, id int) (*T, error)
		CreateOne(ctx context.Context, data *T) (*T, error)
		UpdateOne(ctx context.Context, id int, data *T) (*T, error)
		DeleteOne(ctx context.Context, id int) error
	}

	base[T any] struct {
		modelBase model.Base[T]
	}
)

func NewBase[T any](modelBase model.Base[T]) Base[T] {
	return &base[T]{
		modelBase,
	}
}

func (u *base[T]) Find(ctx context.Context, search string, filters []dto.Filter, ascending []string, descending []string, pagination dto.Pagination) ([]T, *dto.PaginationInfo, error) {
	return u.modelBase.Find(ctx, search, filters, ascending, descending, pagination)
}

func (u *base[T]) FindOne(ctx context.Context, id int) (*T, error) {
	return u.modelBase.FindOne(ctx, id)

}

func (u *base[T]) CreateOne(ctx context.Context, data *T) (*T, error) {
	return u.modelBase.CreateOne(ctx, data)
}

func (u *base[T]) UpdateOne(ctx context.Context, id int, data *T) (*T, error) {
	return u.modelBase.UpdateOne(ctx, id, data)
}

func (u *base[T]) DeleteOne(ctx context.Context, id int) error {
	return u.modelBase.DeleteOne(ctx, id)
}
