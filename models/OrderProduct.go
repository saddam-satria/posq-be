package models

import "github.com/google/uuid"

type OrderProduct struct {
	OrderProductId   uuid.UUID `json:"id" gorm:"column:orderProductId;type:uuid;primaryKey;default:gen_random_uuid()"`
	Quantity         int       `json:"quantity" gorm:"column:quantity;type:integer"`
	Description      *string   `json:"note" gorm:"column:description;type:text"`
	OrderId          string    `json:"order_id" gorm:"column:order_id;type:uuid"`
	ProductVariantId string    `json:"variant_id" gorm:"column:product_variant_id;type:uuid"`
}

func (o *OrderProduct) TableName() string {
	return "orderProduct"
}
