package endpoints

import (
	"github.com/gofiber/fiber/v2"
	"github.com/saddam-satria/posq-be/commons"
	"github.com/saddam-satria/posq-be/domains"
	"github.com/saddam-satria/posq-be/models"
	"github.com/saddam-satria/posq-be/repositories"
)

func GetVariants(ctx *fiber.Ctx) error {
	itemId := ctx.Params("itemId")

	var variants []models.ProductVariant

	repositories.GetVariants(itemId, &variants)
	code := fiber.StatusNotFound
	if variants == nil || !commons.IsUUID(itemId) || len(variants) < 1 {
		ctx.Status(code).JSON(commons.GetResponse[any](commons.NOT_FOUND[domains.En], code, nil))
		return nil
	}

	ctx.Status(fiber.StatusOK).JSON(commons.GetResponse("success", fiber.StatusOK, variants))

	return nil
}
