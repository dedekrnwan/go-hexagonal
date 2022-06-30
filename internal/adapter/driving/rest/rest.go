package rest

import (
	"context"
	"go-boiler-clean/internal/factory"

	"github.com/labstack/echo/v4"
)

type (
	Rest interface {
		PrepareEcho() (func() error, func(ctx context.Context) error)
	}
	rest struct {
		e *echo.Echo
		f *factory.Factory
	}
)

func New(f *factory.Factory) Rest {
	e := echo.New()
	return &rest{
		e,
		f,
	}
}

func (h *rest) PrepareEcho() (func() error, func(ctx context.Context) error) {
	h.InitRoute(h.f)

	return func() error {
			return h.e.Start(":" + "8081")
		}, func(ctx context.Context) error {
			return h.e.Shutdown(ctx)
		}
}
