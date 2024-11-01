package mock_connections

import (
	"database/sql"

	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewMockPostgresConnection() (conn *gorm.DB, mock sqlmock.Sqlmock) {
	var db *sql.DB
	db, mock, _ = sqlmock.New()
	// db, mock, _ = sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherRegexp))

	var err error
	conn, err = gorm.Open(postgres.New(postgres.Config{
		// DSN:                  "sqlmock_db_0",
		Conn: db,
		// PreferSimpleProtocol: true,
		// DriverName:           "postgres",
	}), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Error),
	})

	if err != nil {
		panic(err)
	}

	return
}
