package resumeRoute

import (
	"github.com/2SSK/portfolio_terminal/backend/internals/handlers/resumeHandler"
	"github.com/gofiber/fiber/v2"
)

func SetupResumeRoute(router fiber.Router) {
	resume := router.Group("/resume")

	resume.Get("/", resumeHandler.GetResume)
	resume.Post("/upload", resumeHandler.AddResume)
	resume.Delete("/", resumeHandler.DeleteResume)
}
