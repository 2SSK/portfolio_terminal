package toolRoute

import (
	toolHandler "github.com/2SSK/portfolio_terminal/backend/internals/handlers/skills/tools"
	"github.com/gofiber/fiber/v2"
)

func SetupToolRoute(router fiber.Router) {
	tool := router.Group("/tool")

	tool.Get("/", toolHandler.GetTools) // /api/skill/tool/
	tool.Post("/", toolHandler.SetTools)
	tool.Delete("/:tool", toolHandler.DeleteTools)
}
