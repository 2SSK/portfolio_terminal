package projectRoute

import (
	"github.com/2SSK/portfolio_terminal/backend/internals/handlers/projectHandler"
	"github.com/gofiber/fiber/v2"
)

func SetupProjectRoute(router fiber.Router) {
	projects := router.Group("/projects")

	// Get all projects
	projects.Get("/", projectHandler.GetAllProjects)
	// // Get specific project
	projects.Get("/:id", projectHandler.GetProject)
	// Add a project
	projects.Post("/", projectHandler.AddProject)
	// Update a project
	projects.Put("/:id", projectHandler.UpdateProject)
	// // Delete a project
	projects.Delete("/:id", projectHandler.DeleteProject)
}
