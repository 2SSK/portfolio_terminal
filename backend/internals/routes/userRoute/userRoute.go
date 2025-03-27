package userRoute

import (
	"github.com/2SSK/portfolio_terminal/backend/internals/handlers/userHandler"
	"github.com/gofiber/fiber/v2"
)

func SetupUserRoute(router fiber.Router) {
	user := router.Group("/user")

	user.Post("/register", userHandler.CreateUser)
	user.Post("/login", userHandler.GetUser)
	user.Post("/refresh", userHandler.GetRefreshToken)
	user.Get("/verify", userHandler.VerifyToken)
}
