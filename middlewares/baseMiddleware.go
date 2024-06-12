package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"github.com/saddam-satria/posq-be/commons"
	"github.com/saddam-satria/posq-be/domains"
	"github.com/saddam-satria/posq-be/domains/apis"
)

func BaseMiddleware(ctx *fiber.Ctx) error {
	headerObj := new(apis.ApiKeyHeader)
	code := fiber.StatusUnauthorized
	if err := ctx.ReqHeaderParser(headerObj); err != nil {
		ctx.Status(code).JSON(commons.GetResponse[any](commons.ACCESS_DENIED[domains.En], code, nil))
		return nil
	}

	apiKey := headerObj.XApiKey

	if apiKey != commons.GoDotEnv("API_KEY") {
		ctx.Status(code).JSON(commons.GetResponse[any](commons.ACCESS_DENIED[domains.En], code, nil))
		return nil
	}

	ctx.Next()
	return nil
}
