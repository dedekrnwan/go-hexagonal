package main

import (
	"go-boiler-clean/internal/adapter/inbound/http_echo"
	"go-boiler-clean/internal/config"
	"go-boiler-clean/internal/factory"
	"go-boiler-clean/pkg/util"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

func init() {
	if os.Getenv("ENV") == "" {
		env := util.NewEnv()
		env.Load()
	}

	err := config.Load("")
	if err != nil {
		logrus.Error(err)
		os.Exit(1)
	}
}

func main() {

	//initialize all needs (db conn, adapter outbound)
	f := initFactory()

	initHttp(f)
}

func initFactory() *factory.Factory {
	return factory.NewFactory()
}

func initHttp(f *factory.Factory) {
	e := echo.New()
	http_echo.RunEcho(e, f)
}
