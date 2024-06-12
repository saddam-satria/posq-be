package models

import (
	"time"

	"github.com/google/uuid"
)

type ProductVariant struct {
	ProductVariantId uuid.UUID  `json:"id" gorm:"column:productVariantId;type:uuid;primaryKey;default:gen_random_uuid()"`
	Name             string     `json:"name" gorm:"column:name;type:varchar(255)"`
	BasePrice        float32    `json:"base_price" gorm:"column:basePrice;type:decimal(10,2)"`
	SalePrice        float32    `json:"sale_price" gorm:"column:salePrice;type:decimal(10,2)"`
	Brand            string     `json:"brand" gorm:"column:brand;type:varchar(255)"`
	Stock            int        `json:"stock" gorm:"column:stock;type:integer"`
	Sku              string     `json:"sku" gorm:"column:sku;type:varchar(15);unique"`
	IsTrack          bool       `json:"track_stock" gorm:"column:isTrack;default:false"`
	CreatedAt        *time.Time `json:"-" gorm:"default_now();column:createdAt"`
	UpdatedAt        *time.Time `json:"-" gorm:"autoUpdateTime:milli;column:updatedAt"`
	IsDeleted        bool       `json:"-" gorm:"default:false;column:isDeleted"`
	IsAvailable      bool       `json:"is_available" gorm:"default:true;column:isAvailable"`
	ProductId        string     `json:"item_id" gorm:"column:product_id;type:uuid"`

	OrderProducts []OrderProduct `json:"items" gorm:"foreignKey:ProductVariantId"`
}

func (u ProductVariant) TableName() string {
	return "productVariant"
}
