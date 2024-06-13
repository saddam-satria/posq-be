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
	file, err := os.OpenFile("./logs/query.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		os.Create("./logs/query.log")
	}

	var writter = file

	if GoDotEnv("DEBUG") == "true" {
		writter = os.Stdout
	}

	newLogger := logger.New(
		log.New(writter, "\r\n", log.LstdFlags),
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
