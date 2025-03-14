package router

import (
	userRoute "github.com/2SSK/portfolio_terminal/backend/internals/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func SetupRoutes(server *fiber.App) {
	api := server.Group("/api", logger.New())

	userRoute.SetupUserRoute(api)
}
