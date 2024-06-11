package middlewares

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/saddam-satria/posq-be/commons"
	"github.com/saddam-satria/posq-be/domains/apis"
)

func Authentication(ctx *fiber.Ctx) error {
	header := new(apis.AuthHeader)

	if err := ctx.ReqHeaderParser(header); err != nil {
		ctx.SendStatus(fiber.StatusUnauthorized)
		ctx.JSON(commons.GetResponse[any]("header error", fiber.StatusUnauthorized, nil))
	}

	token := header.Authorization

	tokenParsed := strings.Split(token, " ")

	if len(tokenParsed) != 2 {
		ctx.SendStatus(fiber.StatusUnauthorized)
		ctx.JSON(commons.GetResponse[any]("access denied", fiber.StatusUnauthorized, nil))
		return nil
	}

	if tokenParsed[0] != commons.GoDotEnv("AUTH_HEADER_KEY") {
		ctx.SendStatus(fiber.StatusUnauthorized)
		ctx.JSON(commons.GetResponse[any]("access denied", fiber.StatusUnauthorized, nil))
		return nil
	}

	jwtToken := tokenParsed[1]

	_, err := commons.VerifyToken(jwtToken)

	if err != nil {
		ctx.SendStatus(fiber.StatusUnauthorized)
		ctx.JSON(commons.GetResponse[any](err.Error(), fiber.StatusUnauthorized, nil))
		return nil
	}

	return ctx.Next()
}
