package usecase

import (
	"context"
	"go-boiler-clean/dto"
	"go-boiler-clean/internal/model"
)

type (
	Base[T any, Y any] interface {
		Find(ctx context.Context, search string, filters []dto.Filter, ascending []string, descending []string, pagination dto.Pagination) ([]Y, *dto.PaginationInfo, error)
		FindOne(ctx context.Context, id int) (*Y, error)
		CreateOne(ctx context.Context, data *T) (*Y, error)
		UpdateOne(ctx context.Context, id int, data *T) (*Y, error)
		DeleteOne(ctx context.Context, id int) error
	}

	base[T any, Y any] struct {
		modelBase model.Base[T, Y]
	}
)

func NewBase[T any, Y any](modelBase model.Base[T, Y]) Base[T, Y] {
	return &base[T, Y]{
		modelBase,
	}
}

func (u *base[T, Y]) Find(ctx context.Context, search string, filters []dto.Filter, ascending []string, descending []string, pagination dto.Pagination) ([]Y, *dto.PaginationInfo, error) {
	return u.modelBase.Find(ctx, search, filters, ascending, descending, pagination)
}

func (u *base[T, Y]) FindOne(ctx context.Context, id int) (*Y, error) {
	return u.modelBase.FindOne(ctx, id)

}

func (u *base[T, Y]) CreateOne(ctx context.Context, data *T) (*Y, error) {
	return u.modelBase.CreateOne(ctx, data)
}

func (u *base[T, Y]) UpdateOne(ctx context.Context, id int, data *T) (*Y, error) {
	return u.modelBase.UpdateOne(ctx, id, data)
}

func (u *base[T, Y]) DeleteOne(ctx context.Context, id int) error {
	return u.modelBase.DeleteOne(ctx, id)
}
