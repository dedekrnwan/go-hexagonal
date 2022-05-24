package adapter

import (
	"gorm.io/gorm"
)

type (
	InBound struct {
	}
)

func NewInBound(connection *gorm.DB) *InBound {
	return &InBound{}
}
