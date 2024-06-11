package repositories

import (
	"github.com/google/uuid"
	"github.com/saddam-satria/posq-be/commons"
	"github.com/saddam-satria/posq-be/models"
)


func FindUserProfileByUserCredentialId(userCredentialId uuid.UUID, result *models.UserCredential) {
	commons.DatabaseConnection.Model(&models.UserCredential{}).Where(&models.UserCredential{UserCredentialId: userCredentialId}).Preload("UserProfile").Find(&result)
}