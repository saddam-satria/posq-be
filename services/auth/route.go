package auth

import (
	"github.com/gofiber/fiber/v2"
	"github.com/saddam-satria/posq-be/services/auth/endpoints"
)

func AuthRoute(route fiber.Router) {
	route.Post("/auth/login", endpoints.Login)
}
