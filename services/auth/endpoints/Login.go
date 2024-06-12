package endpoints

import (
	"github.com/gofiber/fiber/v2"
	"github.com/saddam-satria/posq-be/commons"
	"github.com/saddam-satria/posq-be/domains"
	"github.com/saddam-satria/posq-be/domains/apis"
	"github.com/saddam-satria/posq-be/models"
	"github.com/saddam-satria/posq-be/repositories"
)

func Login(ctx *fiber.Ctx) error {
	request := new(apis.LoginRequest)

	code := fiber.StatusBadRequest
	if err := ctx.BodyParser(request); err != nil {
		ctx.Status(code).JSON(commons.GetResponse[any](commons.BAD_REQUEST[domains.En], code, nil))
		return nil
	}

	var result models.UserCredential
	username := request.Username
	password := request.Password

	repositories.FindUserCredentialByUsername(username, &result)

	isPasswordMatch := commons.VerifyHashed(password, result.Password)

	if !isPasswordMatch {
		ctx.Status(code).JSON(commons.GetResponse[any](commons.NOT_FOUND[domains.En], code, nil))
		return nil
	}

	accessToken, _ := commons.GenerateToken(result.UserCredentialId.String(), result.Username, 1000*60*60*24*30, commons.GoDotEnv("SECRET_KEY"))
	refreshToken, _ := commons.GenerateToken(result.UserCredentialId.String(), result.Username, 1000*60*60*24*30, commons.GoDotEnv("SECRET_KEY"))

	response := apis.LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	ctx.Status(fiber.StatusOK).JSON(commons.GetResponse(commons.SUCCESS[domains.En], fiber.StatusOK, response))

	return nil
}
