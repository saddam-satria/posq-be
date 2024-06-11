package utils

import (
	"github.com/gofiber/fiber/v2"
	"github.com/saddam-satria/posq-be/services/auth"
	"github.com/saddam-satria/posq-be/services/product"
	"github.com/saddam-satria/posq-be/services/user"
)

func GetRoute(route *fiber.App) {
	auth.AuthRoute(route)
	user.UserRoute(route)
	product.ProductRoute(route)
}
