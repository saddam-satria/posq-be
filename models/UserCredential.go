package models

import (
	"time"

	"github.com/google/uuid"
)

type UserCredential struct {
	UserCredentialId uuid.UUID  `json:"userCredentialId" gorm:"column:userCredentialId;type:uuid;primaryKey;default:gen_random_uuid()"`
	Username         string     `json:"email" gorm:"unique;type:varchar(255)"`
	Password         string     `json:"password" gorm:"type:text"`
	CreatedAt        *time.Time `json:"createdAt" gorm:"default_now();column:createdAt"`
	UpdatedAt        *time.Time `json:"updatedAt" gorm:"autoUpdateTime:milli;column:updatedAt"`
	IsDeleted        bool       `json:"isDeleted" gorm:"default:false;column:isDeleted"`

	UserProfile *UserProfile `gorm:"foreignKey:UserCredentialId"`
	Orders      []Order      `gorm:"foreignKey:UserCredentialId"`
}

func (u UserCredential) TableName() string {
	return "userCredential"
}
