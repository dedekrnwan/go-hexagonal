package config

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/jinzhu/configor"
	"github.com/sirupsen/logrus"
)

type Configuration struct {
	Server    Server     `env:"server"`
	Swagger   Swagger    `env:"swagger"`
	JWT       JWT        `env:"jwt"`
	Databases []Database `env:"databases"`
}

type Server struct {
	AppName string `env:"app_name"`
	AppKey  string `env:"app_key"`
	Port    string `env:"port"`
	Version string `env:"version"`
}

type Swagger struct {
	SwaggerScheme string `env:"swagger_scheme"`
	SwaggerPrefix string `env:"swagger_prefix"`
}

type JWT struct {
	Secret string `env:"secret"`
}

type Database struct {
	DBHost        string `env:"db_host"`
	DBUser        string `env:"db_user"`
	DBPass        string `env:"db_pass"`
	DBPort        string `env:"db_port"`
	DBName        string `env:"db_name"`
	DBProvider    string `env:"db_provider"`
	DBSSL         string `env:"db_ssl"`
	DBTZ          string `env:"db_tz"`
	DBAutomigrate bool   `env:"db_automigrate"`
}

var Config *Configuration = &Configuration{}

func Load(path string) error {
	if path == "" {
		path = "config.yml"
	}
	if os.Getenv("ENV") == "debug" {
		path = "../../config.yml"
	}
	if os.Getenv("ENV") != "development" && os.Getenv("ENV") != "debug" {
		path = fmt.Sprintf("/run/secrets/%s", os.Getenv("CONFIG"))
	}
	err := configor.New(&configor.Config{AutoReload: true, AutoReloadInterval: time.Minute}).Load(Config, path)
	if err != nil {
		logrus.Info(err)
		return err
	}

	return nil
}

func (Configuration) String() string {
	sb := strings.Builder{}
	return sb.String()
}

func (c *Configuration) Raw() string {
	bytes, err := json.Marshal(c)
	if err != nil {
		return err.Error()
	}
	return string(bytes)
}
