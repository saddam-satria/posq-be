package commons

import (
	"github.com/saddam-satria/posq-be/domains"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DatabaseConnection *gorm.DB


func DatabaseConnect(ctx domains.DatabaseConnection) error {
	dsn := "host=" + ctx.HOST + " user=" + ctx.USER + " password=" + ctx.PASSWORD + " dbname=" + ctx.DATABASE +  " sslmode=require"
	db, error := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if error != nil {
		return error
	}

	DatabaseConnection = db
	return error
}