package usecase

import (
	"context"
	"go-boiler-clean/dto"
	modelPsqlGorm "go-boiler-clean/internal/model/psqlGorm"
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
		modelPsqlGormBase modelPsqlGorm.Base[T]
	}
)

func NewBase[T any](modelPsqlGormBase modelPsqlGorm.Base[T]) Base[T] {
	return &base[T]{
		modelPsqlGormBase,
	}
}

func (u *base[T]) Find(ctx context.Context, search string, filters []dto.Filter, ascending []string, descending []string, pagination dto.Pagination) ([]T, *dto.PaginationInfo, error) {
	return u.modelPsqlGormBase.Find(ctx, search, filters, ascending, descending, pagination)
}

func (u *base[T]) FindOne(ctx context.Context, id int) (*T, error) {
	return u.modelPsqlGormBase.FindOne(ctx, id)

}

func (u *base[T]) CreateOne(ctx context.Context, data *T) (*T, error) {
	return u.modelPsqlGormBase.CreateOne(ctx, data)
}

func (u *base[T]) UpdateOne(ctx context.Context, id int, data *T) (*T, error) {
	return u.modelPsqlGormBase.UpdateOne(ctx, id, data)
}

func (u *base[T]) DeleteOne(ctx context.Context, id int) error {
	return u.modelPsqlGormBase.DeleteOne(ctx, id)
}
