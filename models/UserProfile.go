package models

import "time"

type UserProfile struct {
	UserProfileId  string     `json:"id" gorm:"column:userProfileId;type:uuid;primaryKey;default:gen_random_uuid()"`
	Name           string     `json:"name" gorm:"column:name;type:varchar(255)"`
	Email          string     `json:"email" gorm:"column:email;unique;type:varchar(255)"`
	PhoneNumber    string     `json:"phone_number" gorm:"column:phoneNumber;type:varchar(15);unique"`
	ProfilePicture *string     `json:"profile_picture" gorm:"column:profilePicture;type:text"`
	CreatedAt      *time.Time `json:"createdAt" gorm:"default_now();column:createdAt"`
	UpdatedAt      *time.Time `json:"updatedAt" gorm:"autoUpdateTime:milli;column:updatedAt"`

	UserCredentialId string `json:"user_credential_id" gorm:"column:user_credential_id;type:uuid;unique"`	
}

func(u UserProfile) TableName() string{
	return "userProfile"
}