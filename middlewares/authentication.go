package middlewares

import (
	"fmt"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/saddam-satria/posq-be/commons"
	"github.com/saddam-satria/posq-be/domains"
	"github.com/saddam-satria/posq-be/domains/apis"
	"github.com/saddam-satria/posq-be/models"
	"github.com/saddam-satria/posq-be/repositories"
)

func Authentication(ctx *fiber.Ctx) error {
	header := new(apis.AuthHeader)
	code := fiber.StatusUnauthorized
	if err := ctx.ReqHeaderParser(header); err != nil {
		ctx.Status(code).JSON(commons.GetResponse[any](commons.ACCESS_DENIED[domains.En], code, nil))
		return nil
	}

	token := header.Authorization

	tokenParsed := strings.Split(token, " ")

	if len(tokenParsed) != 2 {
		ctx.Status(code).JSON(commons.GetResponse[any](commons.ACCESS_DENIED[domains.En], code, nil))
		return nil
	}

	if tokenParsed[0] != commons.GoDotEnv("AUTH_HEADER_KEY") {
		ctx.Status(code).JSON(commons.GetResponse[any](commons.ACCESS_DENIED[domains.En], code, nil))
		return nil
	}

	jwtToken := tokenParsed[1]

	parsedToken, err := commons.VerifyToken(jwtToken)

	if err != nil {
		ctx.Status(code).JSON(commons.GetResponse[any](err.Error(), code, nil))
		return nil
	}

	userId := fmt.Sprintf("%v", parsedToken["id"])
	var userCredential models.UserCredential
	userUuid, _ := uuid.Parse(userId)

	if err := repositories.FindUserCredentialById(userUuid, &userCredential); err != nil {
		ctx.Status(code).JSON(commons.GetResponse[any](commons.ACCESS_DENIED[domains.En], code, nil))
		return nil
	}

	return ctx.Next()
}
