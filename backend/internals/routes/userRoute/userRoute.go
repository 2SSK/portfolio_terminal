package userRoute

import (
	"github.com/2SSK/portfolio_terminal/backend/internals/handlers/userHandler"
	"github.com/gofiber/fiber/v2"
)

func SetupUserRoute(router fiber.Router) {
	user := router.Group("/user")

	user.Post("/", userHandler.CreateUser)
	user.Get("/", userHandler.GetUser)
}
