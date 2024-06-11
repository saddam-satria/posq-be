package main

import (
	"fmt"

	"github.com/saddam-satria/posq-be/models"
	"gorm.io/gorm"
)

func migrate(db *gorm.DB) {
	if err := db.AutoMigrate(&models.UserCredential{}, &models.UserProfile{}); err != nil {
		panic("Migration Failed Reason: " + err.Error())
	}

	fmt.Println("Migrate Success")
}