package apis

import (
	"time"

	"github.com/google/uuid"
)

type OrderProduct struct {
	OrderProductId   uuid.UUID `json:"id" gorm:"column:orderProductId;primaryKey"`
	Quantity         int       `json:"quantity" gorm:"column:quantity"`
	Description      *string   `json:"note" gorm:"column:description"`
	Name             string    `json:"variant_name" gorm:"column:name"`
	BasePrice        float32   `json:"base_price" gorm:"column:basePrice"`
	SalePrice        float32   `json:"price" gorm:"column:salePrice"`
	Brand            string    `json:"brand" gorm:"column:brand"`
	Stock            int       `json:"stock" gorm:"column:stock"`
	Sku              string    `json:"sku" gorm:"column:sku"`
	OrderId          string    `json:"order_id" gorm:"column:order_id"`
	ProductVariantId string    `json:"variant_id" gorm:"column:product_variant_id"`
	ProductName      string    `json:"item_name" gorm:"column:item_name"`
}

type OrderResponse struct {
	OrderId     uuid.UUID  `json:"id" gorm:"column:orderId;primaryKey"`
	ReferenceId string     `json:"reference_id" gorm:"column:referenceId"`
	OrderNote   *string    `json:"payment_note" gorm:"column:orderNote"`
	PaymentType string     `json:"payment_type" gorm:"column:paymentType"`
	Tendered    float32    `json:"tendered" gorm:"column:tendered"`
	CreatedAt   *time.Time `json:"-" gorm:"column:createdAt"`

	CustomerId       *string `json:"customer_id" gorm:"column:customer_id"`
	UserCredentialId string  `json:"user_id" gorm:"column:user_credential_id"`

	SubTotal       float32        `json:"subtotal" gorm:"-"`
	Changed        float32        `json:"change" gorm:"-"`
	OrderTimestamp int64          `json:"order_timestamp" gorm:"-"`
	Products       []OrderProduct `json:"items" gorm:"foreignKey:OrderId"`
}

func (o *OrderResponse) TableName() string {
	return "order"
}
func (o *OrderProduct) TableName() string {
	return "orderProduct"
}
