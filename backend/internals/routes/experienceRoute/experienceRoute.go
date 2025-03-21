package experienceRoute

import (
	"github.com/2SSK/portfolio_terminal/backend/internals/handlers/experienceHandler"
	"github.com/2SSK/portfolio_terminal/backend/utils/middleware"
	"github.com/gofiber/fiber/v2"
)

func SetupExperienceRoute(router fiber.Router) {
	experience := router.Group("/experience")

	experience.Use(middleware.VerifyUser())

	experience.Post("/", experienceHandler.AddExperience)
	experience.Get("/", experienceHandler.GetAllExperience)
	experience.Get("/:id", experienceHandler.GetExperience)
	experience.Put("/:id", experienceHandler.UpdateExperience)
	experience.Delete("/:id", experienceHandler.DeleteExperience)
}
