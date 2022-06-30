package usecase

import (
	"context"
	"go-boiler-clean/dto"
	modelBase "go-boiler-clean/internal/model/base"
)

type (
	Base[T any, Y any] interface {
		Find(ctx context.Context, search string, filters []dto.Filter, ascending []string, descending []string, pagination dto.Pagination) ([]Y, *dto.PaginationInfo, error)
		FindOne(ctx context.Context, id int) (*Y, error)
		CreateOne(ctx context.Context, data *Y) (*Y, error)
		UpdateOne(ctx context.Context, id int, data *Y) (*Y, error)
		DeleteOne(ctx context.Context, id int) error
	}

	base[T any, Y any] struct {
		mdlBase modelBase.Base[T, Y]
	}
)

func NewBase[T any, Y any](mdlBase modelBase.Base[T, Y]) Base[T, Y] {
	return &base[T, Y]{
		mdlBase,
	}
}

func (u *base[T, Y]) Find(ctx context.Context, search string, filters []dto.Filter, ascending []string, descending []string, pagination dto.Pagination) ([]Y, *dto.PaginationInfo, error) {
	return u.mdlBase.Find(ctx, search, filters, ascending, descending, pagination)
}

func (u *base[T, Y]) FindOne(ctx context.Context, id int) (*Y, error) {
	return u.mdlBase.FindOne(ctx, id)

}

func (u *base[T, Y]) CreateOne(ctx context.Context, data *Y) (*Y, error) {
	return u.mdlBase.CreateOne(ctx, data)
}

func (u *base[T, Y]) UpdateOne(ctx context.Context, id int, data *Y) (*Y, error) {
	return u.mdlBase.UpdateOne(ctx, id, data)
}

func (u *base[T, Y]) DeleteOne(ctx context.Context, id int) error {
	return u.mdlBase.DeleteOne(ctx, id)
}
