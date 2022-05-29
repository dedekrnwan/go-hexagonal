package orm

import (
	"context"
	"fmt"
	"go-boiler-clean/dto"
	"math"

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

		//building only
		buildFilter(ctx context.Context, tx *gorm.DB, filters []dto.Filter)
		// buildOrder(ctx context.Context, tx *gorm.DB, ascending []string, descending []string)
		buildPagination(ctx context.Context, tx *gorm.DB, pagination dto.Pagination) *dto.PaginationInfo
	}

	base[T any] struct {
		connectionGrom *gorm.DB
		entity         T
	}
)

func NewBase[T any](connectionGrom *gorm.DB, entity T) Base[T] {
	return &base[T]{
		connectionGrom,
		entity,
	}
}

func (m *base[T]) GetDBConnector() *gorm.DB {
	return m.connectionGrom
}

func (m *base[T]) Find(ctx context.Context, search string, filters []dto.Filter, ascending []string, descending []string, pagination dto.Pagination) ([]T, *dto.PaginationInfo, error) {
	query := m.connectionGrom.Model(m.entity)

	m.buildFilter(ctx, query, filters)
	info := m.buildPagination(ctx, query, pagination)

	result := []T{}
	err := query.Find(&result).Error

	if err != nil {
		return nil, info, err
	}
	return result, info, nil
}

func (m *base[T]) FindOne(ctx context.Context, id int) (*T, error) {
	query := m.connectionGrom.Model(m.entity)
	result := new(T)
	err := query.Where("id", id).First(result).Error
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (m *base[T]) CreateOne(ctx context.Context, data *T) (*T, error) {
	query := m.connectionGrom.Model(m.entity)
	err := query.Create(data).Error
	return data, err
}
func (m *base[T]) CreateMany(ctx context.Context, data []T) ([]T, error) {
	err := m.connectionGrom.Model(m.entity).Create(&data).Error
	return data, err
}

func (m *base[T]) UpdateOne(ctx context.Context, id int, data *T) (*T, error) {
	err := m.connectionGrom.Model(data).Updates(data).Error
	return data, err
}

func (m *base[T]) DeleteOne(ctx context.Context, id int) error {
	return m.connectionGrom.Model(m.entity).Where("id = ?", id).Error
}

func (m *base[T]) buildFilter(ctx context.Context, tx *gorm.DB, filters []dto.Filter) {
	for _, v := range filters {
		if v.Operator == "like" {
			v.Value = fmt.Sprintf("%s%s%s", "%", v.Value, "%")
		}
		tx.Where(fmt.Sprintf("%s %s ?", v.Field, v.Operator), v.Value)
	}
}

// func (m *base[T]) buildOrder(ctx context.Context, tx *gorm.DB, ascending []string, descending []string) {
// }

func (m *base[T]) buildPagination(ctx context.Context, tx *gorm.DB, pagination dto.Pagination) *dto.PaginationInfo {
	info := &dto.PaginationInfo{}
	if pagination.Page != nil {
		limit := 10
		if pagination.Limit != nil {
			limit = *pagination.Limit
		}
		page := 0
		if *pagination.Page >= 0 {
			page = *pagination.Page
		}

		tx.Count(&info.Count)
		offset := (page - 1) * limit
		tx.Limit(limit).Offset(offset)
		info.TotalPage = int64(math.Ceil(float64(info.Count) / float64(limit)))

		info.Pagination = pagination
	}
	return info
}
