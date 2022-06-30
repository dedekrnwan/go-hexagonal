package factory

import (
	"go-boiler-clean/internal/adapter/driven/database"
	"go-boiler-clean/internal/adapter/driven/orm"
	"go-boiler-clean/internal/usecase"

	"gorm.io/gorm"
)

type (
	Factory struct {
		ConnectionGorm *gorm.DB

		Usecase *usecase.Usecase

		Orm *orm.Orm
	}
)

func NewFactory() (f *Factory, err error) {
	f = &Factory{}
	err = f.setupAdapterOutDrivenOrm()
	if err != nil {
		return
	}

	f.setupUseCase()

	return
}
func (f *Factory) setupAdapterOutDrivenOrm() (err error) {
	database.Init()

	gormConnection, err := database.Connection[gorm.DB]("postgres")

	f.Orm = orm.New(gormConnection)
	return nil
}

func (f *Factory) setupUseCase() {
	f.Usecase = usecase.New(
		f.Orm.User,
	)
}
