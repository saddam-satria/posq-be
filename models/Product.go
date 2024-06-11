package models

import (
	"time"

	"github.com/google/uuid"
)

type Product struct {
	ProductId   uuid.UUID  `json:"id" gorm:"column:productId;type:uuid;primaryKey;default:gen_random_uuid()"`
	Name        string     `json:"name" gorm:"column:name;type:varchar(255)"`
	Image       *string    `json:"image" gorm:"column:image;type:text"`
	Category    string     `json:"category" gorm:"column:category;"`
	CreatedAt   *time.Time `json:"-" gorm:"default_now();column:createdAt"`
	UpdatedAt   *time.Time `json:"-" gorm:"autoUpdateTime:milli;column:updatedAt"`
	IsDeleted   bool       `json:"-" gorm:"default:false;column:isDeleted"`
	IsAvailable bool       `json:"is_available" gorm:"default:true;column:isAvailable"`

	ProductVariants []ProductVariant `json:"variants" gorm:"foreignKey:ProductId"`

	Price      float32 `json:"price" gorm:"-" omitempty:"true"`
	RangePrice string  `json:"price_range" gorm:"-" omitempty:"true"`
}

func (u Product) TableName() string {
	return "product"
}
