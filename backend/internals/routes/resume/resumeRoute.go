package resumeRoute

import (
	resumeHandler "github.com/2SSK/portfolio_terminal/backend/internals/handlers/resume"
	"github.com/gofiber/fiber/v2"
)

func SetupResumeRoute(router fiber.Router) {
	resume := router.Group("/resume")

	resume.Get("/", resumeHandler.GetResume)
	resume.Post("/", resumeHandler.SetResume)
	resume.Delete("/", resumeHandler.DeleteResume)
}
