package userRoute

import (
	userHandler "github.com/2SSK/portfolio_terminal/backend/internals/handlers"
	"github.com/gofiber/fiber/v2"
)

func SetupUserRoute(router fiber.Router) {
	user := router.Group("/user")

	user.Post("/", userHandler.Login)
}
