package adapter

import (
	"go-boiler-clean/internal/adapter/outbound/orm"

	"gorm.io/gorm"
)

type (
	OutBound struct {
		Orm *orm.Orm
	}
)

func NewOutBound(
	gormConnection *gorm.DB,
) *OutBound {
	ormInstance := orm.New(gormConnection)

	return &OutBound{
		Orm: ormInstance,
	}
}
