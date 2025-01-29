package router

import (
	adminRoute "github.com/2SSK/portfolio_terminal/backend/internals/routes/admin"
	bioRoute "github.com/2SSK/portfolio_terminal/backend/internals/routes/bio"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api", logger.New())

	adminRoute.SetupAdminRoutes(api)
	bioRoute.SetupBioRoutes(api)
}
