package endpoints

import (
	"github.com/gofiber/fiber/v2"
	"github.com/saddam-satria/posq-be/commons"
	"github.com/saddam-satria/posq-be/domains"
	"github.com/saddam-satria/posq-be/models"
	"github.com/saddam-satria/posq-be/repositories"
)

func GetCustomers(ctx *fiber.Ctx) error {
	var customers []models.Customer

	repositories.GetCustomers(&customers)

	ctx.Status(fiber.StatusOK).JSON(commons.GetResponse(commons.SUCCESS[domains.En], fiber.StatusOK, customers))

	return nil
}
