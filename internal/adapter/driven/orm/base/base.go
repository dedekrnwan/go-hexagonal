package base

import (
	"context"
	"encoding/json"
	"fmt"
	"go-boiler-clean/dto"
	"math"

	"gorm.io/gorm"
)

type (
	Base[T any, Y any] interface {
		GetDBConnector() *gorm.DB

		Find(ctx context.Context, search string, filters []dto.Filter, ascending []string, descending []string, pagination dto.Pagination) ([]Y, *dto.PaginationInfo, error)
		FindOne(ctx context.Context, id int) (*Y, error)
		CreateOne(ctx context.Context, data *Y) (*Y, error)
		CreateMany(ctx context.Context, payload []Y) ([]Y, error)
		UpdateOne(ctx context.Context, id int, data *Y) (*Y, error)
		DeleteOne(ctx context.Context, id int) error

		//building only
		buildFilter(ctx context.Context, tx *gorm.DB, filters []dto.Filter)
		// buildOrder(ctx context.Context, tx *gorm.DB, ascending []string, descending []string)
		buildPagination(ctx context.Context, tx *gorm.DB, pagination dto.Pagination) *dto.PaginationInfo
	}

	base[T any, Y any] struct {
		connectionGrom *gorm.DB
		entity         T
	}
)

func NewBase[T any, Y any](connectionGrom *gorm.DB, entity T, dt Y) Base[T, Y] {
	return &base[T, Y]{
		connectionGrom,
		entity,
	}
}

func (m *base[T, Y]) GetDBConnector() *gorm.DB {
	return m.connectionGrom
}

func (m *base[T, Y]) Find(ctx context.Context, search string, filters []dto.Filter, ascending []string, descending []string, pagination dto.Pagination) ([]Y, *dto.PaginationInfo, error) {
	query := m.connectionGrom.Model(m.entity)

	m.buildFilter(ctx, query, filters)
	info := m.buildPagination(ctx, query, pagination)

	result := []Y{}
	err := query.Find(&result).Error

	if err != nil {
		return nil, info, err
	}
	return result, info, nil
}

func (m *base[T, Y]) FindOne(ctx context.Context, id int) (*Y, error) {
	query := m.connectionGrom.Model(m.entity)
	result := new(Y)
	err := query.Where("id", id).First(result).Error
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (m *base[T, Y]) CreateOne(ctx context.Context, data *Y) (*Y, error) {
	query := m.connectionGrom.Model(m.entity)
	err := query.Create(data).Error
	if err != nil {
		return nil, err
	}
	result := new(Y)

	byteJson, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(byteJson, result)
	if err != nil {
		return nil, err
	}

	return result, err
}

func (m *base[T, Y]) CreateMany(ctx context.Context, data []Y) ([]Y, error) {
	err := m.connectionGrom.Model(m.entity).Create(&data).Error
	if err != nil {
		return nil, err
	}

	result := make([]Y, 0)

	byteJson, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(byteJson, result)
	if err != nil {
		return nil, err
	}

	return result, err
}

func (m *base[T, Y]) UpdateOne(ctx context.Context, id int, data *Y) (*Y, error) {
	err := m.connectionGrom.Model(data).Updates(data).Error
	if err != nil {
		return nil, err
	}
	result := new(Y)

	byteJson, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(byteJson, result)
	if err != nil {
		return nil, err
	}
	return result, err
}

func (m *base[T, Y]) DeleteOne(ctx context.Context, id int) error {
	return m.connectionGrom.Model(m.entity).Where("id = ?", id).Error
}

func (m *base[T, Y]) buildFilter(ctx context.Context, tx *gorm.DB, filters []dto.Filter) {
	for _, v := range filters {
		if v.Operator == "like" {
			v.Value = fmt.Sprintf("%s%s%s", "%", v.Value, "%")
		}
		tx.Where(fmt.Sprintf("%s %s ?", v.Field, v.Operator), v.Value)
	}
}

// func (m *base[T, Y]) buildOrder(ctx context.Context, tx *gorm.DB, ascending []string, descending []string) {
// }

func (m *base[T, Y]) buildPagination(ctx context.Context, tx *gorm.DB, pagination dto.Pagination) *dto.PaginationInfo {
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
