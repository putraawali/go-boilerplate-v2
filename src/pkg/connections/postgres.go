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

func NewPostgreConnection() (pg *gorm.DB, err error) {
	config := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("PG_DB_HOST"),
		os.Getenv("PG_DB_USER"),
		os.Getenv("PG_DB_PASSWORD"),
		os.Getenv("PG_DB_NAME"),
		os.Getenv("PG_DB_PORT"),
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

	pg, err = gorm.Open(postgres.Open(config), &gorm.Config{
		Logger: newLog,
	})
	if err != nil {
		return
	}

	fmt.Println("Success connected to database")

	if err = pg.Debug().AutoMigrate(models.User{}); err != nil {
		return
	}

	return
}
