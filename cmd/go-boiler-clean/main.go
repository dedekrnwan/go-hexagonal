package main

import (
	"go-boiler-clean/internal/adapter/driving/rest"
	"go-boiler-clean/internal/config"
	"go-boiler-clean/internal/factory"
	"go-boiler-clean/pkg/util"
	"os"
	"sync"
	"time"

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
	f, err := factory.NewFactory()
	if err != nil {
		panic(err)
	}

	//rest
	restInstance := rest.New(f)
	starterEcho, stopperEcho := restInstance.PrepareEcho()

	//grpc

	wg := new(sync.WaitGroup)

	wg.Add(2)
	go func() {
		util.StartProcessAtBackground(starterEcho)
		util.GracefullStopProcessAtBackground(time.Second*10, stopperEcho)
		wg.Done()
	}()

	// //sample multiple process
	// go func() {
	// 	time.Sleep(time.Second * 20)
	// 	logrus.Info("testing log")
	// 	wg.Done()
	// }()
	wg.Wait()
}
