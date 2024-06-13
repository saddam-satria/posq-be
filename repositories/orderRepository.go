package repositories

import (
	"database/sql"
	"errors"

	"github.com/saddam-satria/posq-be/commons"
	"github.com/saddam-satria/posq-be/domains/apis"
	"github.com/saddam-satria/posq-be/domains/query"
	"github.com/saddam-satria/posq-be/models"
	"gorm.io/gorm"
)

func Checkout(order *models.Order, response *apis.CheckoutResponse) error {
	db := commons.DatabaseConnection.Begin(&sql.TxOptions{})

	productVariantIds := []string{}

	for _, item := range order.Products {
		productVariantIds = append(productVariantIds, item.ProductVariantId)
	}

	var productVariants []models.ProductVariant

	if err := commons.DatabaseConnection.Model(&models.ProductVariant{}).Where("\"productVariantId\" IN ?", productVariantIds).Find(&productVariants).Error; err != nil {
		db.Rollback()
		return err
	}

	if len(productVariants) != len(order.Products) {
		return errors.New("variant not found")
	}

	totalAmount := float32(0)

	for index, variant := range productVariants {
		for _, requestVariant := range order.Products {
			if requestVariant.ProductVariantId == variant.ProductVariantId.String() {
				if requestVariant.Quantity > variant.Stock {
					db.Rollback()
					return errors.New("run out of stock")
				}
				currentStock := variant.Stock - requestVariant.Quantity
				variant.Stock = currentStock
				db.Save(variant)
			}
		}
		totalAmount += variant.SalePrice * float32(order.Products[index].Quantity)
	}

	if totalAmount > order.Tendered {
		db.Rollback()
		return errors.New("less payment")
	}

	query := db.Create(&order)

	if err := query.Error; err != nil {
		db.Rollback()
		return err
	}

	response.Changed = order.Tendered - totalAmount
	response.ReferenceId = order.ReferenceId
	response.Tendered = order.Tendered
	response.Total = totalAmount

	return db.Commit().Error

}

func GetOrders(order *[]apis.OrderResponse) {
	commons.DatabaseConnection.Preload("Products", func(tx *gorm.DB) *gorm.DB {
		return tx.Joins("JOIN \"productVariant\" AS pv ON pv.\"productVariantId\" = \"orderProduct\".product_variant_id").Joins("JOIN product ON product.\"productId\" = pv.product_id").Select(
			"\"orderProduct\".*", "pv.name", "pv.\"basePrice\"", "pv.\"salePrice\"", "pv.brand", "pv.stock", "pv.sku", "pv.\"productVariantId\"", "pv.product_id", "product.name AS item_name",
		)
	}).Order("\"createdAt\" DESC").Find(&order)
}

func GetOrderReceipt(referenceId string, order *apis.ReceiptResponse) error {
	return commons.DatabaseConnection.Preload("Products", func(tx *gorm.DB) *gorm.DB {
		return tx.Joins("JOIN \"productVariant\" AS pv ON pv.\"productVariantId\" = \"orderProduct\".product_variant_id").Joins("JOIN product ON product.\"productId\" = pv.product_id").Select(
			"\"orderProduct\".*", "pv.name", "pv.\"basePrice\"", "pv.\"salePrice\"", "pv.brand", "pv.stock", "pv.sku", "pv.\"productVariantId\"", "pv.product_id", "product.name AS item_name",
		)
	}).Where(&models.Order{ReferenceId: referenceId}).First(&order).Error
}

func CountTotalOrderByToday(order *query.CountAggregate) {
	commons.DatabaseConnection.Model(&models.Order{}).Select("count(*) as total").Where("CAST(\"createdAt\" AS DATE) = CURRENT_DATE").Find(&order)
}
