package endpoints

import (
	"github.com/gofiber/fiber/v2"
	"github.com/saddam-satria/posq-be/commons"
	"github.com/saddam-satria/posq-be/domains"
	"github.com/saddam-satria/posq-be/domains/apis"
	"github.com/saddam-satria/posq-be/repositories"
)

func GetOrders(ctx *fiber.Ctx) error {
	var orders []apis.OrderResponse
	code := fiber.StatusOK
	repositories.GetOrders(&orders)

	for index, item := range orders {
		totalAmount := float32(0)

		for _, product := range item.Products {
			totalAmount += float32(product.Quantity) * product.SalePrice
		}

		orders[index].SubTotal = totalAmount
		orders[index].Changed = item.Tendered - totalAmount
		orders[index].OrderTimestamp = item.CreatedAt.Unix()
	}

	ctx.Status(code).JSON(commons.GetResponse(commons.SUCCESS[domains.En], code, orders))
	return nil
}
