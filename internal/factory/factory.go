package factory

import (
	adapter "go-boiler-clean/internal/adapter"
	"go-boiler-clean/internal/database"

	"gorm.io/gorm"
)

type (
	Factory struct {
		ConnectionGorm *gorm.DB

		Adapter struct {
			OutBound *adapter.OutBound
		}
	}
)

func NewFactory() *Factory {
	f := &Factory{}
	f.setupDb()
	// f.SetupModelPsqlGorm()
	f.setupAdapterOutBound()

	return f
}

func (f *Factory) setupDb() {
	database.Init()

	conn := "postgres"
	db, err := database.Connection[gorm.DB](conn)
	if err != nil {
		panic("Failed setup db, connection is undefined")
	}
	f.ConnectionGorm = db
}

func (f *Factory) setupAdapterOutBound() {
	f.Adapter.OutBound = adapter.NewOutBound(f.ConnectionGorm)
}

// func (f *Factory) SetupModelPsqlGorm() {
// 	if f.ConnectionGorm == nil {
// 		panic("Failed setup model, db is undefined")
// 	}

// 	f.Model.PsqlGorm.User = modelPsqlGorm.NewUser(f.ConnectionGorm)
// }
