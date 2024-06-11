package product

import (
	"github.com/gofiber/fiber/v2"
	"github.com/saddam-satria/posq-be/middlewares"
	"github.com/saddam-satria/posq-be/services/product/endpoints"
)

func ProductRoute(route *fiber.App) {
	route.Get("/api/v1/item", middlewares.Authentication, endpoints.GetProducts)
	route.Get("/api/v1/item/:itemId/variants", middlewares.Authentication, endpoints.GetVariants)
}
