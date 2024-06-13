package order

import (
	"github.com/gofiber/fiber/v2"
	"github.com/saddam-satria/posq-be/services/order/endpoints"
)

func OrderRoute(route fiber.Router) {
	route.Post("/order", endpoints.Checkout)
	route.Get("/order/report", endpoints.GetOrders)
	route.Post("/order/:referenceId/receipt", endpoints.GetReceipt)
}
