package seeder

import (
	"context"
	"go-boiler-clean/entity"
	"go-boiler-clean/internal/factory"
	modelPsqlGorm "go-boiler-clean/internal/model/psqlGorm"

	"github.com/bxcodec/faker/v3"
)

type (
	User interface {
		Seed(context.Context, int) error
	}

	user struct {
		model modelPsqlGorm.User
	}
)

func NewUser(f *factory.Factory) User {
	return &user{
		model: f.Model.ModelPsqlGormUser,
	}
}

func (s *user) Seed(ctx context.Context, rows int) error {
	data := []entity.User{}
	ch := make(chan entity.User)
	active := true
	go func(ch chan entity.User) {
		for i := 0; i < rows; i++ {
			ch <- entity.User{
				FirstName: faker.FirstName(),
				LastName:  faker.LastName(),
				Email:     faker.Email(),
				Phone:     faker.Phonenumber(),
				IsActive:  &active,
				Password:  faker.Password(),
			}
		}
	}(ch)
	for i := 0; i < rows; i++ {
		data = append(data, <-ch)
	}
	_, err := s.model.CreateMany(ctx, data)
	close(ch)
	return err
}
