package connections

import (
	"fmt"
	"go-boilerplate-v2/src/models"
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewMySQLConnection() (mySql *gorm.DB, err error) {
	config := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("MSYQL_DB_USER"),
		os.Getenv("MSYQL_DB_PASSWORD"),
		os.Getenv("MSYQL_DB_HOST"),
		os.Getenv("MSYQL_DB_PORT"),
		os.Getenv("MSYQL_DB_NAME"),
	)

	var newLog logger.Interface

	if os.Getenv("ENVIRONMENT") != "production" {
		newLog = logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
			logger.Config{
				SlowThreshold:             time.Second, // Slow SQL threshold
				LogLevel:                  logger.Info, // Log level
				IgnoreRecordNotFoundError: false,       // Ignore ErrRecordNotFound error for logger
				Colorful:                  true,        // Disable color
			},
		)
	}

	mySql, err = gorm.Open(postgres.Open(config), &gorm.Config{
		Logger: newLog,
	})
	if err != nil {
		return
	}

	fmt.Println("Success connected to database")

	if err = mySql.Debug().AutoMigrate(models.User{}); err != nil {
		return
	}

	return
}
