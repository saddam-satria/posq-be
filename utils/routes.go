package utils

import (
	"github.com/gofiber/fiber/v2"
	"github.com/saddam-satria/posq-be/middlewares"
	"github.com/saddam-satria/posq-be/services/auth"
	"github.com/saddam-satria/posq-be/services/customer"
	"github.com/saddam-satria/posq-be/services/order"
	"github.com/saddam-satria/posq-be/services/product"
	"github.com/saddam-satria/posq-be/services/report"
	"github.com/saddam-satria/posq-be/services/user"
)

func GetRoute(route fiber.Router) {
	originRoute := route.Group("/")
	auth.AuthRoute(originRoute)

	protectedRoute := route.Group("/", middlewares.Authentication)
	user.UserRoute(protectedRoute)
	product.ProductRoute(protectedRoute)
	customer.CustomerRoute(protectedRoute)
	order.OrderRoute(protectedRoute)
	report.ReportRoute(protectedRoute)
}
