package seeder

import (
	"context"
	"go-boiler-clean/internal/factory"
)

type (
	Seeder interface {
		Seed(context.Context, int) error
	}

	seeder struct {
		User
	}
)

func NewSeeder(f *factory.Factory) Seeder {
	return &seeder{
		User: NewUser(f),
	}
}

func (s *seeder) Seed(ctx context.Context, rows int) error {
	// s.User.Seed(ctx, rows)
	return nil
}
