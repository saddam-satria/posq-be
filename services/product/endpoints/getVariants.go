package endpoints

import (
	"github.com/gofiber/fiber/v2"
	"github.com/saddam-satria/posq-be/commons"
	"github.com/saddam-satria/posq-be/models"
	"github.com/saddam-satria/posq-be/repositories"
)

func GetVariants(ctx *fiber.Ctx) error {
	itemId := ctx.Params("itemId")

	var variants []models.ProductVariant

	repositories.GetVariants(itemId, &variants)

	if variants == nil || !commons.IsUUID(itemId) || len(variants) < 1 {

		ctx.SendStatus(fiber.StatusNotFound)
		ctx.JSON(commons.GetResponse[any]("data tidak ditemukan", fiber.StatusNotFound, nil))

		return nil
	}

	ctx.SendStatus(fiber.StatusOK)
	ctx.JSON(commons.GetResponse("success", fiber.StatusOK, variants))

	return nil
}
