package factory

import (
	adapter "go-boiler-clean/internal/adapter"
	"go-boiler-clean/internal/database"
	"go-boiler-clean/internal/usecase"

	"gorm.io/gorm"
)

type (
	Factory struct {
		ConnectionGorm *gorm.DB

		Adapter struct {
			OutBound *adapter.OutBound
			InBound  *adapter.InBound
		}

		Usecase *usecase.Usecase
	}
)

func NewFactory() *Factory {
	f := &Factory{}
	f.setupDb()
	// f.SetupModelPsqlGorm()
	f.setupAdapterOutBound()
	f.setupUseCase()
	f.setupAdapterInBound()

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

func (f *Factory) setupUseCase() {
	f.Usecase = usecase.New(
		f.Adapter.OutBound.Orm.User,
	)
}

func (f *Factory) setupAdapterInBound() {
	f.Adapter.InBound = adapter.NewInBound(
		f.Usecase.UsecaseUser,
	)
}
