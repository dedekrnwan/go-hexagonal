package rest

import (
	"fmt"
	"go-boiler-clean/internal/adapter/inbound/rest/handler/auth"
	"go-boiler-clean/internal/adapter/inbound/rest/handler/user"
	"go-boiler-clean/internal/adapter/inbound/rest/middlewares"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (r *rest) InitRoute() {
	r.e.GET("/", func(c echo.Context) error {
		message := fmt.Sprintf("Welcome to %s version %s", "go-boiler-clean", "0.0.1")
		return c.String(http.StatusOK, message)
	})

	middlewares.Init(r.e)

	r.initV1Route(r.e.Group("/v1"))
}

func (r *rest) initV1Route(g *echo.Group) {
	user.NewHandler(
		r.usecaseUser,
	).Route(g.Group("/users"))

	auth.NewHandler(
		r.usecaseAuth,
	).Route(g.Group("/auth"))
}
