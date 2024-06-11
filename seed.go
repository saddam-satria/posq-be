package main

import (
	"fmt"

	"github.com/saddam-satria/posq-be/commons"
	"github.com/saddam-satria/posq-be/models"
	"gorm.io/gorm"
)

func Seed(db *gorm.DB) {
	password, _ := commons.HashedBcrypt("admin")
	userCredential := models.UserCredential{
		Username: "admin",
		Password: password,
		UserProfile: &models.UserProfile{
			Name:        "admin",
			Email:       "admin@gmail.com",
			PhoneNumber: "089920001000",
		},
	}

	product := models.Product{
		Name:     "Racing motor oli",
		Category: "product",
		ProductVariants: []models.ProductVariant{{
			Name:      "5 liter",
			BasePrice: 40000,
			SalePrice: 10000,
			Brand:     "yamalube",
			Stock:     20,
			Sku:       "HJG",
		},
			{
				Name:      "2 liter",
				BasePrice: 30000,
				SalePrice: 10000,
				Brand:     "yamalube",
				Stock:     20,
				Sku:       "KKJ",
			},
		},
	}

	db.Create(&userCredential)
	db.Create(&product)

	fmt.Println("seed success")
}
