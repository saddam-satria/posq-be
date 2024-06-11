package endpoints

import (
	"github.com/gofiber/fiber/v2"
	"github.com/saddam-satria/posq-be/commons"
	"github.com/saddam-satria/posq-be/domains/apis"
	"github.com/saddam-satria/posq-be/models"
	"github.com/saddam-satria/posq-be/repositories"
)

func Login(ctx *fiber.Ctx) error {
	request:= new(apis.LoginRequest)

	if err:= ctx.BodyParser(request);err!= nil{
		ctx.SendStatus(fiber.StatusBadRequest)
		ctx.JSON(commons.GetResponse[any](err.Error(), fiber.StatusBadRequest, nil))
		return nil
	}

	var result models.UserCredential
	username:= request.Username
	password:= request.Password


	repositories.FindUserCredentialByUsername(username,&result)

	isPasswordMatch := commons.VerifyHashed(password,result.Password)

	if !isPasswordMatch {
		ctx.SendStatus(fiber.StatusBadRequest)
		ctx.JSON(commons.GetResponse[any]("user tidak ditemukan", fiber.StatusBadRequest, nil))
		return nil
	}

	accessToken,_:= commons.GenerateToken(result.UserCredentialId.String(),result.Username,1000 * 60 * 60 * 24 * 30, commons.GoDotEnv("SECRET_KEY"))
	refreshToken,_:= commons.GenerateToken(result.UserCredentialId.String(),result.Username,1000 * 60 * 60 * 24 * 30, commons.GoDotEnv("SECRET_KEY"))
	

	response:= apis.LoginResponse{
		AccessToken:accessToken,
		RefreshToken: refreshToken,
	}

	ctx.SendStatus(fiber.StatusOK)
	ctx.JSON(commons.GetResponse("success",fiber.StatusOK,response))

	return nil
}