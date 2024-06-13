package repositories

import (
	"database/sql"
	"errors"

	"github.com/saddam-satria/posq-be/commons"
	"github.com/saddam-satria/posq-be/domains/apis"
	"github.com/saddam-satria/posq-be/models"
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
		totalAmount += variant.BasePrice * float32(order.Products[index].Quantity)
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

	db.Commit()
	return nil
}
