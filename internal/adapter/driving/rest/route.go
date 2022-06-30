package rest

import (
	"fmt"
	"go-boiler-clean/internal/adapter/driving/rest/handler/user"
	"go-boiler-clean/internal/adapter/driving/rest/middlewares"
	"go-boiler-clean/internal/factory"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (r *rest) InitRoute(f *factory.Factory) {
	r.e.GET("/", func(c echo.Context) error {
		message := fmt.Sprintf("Welcome to %s version %s", "go-boiler-clean", "0.0.1")
		return c.String(http.StatusOK, message)
	})

	middlewares.Init(r.e)

	r.initV1Route(r.e.Group("/v1"), f)
}

func (r *rest) initV1Route(g *echo.Group, f *factory.Factory) {
	user.NewHandler(f).Route(g.Group("/users"))
}
