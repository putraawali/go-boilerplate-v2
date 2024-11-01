package repositories_test

import (
	"context"
	"errors"
	src_mock "go-boilerplate-v2/src/mocks"
	"go-boilerplate-v2/src/models"
	mock_connections "go-boilerplate-v2/src/pkg/connections/mocks"
	"go-boilerplate-v2/src/repositories"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
)

type userRepositoryTest struct {
	suite.Suite
	repoPostgres repositories.UserRepository
	repoMysql    repositories.UserRepository
	mysql        *gorm.DB
	postgres     *gorm.DB
	mockMysql    sqlmock.Sqlmock
	mockPostgres sqlmock.Sqlmock
	ctx          context.Context
}

func TestUserRepository(t *testing.T) {
	suite.Run(t, new(userRepositoryTest))
}

// Before each test
func (u *userRepositoryTest) SetupTest() {
	u.mysql, u.mockMysql = mock_connections.NewMockMySQLConnection()
	u.postgres, u.mockPostgres = mock_connections.NewMockPostgresConnection()

	diPostgres := src_mock.NewMockDependencies(src_mock.Dependencies{
		Postgres: u.postgres,
	})

	diMysql := src_mock.NewMockDependencies(src_mock.Dependencies{
		Mysql: u.mysql,
	})

	u.repoPostgres = repositories.NewUserRepository(diPostgres)
	u.repoMysql = repositories.NewUserRepository(diMysql)

	u.ctx = context.WithValue(context.TODO(), "request-id", "213")
}

func (u *userRepositoryTest) TestInsertPostgres() {
	u.Run("Success insert user as postgres", func() {
		data := models.User{
			Email:     "email@mail.com",
			FirstName: "First Name",
			LastName:  "Last Name",
			Phone:     "08123123123",
			Password:  "pw-123",
		}

		query := "INSERT INTO \"users\" "

		u.mockPostgres.ExpectBegin()

		u.mockPostgres.ExpectQuery(query).
			WithArgs(
				data.Email,
				data.FirstName,
				data.LastName,
				data.Phone,
				sqlmock.AnyArg(),
				sqlmock.AnyArg(),
				sqlmock.AnyArg(),
			).
			WillReturnRows(sqlmock.NewRows([]string{"user_id"}).AddRow(1))

		u.mockPostgres.ExpectCommit()

		err := u.repoPostgres.Insert(u.ctx, &data)
		u.Nil(err, "error should be nil")
	})

	u.Run("Failed insert user as postgres", func() {
		data := models.User{
			Email:     "email@mail.com",
			FirstName: "First Name",
			LastName:  "Last Name",
			Phone:     "08123123123",
			Password:  "pw-123",
		}

		query := "INSERT INTO \"users\" "

		u.mockPostgres.ExpectBegin()

		u.mockPostgres.ExpectQuery(query).
			WithArgs(
				data.Email,
				data.FirstName,
				data.LastName,
				data.Phone,
				sqlmock.AnyArg(),
				sqlmock.AnyArg(),
				sqlmock.AnyArg(),
			).
			WillReturnError(errors.New("mock error"))

		u.mockPostgres.ExpectRollback()

		err := u.repoPostgres.Insert(u.ctx, &data)
		u.NotNil(err, "error should be not nil")
	})

	u.Run("Success insert user as mysql", func() {
		data := models.User{
			Email:     "email@mail.com",
			FirstName: "First Name",
			LastName:  "Last Name",
			Phone:     "08123123123",
			Password:  "pw-123",
		}

		query := "INSERT INTO `users` "

		u.mockMysql.ExpectBegin()

		u.mockMysql.ExpectExec(query).
			WithArgs(
				data.Email,
				data.FirstName,
				data.LastName,
				data.Phone,
				sqlmock.AnyArg(),
				sqlmock.AnyArg(),
				sqlmock.AnyArg(),
			).
			WillReturnResult(sqlmock.NewResult(1, 1))

		u.mockMysql.ExpectCommit()

		err := u.repoMysql.Insert(u.ctx, &data)
		u.Nil(err, "error should be nil")
	})

	u.Run("Failed insert user as mysql", func() {
		data := models.User{
			Email:     "email@mail.com",
			FirstName: "First Name",
			LastName:  "Last Name",
			Phone:     "08123123123",
			Password:  "pw-123",
		}

		query := "INSERT INTO `users` "

		u.mockMysql.ExpectBegin()

		u.mockMysql.ExpectExec(query).
			WithArgs(
				data.Email,
				data.FirstName,
				data.LastName,
				data.Phone,
				sqlmock.AnyArg(),
				sqlmock.AnyArg(),
				sqlmock.AnyArg(),
			).
			WillReturnError(errors.New("mock error"))

		u.mockMysql.ExpectRollback()

		err := u.repoMysql.Insert(u.ctx, &data)
		u.NotNil(err, "error should be not nil")
	})
}
