package main

import (
	"fmt"

	"github.com/saddam-satria/posq-be/commons"
	"github.com/saddam-satria/posq-be/models"
	"gorm.io/gorm"
)

func Seed(db *gorm.DB) {
	password, _ :=commons.HashedBcrypt("admin")
	userCredential:= models.UserCredential{
		Username: "admin",
		Password: password,
		UserProfile: &models.UserProfile{
			Name: "admin",
			Email: "admin@gmail.com",
			PhoneNumber: "089920001000",
		},
	}

	db.Create(&userCredential)

	fmt.Println("seed success")
}