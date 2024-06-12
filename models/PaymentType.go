package models

import "database/sql/driver"

type PaymentType string

const (
	CASH PaymentType = "CASH"
	DEBT PaymentType = "DEBT"
)

func (ct *PaymentType) Scan(value interface{}) error {
	*ct = PaymentType(value.([]byte))
	return nil
}

func (ct PaymentType) Value() (driver.Value, error) {
	return string(ct), nil
}
