package database

import (
	"context"
	"fmt"
	"reflect"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/bsontype"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Db interface {
	Init() (interface{}, error)
	// InitGorm() (*gorm.DB, error)
	// InitMongo() (*mongo.Client, error)
}

type db struct {
	Host string
	User string
	Pass string
	Port string
	Name string
}

type dbPostgreSQL struct {
	db
	SslMode     string
	Tz          string
	AutoMigrate bool
}

type dbMongo struct {
	db
}

func (c *dbPostgreSQL) Init() (interface{}, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s", c.Host, c.User, c.Pass, c.Name, c.Port, c.SslMode, c.Tz)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Error),
	})
	if err != nil {
		return nil, err
	}

	if c.AutoMigrate {
		logrus.Info("[*] Gorm auto migration from entities")
		db.AutoMigrate(Entity...)
	}
	return db, nil
}

func (c *dbMongo) Init() (interface{}, error) {
	rb := bson.NewRegistryBuilder()
	rb.RegisterTypeMapEntry(bsontype.EmbeddedDocument, reflect.TypeOf(bson.M{}))

	clientOptions := options.Client().
		ApplyURI(c.Host).
		SetAuth(options.Credential{
			Username:   c.User,
			Password:   c.Pass,
			AuthSource: c.Name,
		}).
		SetRegistry(rb.Build())
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return nil, err
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		return nil, err
	}
	db := client.Database(c.Name)
	return db, nil
}
