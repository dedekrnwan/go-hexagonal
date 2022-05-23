package main

import (
	"context"
	"go-boiler-clean/database"
	"go-boiler-clean/database/seeder"
	"go-boiler-clean/internal/adapters/http_echo"
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
	ctx := context.Background()

	initDatabase()
	f := initFactory()

	//seder
	seeder := seeder.NewSeeder(f)
	err := seeder.Seed(ctx, 100)
	if err != nil {
		ctx.Done()
		os.Exit(1)
	}
	initHttp(f)
}

func initDatabase() {
	database.Init()
}

func initFactory() *factory.Factory {
	return factory.NewFactory()
}

func initHttp(f *factory.Factory) {
	e := echo.New()
	http_echo.RunEcho(e, f)
}
