package factory

import (
	adapter "go-boiler-clean/internal/adapter"
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

func NewFactory() (f *Factory, err error) {
	f = &Factory{}
	err = f.setupAdapterOutBound()
	if err != nil {
		return
	}

	f.setupUseCase()
	f.setupAdapterInBound()

	return
}
func (f *Factory) setupAdapterOutBound() (err error) {
	f.Adapter.OutBound, err = adapter.NewOutBound()
	return
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
