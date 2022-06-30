package repository

import (
	"context"
	"database/sql"
	"go-boiler-clean/dto"
	"go-boiler-clean/internal/adapter/driven/orm/entity"
	"reflect"
	"regexp"
	"strings"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/bxcodec/faker"
	"github.com/go-test/deep"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type (
	UserSuite struct {
		suite.Suite

		mock  sqlmock.Sqlmock
		model User
	}
)

func (s *UserSuite) SetupTest() {
	var (
		db  *sql.DB
		err error
	)

	db, s.mock, err = sqlmock.New()
	require.NoError(s.T(), err)

	dialector := postgres.New(postgres.Config{
		DriverName:           "postgres",
		DSN:                  "sqlmock_db_0",
		Conn:                 db,
		PreferSimpleProtocol: true,
	})

	gdb, err := gorm.Open(dialector, &gorm.Config{})
	require.NoError(s.T(), err)

	s.model = NewUser(gdb)
}

func (s *UserSuite) TestFindOne() {
	isActive := true
	userOutput := entity.User{}
	user := entity.User{}
	user.ID = 1
	user.FirstName = faker.FirstName
	user.LastName = faker.LastName
	user.Email = faker.Email
	user.Phone = faker.PhoneNumber
	user.IsActive = &isActive
	user.Password = faker.PASSWORD
	reflectUser := reflect.ValueOf(userOutput)
	cols := []string{}
	cols = append(cols, "id")
	for i := 0; i < reflectUser.NumField(); i++ {
		if !strings.Contains(reflectUser.Type().Field(i).Tag.Get("json"), "omitempty") && reflectUser.Type().Field(i).Tag.Get("json") != "" {
			cols = append(cols, reflectUser.Type().Field(i).Tag.Get("json"))
		}
	}
	s.mock.ExpectQuery(
		regexp.QuoteMeta(
			`SELECT * FROM "users" WHERE "id" = $1 ORDER BY "users"."id" LIMIT 1`),
	).
		WithArgs(user.ID).
		WillReturnRows(
			sqlmock.NewRows(cols).
				AddRow(user.ID, user.FirstName, user.LastName, user.Email, user.Phone, user.IsActive, user.Password),
		)

	res, err := s.model.FindOne(context.TODO(), user.ID, nil)

	require.NoError(s.T(), err)
	require.Nil(s.T(), deep.Equal(&user, res))
}

func (s *UserSuite) TestFind() {
	isActive := true
	userOutput := entity.User{}
	user := entity.User{}
	user.ID = 1
	user.FirstName = faker.FirstName
	user.LastName = faker.LastName
	user.Email = faker.Email
	user.Phone = faker.PhoneNumber
	user.IsActive = &isActive
	user.Password = faker.PASSWORD

	users := []entity.User{}
	users = append(users, user)

	infoPagination := dto.PaginationInfo{
		Pagination: dto.Pagination{},
		Count:      0,
		TotalPage:  0,
	}

	reflectUser := reflect.ValueOf(userOutput)

	cols := []string{}
	cols = append(cols, "id")

	for i := 0; i < reflectUser.NumField(); i++ {
		if !strings.Contains(reflectUser.Type().Field(i).Tag.Get("json"), "omitempty") && reflectUser.Type().Field(i).Tag.Get("json") != "" {
			cols = append(cols, reflectUser.Type().Field(i).Tag.Get("json"))
		}
	}

	s.mock.ExpectQuery(
		regexp.QuoteMeta(
			`SELECT * FROM "users"`),
	).
		WillReturnRows(
			sqlmock.NewRows(cols).
				AddRow(user.ID, user.FirstName, user.LastName, user.Email, user.Phone, user.IsActive, user.Password),
		)

	res, info, err := s.model.Find(context.TODO(), "", []dto.Filter{}, []string{}, []string{}, dto.Pagination{}, nil)

	require.NoError(s.T(), err)
	require.Nil(s.T(), deep.Equal(users, res))
	require.Nil(s.T(), deep.Equal(&infoPagination, info))
}

func (s *UserSuite) TestCreateOne() {
	isActive := true
	user := entity.User{}
	user.ID = 1
	user.FirstName = faker.FirstName
	user.LastName = faker.LastName
	user.Email = faker.Email
	user.Phone = faker.PhoneNumber
	user.IsActive = &isActive
	user.Password = faker.PASSWORD
	user.CreatedAt = time.Now()

	s.mock.ExpectBegin()
	s.mock.ExpectQuery(
		regexp.QuoteMeta(
			`INSERT INTO "users" ("created_at","created_by","modified_at","modified_by","deleted_at","deleted_by","first_name","last_name","email","phone","is_active","password","id") VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13) RETURNING "id"`),
	).
		WithArgs(user.CreatedAt, nil, time.Time{}, nil, time.Time{}, nil, user.FirstName, user.LastName, user.Email, user.Phone, user.IsActive, user.Password, user.ID).
		WillReturnRows(
			sqlmock.NewRows([]string{"id"}).
				AddRow(user.ID),
		)
	s.mock.ExpectCommit()

	// res, err := s.model.CreateOne(context.TODO(), &user)

	// require.NoError(s.T(), err)
	// require.Nil(s.T(), deep.Equal(&user, res))
}

func TestUserSuite(t *testing.T) {
	suite.Run(t, new(UserSuite))
}
