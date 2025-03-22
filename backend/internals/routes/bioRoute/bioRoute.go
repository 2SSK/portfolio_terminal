package bioRoute

import (
	"github.com/2SSK/portfolio_terminal/backend/internals/handlers/bioHandler"
	"github.com/2SSK/portfolio_terminal/backend/utils/middleware"
	"github.com/gofiber/fiber/v2"
)

func SetupBioRoute(router fiber.Router) {
	bio := router.Group("/bio")

	bio.Use(middleware.VerifyUser())

	bio.Get("/", bioHandler.GetBio)
	bio.Post("/", bioHandler.AddBio)
}
