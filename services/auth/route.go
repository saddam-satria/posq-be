package auth

import (
	"github.com/gofiber/fiber/v2"
	"github.com/saddam-satria/posq-be/services/auth/endpoints"
)


func AuthRoute(route *fiber.App){
	route.Post("/api/v1/auth/login", endpoints.Login)
}