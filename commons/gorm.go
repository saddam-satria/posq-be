package commons

import (
	"log"
	"os"
	"time"

	"github.com/saddam-satria/posq-be/domains"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DatabaseConnection *gorm.DB

func DatabaseConnect(ctx domains.DatabaseConnection) error {

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             time.Second,
			LogLevel:                  logger.Info,
			IgnoreRecordNotFoundError: true,
			ParameterizedQueries:      true,
			Colorful:                  false,
		},
	)

	dsn := "host=" + ctx.HOST + " user=" + ctx.USER + " password=" + ctx.PASSWORD + " dbname=" + ctx.DATABASE + " sslmode=require"
	db, error := gorm.Open(postgres.Open(dsn), &gorm.Config{Logger: newLogger})

	if error != nil {
		return error
	}

	DatabaseConnection = db
	return error
}
