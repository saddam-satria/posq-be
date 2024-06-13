package models

import (
	"time"

	"github.com/google/uuid"
)

type Order struct {
	OrderId     uuid.UUID   `json:"id" gorm:"column:orderId;type:uuid;primaryKey;default:gen_random_uuid()"`
	ReferenceId string      `json:"reference_id" gorm:"column:referenceId;unique;type:varchar(50)"`
	OrderNote   *string     `json:"payment_note" gorm:"column:orderNote;type:text"`
	PaymentType PaymentType `json:"payment_type" gorm:"type:payment_type;column:paymentType"`
	Tendered    float32     `json:"tendered" gorm:"column:tendered;type:decimal(10,2)"`
	CreatedAt   *time.Time  `json:"order_timestamp" gorm:"default_now();column:createdAt"`
	UpdatedAt   *time.Time  `json:"-" gorm:"autoUpdateTime:milli;column:updatedAt"`
	CardNumber  *string     `json:"card_number" gorm:"column:cardNumber;type:varchar(255)"`

	CustomerId       *string `json:"customer_id" gorm:"column:customer_id;type:uuid"`
	UserCredentialId string  `json:"user_id" gorm:"column:user_credential_id;type:uuid"`

	Products        []OrderProduct   `json:"items" gorm:"foreignKey:OrderId"`
	ProductVariants []ProductVariant `gorm:"many2many:product_variants;joinForeignKey:order_id;joinTableForeignKey:product_variant_id"`
}

func (o *Order) TableName() string {
	return "order"
}
