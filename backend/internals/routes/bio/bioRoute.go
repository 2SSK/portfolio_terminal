package bioRoute

import (
	bioHandler "github.com/2SSK/portfolio_terminal/backend/internals/handlers/bio"
	"github.com/gofiber/fiber/v2"
)

func SetupBioRoutes(router fiber.Router) {
	bio := router.Group("/bio")

	bio.Get("/", bioHandler.GetBio)
	bio.Post("/", bioHandler.SetBio)
}
