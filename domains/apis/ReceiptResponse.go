package apis

import (
	"time"

	"github.com/google/uuid"
)

type ReceiptResponse struct {
	OrderId     uuid.UUID  `json:"id" gorm:"column:orderId;primaryKey"`
	ReferenceId string     `json:"reference_id" gorm:"column:referenceId"`
	OrderNote   *string    `json:"payment_note" gorm:"column:orderNote"`
	PaymentType string     `json:"payment_type" gorm:"column:paymentType"`
	Tendered    float32    `json:"tendered" gorm:"column:tendered"`
	CreatedAt   *time.Time `json:"-" gorm:"column:createdAt"`

	CardNumber *string `json:"card_number" gorm:"column:cardNumber"`
	Address    string  `json:"address" gorm:"_"`
	Logo       string  `json:"logo" gorm:"-"`
	TotalItem  int     `json:"total_item" gorm:"-"`

	CustomerId       *string `json:"customer_id" gorm:"column:customer_id"`
	UserCredentialId string  `json:"user_id" gorm:"column:user_credential_id"`

	SubTotal       float32        `json:"total" gorm:"-"`
	Changed        float32        `json:"changed" gorm:"-"`
	OrderTimestamp int64          `json:"order_timestamp" gorm:"-"`
	Products       []OrderProduct `json:"items" gorm:"foreignKey:OrderId"`
}

func (o *ReceiptResponse) TableName() string {
	return "order"
}
