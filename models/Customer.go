package models

import (
	"time"

	"github.com/google/uuid"
)

type Customer struct {
	CustomerId  uuid.UUID  `json:"id" gorm:"column:customerId;type:uuid;primaryKey;default:gen_random_uuid()"`
	Name        string     `json:"name" gorm:"column:name;type:varchar(255)"`
	Email       string     `json:"email" gorm:"column:email;type:varchar(255);unique"`
	Phonenumber string     `json:"phone" gorm:"column:phonenumber;type:varchar(15);unique"`
	Address     *string    `json:"address" gorm:"column:address;type:text"`
	CreatedAt   *time.Time `json:"-" gorm:"default_now();column:createdAt"`
	UpdatedAt   *time.Time `json:"-" gorm:"autoUpdateTime:milli;column:updatedAt"`
	IsDeleted   bool       `json:"-" gorm:"default:false;column:isDeleted"`

	Orders *[]Order `json:"orders" gorm:"foreignKey:CustomerId"`
}

func (c Customer) TableName() string {
	return "customer"
}
