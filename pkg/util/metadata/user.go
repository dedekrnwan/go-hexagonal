package metadata

import (
	"context"
	"fmt"
	"go-boiler-clean/entity"
)

type (
	User interface {
	}
	user struct{}
)

func NewUser() User {
	return &user{}
}

func (m *user) GetUser(ctx context.Context) (*entity.User, error) {
	data, ok := ctx.Value(metadataUser).(*entity.User)
	if !ok {
		return nil, fmt.Errorf("error getting metadata user from context")
	}
	return data, nil
}

func (m *user) SetUser(ctx context.Context, data *entity.User) context.Context {
	return context.WithValue(ctx, metadataUser, data)
}
