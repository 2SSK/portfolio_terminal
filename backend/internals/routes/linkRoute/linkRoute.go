package linkRoute

import (
	linkHandler "github.com/2SSK/portfolio_terminal/backend/internals/handlers/linkHandler"
	"github.com/2SSK/portfolio_terminal/backend/utils/middleware"
	"github.com/gofiber/fiber/v2"
)

func SetupLinkRoute(router fiber.Router) {
	link := router.Group("/link")

	// Apply middleware
	link.Use(middleware.VerifyUser())

	link.Get("/", linkHandler.GetAllLinks)

	link.Get("/:id", linkHandler.GetLink)

	link.Post("/", linkHandler.CreateLink)

	link.Put("/:id", linkHandler.UpdateLink)

	link.Delete("/:id", linkHandler.DeleteLink)
}
