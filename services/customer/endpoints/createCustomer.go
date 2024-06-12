package endpoints

import (
	"github.com/gofiber/fiber/v2"
	"github.com/saddam-satria/posq-be/commons"
	"github.com/saddam-satria/posq-be/domains"
	"github.com/saddam-satria/posq-be/domains/apis"
	"github.com/saddam-satria/posq-be/models"
	"github.com/saddam-satria/posq-be/repositories"
)

func CreateCustomer(ctx *fiber.Ctx) error {
	newCustomer := new(apis.CreateCustomerRequest)
	code := fiber.StatusBadRequest
	if err := ctx.BodyParser(newCustomer); err != nil {
		ctx.Status(code).JSON(commons.GetResponse[any](commons.BAD_REQUEST[domains.En], code, nil))
		return nil
	}

	var insertedCustomer models.Customer

	customer := models.Customer{
		Name:        newCustomer.Name,
		Email:       newCustomer.Email,
		Phonenumber: newCustomer.Phone,
		Address:     newCustomer.Address,
	}

	if err := repositories.CreateCustomer(&customer, &insertedCustomer).Error; err != nil {
		ctx.Status(code).JSON(commons.GetResponse[any](commons.DUPLICATE_ENTRY[domains.En], code, nil))
		return nil
	}

	ctx.Status(fiber.StatusCreated).JSON(commons.GetResponse(commons.SUCCESS[domains.En], fiber.StatusCreated, customer))
	return nil
}
