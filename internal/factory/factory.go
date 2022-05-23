package factory

import (
	"go-boiler-clean/database"
	modelPsqlGorm "go-boiler-clean/internal/model/psqlGorm"

	"gorm.io/gorm"
)

type (
	Factory struct {
		ConnectionGorm *gorm.DB

		Model struct {
			ModelPsqlGormUser modelPsqlGorm.User
		}
	}
)

func NewFactory() *Factory {
	f := &Factory{}
	f.SetupDb()
	f.SetupModel()

	return f
}

func (f *Factory) SetupDb() {
	conn := "postgres"
	db, err := database.Connection(conn)
	if err != nil {
		panic("Failed setup db, connection is undefined")
	}
	dbGorm, ok := db.(*gorm.DB)
	if !ok {
		panic("Failed setup db, db is not gorm")
	}
	f.ConnectionGorm = dbGorm
}

func (f *Factory) SetupModel() {
	if f.ConnectionGorm == nil {
		panic("Failed setup model, db is undefined")
	}

	// modelPsqlGormBase := modelPsqlGorm.NewBase[entity.Entities](dbGorm)
	f.Model.ModelPsqlGormUser = modelPsqlGorm.NewUser(f.ConnectionGorm)
}
