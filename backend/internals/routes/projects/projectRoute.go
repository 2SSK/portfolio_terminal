package projectRoute

import (
	projectHandler "github.com/2SSK/portfolio_terminal/backend/internals/handlers/project"
	"github.com/gofiber/fiber/v2"
)

func SetupProjectRoute(router fiber.Router) {
	project := router.Group("/project")

	project.Get("/", projectHandler.GetProject)
	project.Post("/", projectHandler.SetProject)
	project.Delete("/:projectName", projectHandler.DeleteProject)
}
