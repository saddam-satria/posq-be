package customer

import (
	"github.com/gofiber/fiber/v2"
	"github.com/saddam-satria/posq-be/services/customer/endpoints"
)

func CustomerRoute(route fiber.Router) {
	route.Get("/customer", endpoints.GetCustomers)
	route.Post("/customer", endpoints.CreateCustomer)
}
