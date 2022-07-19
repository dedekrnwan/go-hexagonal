package factory

import (
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

	//orm
	f.Orm, err = orm.New()
	if err != nil {
		return
	}

	//usecase
	f.setupUseCase()

	return
}

func (f *Factory) setupUseCase() {
	f.Usecase = usecase.New(
		f.Orm.User,
	)
}
