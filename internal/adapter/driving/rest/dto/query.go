package dto

import (
	"fmt"
	"go-boiler-clean/dto"
	"net/http"
	"reflect"
	"strings"
)

type (
	HttpQuery[T any] struct {
		dto.Pagination
		Search     string       `json:"search" query:"search" bson:"search"`
		Ascending  []string     `json:"ascending" query:"ascending" bson:"ascending"`
		Descending []string     `json:"descending" query:"descending" bson:"descending"`
		Filters    []dto.Filter `json:"filters" bson:"filters"`
		entity     T            `json:"-"`

		request *http.Request `json:"-"`
	}
)

func NewHttpQuery[T any](r *http.Request, entity T) *HttpQuery[T] {
	return &HttpQuery[T]{
		entity:  entity,
		request: r,
	}
}

func (q *HttpQuery[T]) BindFilters() {
	queries := q.request.URL.Query()
	reflectEntity := reflect.ValueOf(q.entity)
	q.Filters = []dto.Filter{}

	for key, _ := range queries {
		if strings.HasPrefix(key, "filter") {
			field := strings.SplitN(key, ".", 3)
			if len(field) < 2 {
				continue
			}

			for i := 0; i < reflectEntity.NumField(); i++ {
				if reflectEntity.Type().Field(i).Tag.Get("json") == field[1] {
					filter := dto.Filter{
						Field: fmt.Sprintf("%s", field[1]),
						Value: queries[key][0],
					}
					if len(field) == 2 {
						filter.Operator = "="
					} else {
						filter.Operator = field[2]
					}

					q.Filters = append(q.Filters, filter)
				}
			}
		}
	}
}
