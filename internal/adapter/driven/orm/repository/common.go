package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"go-boiler-clean/dto"
	"math"
	"reflect"
	"strings"

	"gorm.io/gorm"
)

type (
	Common[T any, Y any] interface {
		GetDBConnector() *gorm.DB

		Count(ctx context.Context, filters []dto.Filter) (int64, error)
		Find(ctx context.Context, search string, filters []dto.Filter, ascending []string, descending []string, pagination dto.Pagination, preloads []string, excludesOrder ...string) ([]Y, *dto.PaginationInfo, error)
		FindOne(ctx context.Context, id int, preloads []string) (*Y, error)
		CreateOne(ctx context.Context, data *Y) (*Y, error)
		CreateMany(ctx context.Context, payload []Y) ([]Y, error)
		UpdateOne(ctx context.Context, id int, data *Y) (*Y, error)
		DeleteOne(ctx context.Context, id int) error

		//building only
		buildFilter(ctx context.Context, tx *gorm.DB, filters []dto.Filter)
		buildPreload(ctx context.Context, tx *gorm.DB, preloads []string)
		buildOrder(ctx context.Context, tx *gorm.DB, ascending []string, descending []string, excludes ...string)
		BuildPagination(ctx context.Context, tx *gorm.DB, pagination dto.Pagination) *dto.PaginationInfo
	}

	common[T any, Y any] struct {
		connectionGrom *gorm.DB
		entity         T
	}
)

func NewCommon[T any, Y any](connectionGrom *gorm.DB, entity T, dt Y) Common[T, Y] {
	return &common[T, Y]{
		connectionGrom,
		entity,
	}
}

func (m *common[T, Y]) GetDBConnector() *gorm.DB {
	return m.connectionGrom
}

func (m *common[T, Y]) Count(ctx context.Context, filters []dto.Filter) (count int64, err error) {
	query := m.connectionGrom.Model(m.entity)

	m.buildFilter(ctx, query, filters)
	err = query.Count(&count).Error

	return
}

func (m *common[T, Y]) Find(ctx context.Context, search string, filters []dto.Filter, ascending []string, descending []string, pagination dto.Pagination, preloads []string, excludesOrder ...string) ([]Y, *dto.PaginationInfo, error) {
	query := m.connectionGrom.Model(m.entity)

	m.buildFilter(ctx, query, filters)
	m.buildOrder(ctx, query, ascending, descending, excludesOrder...)
	m.buildPreload(ctx, query, preloads)
	info := m.BuildPagination(ctx, query, pagination)

	result := []Y{}
	err := query.Find(&result).Error

	if err != nil {
		return nil, info, err
	}
	return result, info, nil
}

func (m *common[T, Y]) FindOne(ctx context.Context, id int, preloads []string) (*Y, error) {
	query := m.connectionGrom.Model(m.entity)
	m.buildPreload(ctx, query, preloads)

	result := new(Y)
	err := query.Where("id", id).First(result).Error
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (m *common[T, Y]) CreateOne(ctx context.Context, data *Y) (*Y, error) {
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

func (m *common[T, Y]) CreateMany(ctx context.Context, data []Y) ([]Y, error) {
	err := m.connectionGrom.Model(m.entity).Create(&data).Error
	if err != nil {
		return nil, err
	}

	result := make([]Y, 0)

	byteJson, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(byteJson, &result)
	if err != nil {
		return nil, err
	}

	return result, err
}

func (m *common[T, Y]) UpdateOne(ctx context.Context, id int, data *Y) (*Y, error) {
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

func (m *common[T, Y]) DeleteOne(ctx context.Context, id int) error {
	return m.connectionGrom.Model(m.entity).Where("id = ?", id).Error
}

func (m *common[T, Y]) buildFilter(ctx context.Context, tx *gorm.DB, filters []dto.Filter) {
	for _, v := range filters {
		if v.Operator == "like" {
			v.Value = fmt.Sprintf("%s%s%s", "%", v.Value, "%")
		}

		switch strings.ToLower(v.Operator) {
		case "in":
			tx.Where(fmt.Sprintf("%s %s (?)", v.Field, v.Operator), v.Value)
		default:
			if v.Operator == "is not" && v.Value == nil {
				tx.Where(fmt.Sprintf("%s is not null", v.Field))
			} else if v.Operator == "is" && v.Value == nil {
				tx.Where(fmt.Sprintf("%s is null", v.Field))
			} else {
				tx.Where(fmt.Sprintf("%s %s ?", v.Field, v.Operator), v.Value)
			}
		}
	}
}

func (m *common[T, Y]) buildPreload(ctx context.Context, tx *gorm.DB, preloads []string) {
	for _, v := range preloads {
		tx.Preload(v)
	}
}

func (m *common[T, Y]) buildOrder(ctx context.Context, tx *gorm.DB, ascending []string, descending []string, excludes ...string) {
	columns := reflect.ValueOf(m.entity)
	mapExcludes := make(map[string]string, 0)
	mapAsc := make(map[string]string, 0)
	mapDesc := make(map[string]string, 0)

	for _, v := range excludes {
		mapExcludes[v] = v
	}
	for _, v := range ascending {
		mapAsc[v] = v
	}
	for _, v := range descending {
		mapDesc[v] = v
	}

	ascending = []string{}
	descending = []string{}

loopColumns:
	for i := 0; i < columns.NumField(); i++ {
		column := columns.Type().Field(i).Tag.Get("json")

		if _, ok := mapExcludes[column]; !ok {
			continue loopColumns
		}

		if v, ok := mapAsc[column]; !ok {
			ascending = append(ascending, v)
		}

		if v, ok := mapDesc[column]; !ok {
			descending = append(descending, v)
		}
	}

	if len(ascending) > 0 {
		cols := strings.Join(ascending, ",")
		tx.Order(fmt.Sprintf("%s asc", cols))
	}
	if len(descending) > 0 {
		cols := strings.Join(descending, ",")
		tx.Order(fmt.Sprintf("%s desc", cols))
	}
}

func (m *common[T, Y]) BuildPagination(ctx context.Context, tx *gorm.DB, pagination dto.Pagination) *dto.PaginationInfo {
	info := &dto.PaginationInfo{}
	if pagination.Page != 0 {
		limit := 10
		if pagination.Limit != 0 {
			limit = pagination.Limit
		}
		page := 0
		if pagination.Page >= 0 {
			page = pagination.Page
		}

		tx.Count(&info.Count)
		offset := (page - 1) * limit
		tx.Limit(limit).Offset(offset)
		info.TotalPage = int64(math.Ceil(float64(info.Count) / float64(limit)))

		info.Pagination = pagination
	}
	return info
}
