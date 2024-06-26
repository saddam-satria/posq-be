package repositories

import (
	"fmt"
	"strings"

	"github.com/saddam-satria/posq-be/commons"
	"github.com/saddam-satria/posq-be/domains/query"
	"github.com/saddam-satria/posq-be/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func GetProducts(filter string, order string, sort string, keyword string, result *[]models.Product) {
	if filter == "" {
		filter = "product"
	}

	if order == "" {
		order = "asc"
	}

	if sort == "" {
		sort = "name"
	}

	query := commons.DatabaseConnection.Model(&models.Product{}).Where(&models.Product{Category: filter})

	if keyword != "" {
		query.Where("LOWER(name) LIKE ?", "%"+strings.ToLower(keyword)+"%")
	}

	query.Order(fmt.Sprintf("name %s", order)).Preload("ProductVariants", func(db *gorm.DB) *gorm.DB {
		return db.Order(clause.OrderByColumn{Column: clause.Column{Name: "basePrice"}, Desc: true})
	}).Find(&result)

}

func GetVariants(itemId string, result *[]models.ProductVariant) {
	commons.DatabaseConnection.Model(&models.ProductVariant{}).Order("name asc").Where(&models.ProductVariant{ProductId: itemId}).Find(&result)
}

func CountProductOutStock(product *query.CountAggregate) {
	commons.DatabaseConnection.Model(&models.ProductVariant{}).Select("count(*) AS total").Where("stock < 10").Find(&product)
}

func CountUnavailableService(product *query.CountAggregate) {
	commons.DatabaseConnection.Model(&models.Product{}).Select("count(*) as total").Where("category = service").Where("IsAvailable = false").Find(&product)
}
