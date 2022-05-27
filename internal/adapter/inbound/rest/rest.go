package rest

import (
	"context"
	"go-boiler-clean/internal/usecase"

	"github.com/labstack/echo/v4"
)

type (
	Rest interface {
		PrepareEcho() (func() error, func(ctx context.Context) error)
	}
	rest struct {
		usecaseUser usecase.User
		e           *echo.Echo
	}
)

func New(usecaseUser usecase.User) Rest {
	e := echo.New()
	return &rest{
		usecaseUser,
		e,
	}
}

func (h *rest) PrepareEcho() (func() error, func(ctx context.Context) error) {
	h.InitRoute()

	return func() error {
			return h.e.Start(":" + "8081")
		}, func(ctx context.Context) error {
			return h.e.Shutdown(ctx)
		}
}
