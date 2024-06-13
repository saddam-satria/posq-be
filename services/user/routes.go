package user

import (
	"github.com/gofiber/fiber/v2"
	"github.com/saddam-satria/posq-be/services/user/endpoints"
)

func UserRoute(route fiber.Router) {
	route.Post("/auth/me", endpoints.GetUserProfile)
}
