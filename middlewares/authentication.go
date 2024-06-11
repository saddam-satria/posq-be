package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"github.com/saddam-satria/posq-be/commons"
	"github.com/saddam-satria/posq-be/domains/apis"
)

func Authentication(ctx *fiber.Ctx) error {
	header:= new(apis.AuthHeader)

	
	if err:= ctx.ReqHeaderParser(header); err!=nil{
		ctx.SendStatus(fiber.StatusUnauthorized)
		ctx.JSON(commons.GetResponse[any]("header error", fiber.StatusUnauthorized,nil))
	}

	token:= header.Authorization
	
	_, err:= commons.VerifyToken(token) 

	if err != nil {
		ctx.SendStatus(fiber.StatusUnauthorized)
		ctx.JSON(commons.GetResponse(err.Error(), fiber.StatusUnauthorized,token))
		return nil
	}

	return ctx.Next()
}