package router

import (
	bioRoute "github.com/2SSK/portfolio_terminal/backend/internals/routes/bio"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api", logger.New())

	bioRoute.SetupBioRoutes(api)
}
