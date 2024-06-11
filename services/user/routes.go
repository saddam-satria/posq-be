package user

import (
	"github.com/gofiber/fiber/v2"
	"github.com/saddam-satria/posq-be/middlewares"
	"github.com/saddam-satria/posq-be/services/user/endpoints"
)







func UserRoute(route *fiber.App){
	route.Get("/api/v1/me", middlewares.Authentication,endpoints.GetUserProfile)
}