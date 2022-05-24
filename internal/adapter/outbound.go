package adapter

import (
	adapterOutBoundGorm "go-boiler-clean/internal/adapter/outbound/gorm"

	"gorm.io/gorm"
)

type (
	OutBound struct {
		Gorm *adapterOutBoundGorm.Gorm
	}
)

func NewOutBound(connection *gorm.DB) *OutBound {
	return &OutBound{
		Gorm: adapterOutBoundGorm.New(connection),
	}
}
