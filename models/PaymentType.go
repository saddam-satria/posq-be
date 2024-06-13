package models

import "database/sql/driver"

type PaymentType string

const (
	CASH PaymentType = "cash"
	DEBT PaymentType = "debit"
)

func (ct *PaymentType) Scan(value interface{}) error {
	*ct = PaymentType(value.([]byte))
	return nil
}

func (ct PaymentType) Value() (driver.Value, error) {
	return string(ct), nil
}
