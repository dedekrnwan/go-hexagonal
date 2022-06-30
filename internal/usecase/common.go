package usecase

import (
	"context"
	"go-boiler-clean/dto"
	model "go-boiler-clean/internal/model/common"
)

type (
	Common[T any, Y any] interface {
		Find(ctx context.Context, search string, filters []dto.Filter, ascending []string, descending []string, pagination dto.Pagination) ([]Y, *dto.PaginationInfo, error)
		FindOne(ctx context.Context, id int) (*Y, error)
		CreateOne(ctx context.Context, data *Y) (*Y, error)
		UpdateOne(ctx context.Context, id int, data *Y) (*Y, error)
		DeleteOne(ctx context.Context, id int) error
	}

	common[T any, Y any] struct {
		model model.Common[T, Y]
	}
)

func NewCommon[T any, Y any](model model.Common[T, Y]) Common[T, Y] {
	return &common[T, Y]{
		model,
	}
}

func (u *common[T, Y]) Find(ctx context.Context, search string, filters []dto.Filter, ascending []string, descending []string, pagination dto.Pagination) ([]Y, *dto.PaginationInfo, error) {
	return u.model.Find(ctx, search, filters, ascending, descending, pagination)
}

func (u *common[T, Y]) FindOne(ctx context.Context, id int) (*Y, error) {
	return u.model.FindOne(ctx, id)

}

func (u *common[T, Y]) CreateOne(ctx context.Context, data *Y) (*Y, error) {
	return u.model.CreateOne(ctx, data)
}

func (u *common[T, Y]) UpdateOne(ctx context.Context, id int, data *Y) (*Y, error) {
	return u.model.UpdateOne(ctx, id, data)
}

func (u *common[T, Y]) DeleteOne(ctx context.Context, id int) error {
	return u.model.DeleteOne(ctx, id)
}
