package repositories

import (
	"github.com/google/uuid"
	"github.com/saddam-satria/posq-be/commons"
	"github.com/saddam-satria/posq-be/models"
)

func FindUserCredentialByUsername(username string, result *models.UserCredential) {
	commons.DatabaseConnection.Where(&models.UserCredential{Username: username}).Find(&result)
}

func FindUserCredentialById(userCredentialId uuid.UUID, result *models.UserCredential) error {
	return commons.DatabaseConnection.Model(&models.UserCredential{}).Where(&models.UserCredential{UserCredentialId: userCredentialId}).First(&result).Error
}
