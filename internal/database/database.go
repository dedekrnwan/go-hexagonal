package database

import (
	"errors"
	"fmt"
	"strings"

	"go-boiler-clean/internal/config"

	"github.com/sirupsen/logrus"
)

var (
	dbConnections map[string]interface{}
)

func Init() {
	dbConfigurations := map[string]Db{}
	for _, v := range config.Config.Databases {
		switch v.DBProvider {
		case "postgresql":
			dbConfigurations[strings.ToUpper(v.DBName)] = &dbPostgreSQL{
				db: db{
					Host: v.DBHost,
					Name: v.DBName,
					Port: v.DBPort,
					Pass: v.DBPass,
					User: v.DBUser,
				},
				SslMode:     v.DBSSL,
				Tz:          v.DBTZ,
				AutoMigrate: v.DBAutomigrate,
			}
		case "mongodb":
			dbConfigurations[strings.ToUpper(v.DBName)] = &dbMongo{
				db: db{
					Host: v.DBHost,
					Name: v.DBName,
					Port: v.DBPort,
					Pass: v.DBPass,
					User: v.DBUser,
				},
			}
		default:
			dbConfigurations[strings.ToUpper(v.DBName)] = &dbPostgreSQL{
				db: db{
					Host: v.DBHost,
					Name: v.DBName,
					Port: v.DBPort,
					Pass: v.DBPass,
					User: v.DBUser,
				},
				SslMode: v.DBSSL,
				Tz:      v.DBTZ,
			}
		}
	}

	dbConnections = make(map[string]interface{})
	for k, v := range dbConfigurations {
		db, err := v.Init()
		if err != nil {
			logrus.Info(err)
			panic(fmt.Sprintf("Failed to connect to database %s", k))
		}
		dbConnections[k] = db
		logrus.Info(fmt.Sprintf("Successfully connected to database %s", k))
	}
}

func Connection[T any](name string) (*T, error) {
	conn := dbConnections[strings.ToUpper(name)]
	if conn == nil {
		return nil, errors.New("Connection is undefined")
	}

	connection, ok := conn.(*T)
	if !ok {
		return nil, errors.New("error parsing connection to generics type")
	}

	return connection, nil
}
