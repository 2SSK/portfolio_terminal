package experienceRoute

import (
	experienceHandler "github.com/2SSK/portfolio_terminal/backend/internals/handlers/experience"
	"github.com/gofiber/fiber/v2"
)

func SetupExperienceRoute(router fiber.Router) {
	experience := router.Group("/experience")

	experience.Get("/", experienceHandler.GetExperience)
	experience.Post("/", experienceHandler.SetExperience)
	experience.Delete("/:company", experienceHandler.DeleteExperience)
}
