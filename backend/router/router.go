package router

import (
	"github.com/2SSK/portfolio_terminal/backend/internals/routes/linkRoute"
	"github.com/2SSK/portfolio_terminal/backend/internals/routes/resumeRoute"
	"github.com/2SSK/portfolio_terminal/backend/internals/routes/toolsRoute"
	"github.com/2SSK/portfolio_terminal/backend/internals/routes/userRoute"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func SetupRoutes(server *fiber.App) {
	api := server.Group("/api", logger.New())

	userRoute.SetupUserRoute(api)
	linkRoute.SetupLinkRoute(api)
	resumeRoute.SetupResumeRoute(api)
	toolsRoute.SetupToolsRoute(api)
}
