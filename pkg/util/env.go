package util

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

type Env interface {
	GetString(name string) string
}

type env struct {
	Env
}

type EnvGetter struct{}

func NewEnv() *env {
	return &env{Env: &EnvGetter{}}
}

func (e *env) Load() {

	err := godotenv.Load()

	if err != nil {
		logrus.WithFields(logrus.Fields{
			"cause": err,
		}).Fatal("Load .env file error")

		os.Exit(-1)
	}
}

func (r *EnvGetter) GetString(name string) string {
	return os.Getenv(name)
}

func (e *env) GetString(name string) string {
	if nil == e.Env {
		return ""
	}
	return e.Env.GetString(name)
}

func (e *env) GetBool(name string) bool {
	s := e.GetString(name)
	i, err := strconv.ParseBool(s)
	if err != nil {
		return false
	}
	return i
}

func (e *env) GetInt(name string) int {
	s := e.GetString(name)
	i, err := strconv.Atoi(s)
	if err != nil {
		return 0
	}
	return i
}

func (e *env) GetFloat(name string) float64 {
	s := e.GetString(name)
	i, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return 0
	}
	return i
}
