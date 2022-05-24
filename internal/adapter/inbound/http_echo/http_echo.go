package http_echo

import (
	"go-boiler-clean/internal/factory"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

func RunEcho(e *echo.Echo, f *factory.Factory) {
	InitEchoRoute(e, f)

	// e.Logger.Fatal(e.Start(":" + "8081"))
	logrus.Info(e.Start(":" + "8081"))
}
