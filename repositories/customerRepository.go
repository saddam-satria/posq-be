package repositories

import (
	"github.com/google/uuid"
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

func FindCustomerById(customerId uuid.UUID, result *models.Customer) error {
	return commons.DatabaseConnection.Model(&models.Customer{}).Where(&models.Customer{CustomerId: customerId}).Find(&result).Error
}
