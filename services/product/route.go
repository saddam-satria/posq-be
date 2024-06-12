package product

import (
	"github.com/gofiber/fiber/v2"
	"github.com/saddam-satria/posq-be/services/product/endpoints"
)

func ProductRoute(route fiber.Router) {
	route.Get("/item", endpoints.GetProducts)
	route.Get("/item/:itemId/variants", endpoints.GetVariants)
}
