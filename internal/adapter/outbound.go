package adapter

import (
	"go-boiler-clean/internal/adapter/outbound/database"
	"go-boiler-clean/internal/adapter/outbound/orm"

	"gorm.io/gorm"
)

type (
	OutBound struct {
		Orm *orm.Orm
	}
)

func NewOutBound() (*OutBound, error) {
	database.Init()

	gormConnection, err := database.Connection[gorm.DB]("postgres")
	if err != nil {
		return nil, err
	}
	ormInstance := orm.New(gormConnection)

	return &OutBound{
		Orm: ormInstance,
	}, nil
}
