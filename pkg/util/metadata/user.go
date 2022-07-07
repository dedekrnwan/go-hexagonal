package metadata

import (
	"context"
	"fmt"
	"go-boiler-clean/internal/model/sample"
)

type (
	User interface {
	}
	user struct{}
)

func NewUser() User {
	return &user{}
}

func (m *user) GetUser(ctx context.Context) (*sample.UserEntity, error) {
	data, ok := ctx.Value(metadataUser).(*sample.UserEntity)
	if !ok {
		return nil, fmt.Errorf("error getting metadata user from context")
	}
	return data, nil
}

func (m *user) SetUser(ctx context.Context, data *sample.UserEntity) context.Context {
	return context.WithValue(ctx, metadataUser, data)
}
