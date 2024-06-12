package endpoints

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/saddam-satria/posq-be/commons"
	"github.com/saddam-satria/posq-be/domains"
	"github.com/saddam-satria/posq-be/domains/apis"
	"github.com/saddam-satria/posq-be/models"
	"github.com/saddam-satria/posq-be/repositories"
)

func GetUserProfile(ctx *fiber.Ctx) error {
	header := new(apis.AuthHeader)

	ctx.ReqHeaderParser(header)

	token := header.Authorization

	parsedToken, _ := commons.VerifyToken(token)

	userId := fmt.Sprintf("%v", parsedToken["id"])

	var userProfile models.UserCredential

	userUuid, _ := uuid.Parse(userId)

	repositories.FindUserProfileByUserCredentialId(userUuid, &userProfile)

	code := fiber.StatusNotFound
	if userProfile.UserProfile == nil {
		ctx.Status(code).JSON(commons.GetResponse[any](commons.NOT_FOUND[domains.En], code, nil))
		return nil
	}

	profile := apis.ProfileResponse{
		Id:             userProfile.UserProfile.UserProfileId,
		Name:           userProfile.UserProfile.Name,
		Email:          userProfile.UserProfile.Email,
		Username:       userProfile.Username,
		PhoneNumber:    userProfile.UserProfile.PhoneNumber,
		ProfilePicture: userProfile.UserProfile.ProfilePicture,
		Role:           "",
		RoleName:       "kasir",
		BusinessId:     "",
	}

	ctx.Status(fiber.StatusOK).JSON(commons.GetResponse(commons.SUCCESS[domains.En], fiber.StatusOK, profile))

	return nil
}
