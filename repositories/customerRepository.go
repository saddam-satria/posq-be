package repositories

import (
	"github.com/saddam-satria/posq-be/commons"
	"github.com/saddam-satria/posq-be/models"
	"gorm.io/gorm"
)

func GetCustomers(result *[]models.Customer) {
	query := commons.DatabaseConnection.Model(&models.Customer{})

	query.Find(&result)
}

func CreateCustomer(customer *models.Customer, result *models.Customer) *gorm.DB {
	return commons.DatabaseConnection.Create(&customer).Scan(&result)
}
