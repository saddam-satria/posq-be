package endpoints

import (
	"github.com/gofiber/fiber/v2"
	"github.com/saddam-satria/posq-be/commons"
	"github.com/saddam-satria/posq-be/domains"
	"github.com/saddam-satria/posq-be/domains/apis"
	"github.com/saddam-satria/posq-be/repositories"
)

func GetReceipt(ctx *fiber.Ctx) error {

	referenceId := ctx.Params("referenceId")

	var order apis.ReceiptResponse

	if err := repositories.GetOrderReceipt(referenceId, &order); err != nil {
		ctx.Status(fiber.StatusNotFound).JSON(commons.GetResponse[any](commons.NOT_FOUND[domains.En], fiber.StatusNotFound, nil))
		return nil
	}

	totalAmount := float32(0)
	for _, product := range order.Products {
		totalAmount += float32(product.Quantity) * product.SalePrice
	}

	order.SubTotal = totalAmount
	order.TotalItem = len(order.Products)
	order.Logo = "http://localhost:5000/assets/posq.png"
	order.Changed = order.Tendered - totalAmount
	order.OrderTimestamp = order.CreatedAt.Unix()
	code := fiber.StatusOK

	ctx.Status(code).JSON(commons.GetResponse(commons.SUCCESS[domains.En], code, order))
	return nil
}
