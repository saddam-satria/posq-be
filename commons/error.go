package commons

import (
	"errors"

	"github.com/gofiber/fiber/v2"
)

func ErrorHandler(ctx *fiber.Ctx, err error) error {

	defaultCode := fiber.StatusInternalServerError

	var e *fiber.Error

	if errors.As(err, &e) {
		defaultCode = e.Code
	}

	ctx.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)
	ctx.Status(defaultCode).JSON(GetResponse[any](e.Message, defaultCode, nil))

	return nil

}
