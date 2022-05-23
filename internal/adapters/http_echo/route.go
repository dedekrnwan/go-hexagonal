package http_echo

import (
	"fmt"
	"go-boiler-clean/internal/adapters/http_echo/handler/transaction"
	"go-boiler-clean/internal/adapters/http_echo/handler/user"
	"go-boiler-clean/internal/factory"
	"net/http"

	"github.com/labstack/echo/v4"
)

func InitEchoRoute(
	e *echo.Echo,
	f *factory.Factory,
) {
	e.GET("/", func(c echo.Context) error {
		message := fmt.Sprintf("Welcome to %s version %s", "go-boiler-clean", "0.0.1")
		return c.String(http.StatusOK, message)
	})

	InitV1Route(e.Group("/v1"), f)
}

func InitV1Route(e *echo.Group, f *factory.Factory) {
	transaction.NewHandler().Route(e.Group("/transactions"))
	user.NewHandler(f).Route(e.Group("/users"))
}
