package repositories

import (
	"github.com/saddam-satria/posq-be/commons"
	"github.com/saddam-satria/posq-be/models"
)

func FindUserCredentialByUsername(username string, result *models.UserCredential){
	commons.DatabaseConnection.Where(&models.UserCredential{Username: username}).Find(&result)
}