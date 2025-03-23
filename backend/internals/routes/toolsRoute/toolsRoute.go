package toolsRoute

import (
	"github.com/2SSK/portfolio_terminal/backend/internals/handlers/toolsHandler"
	"github.com/gofiber/fiber/v2"
)

func SetupToolsRoute(router fiber.Router) {
	tools := router.Group("/tools")

	tools.Get("/", toolsHandler.GetAllTools)
	tools.Get("/programming-lang", toolsHandler.GetAllProgrammingLangs)
	tools.Get("/software-tools", toolsHandler.GetAllSoftwareTools)
	tools.Get("/frameworks", toolsHandler.GetAllFrameworks)
	tools.Get("/databases", toolsHandler.GetAllDatabases)

	tools.Post("/programming-lang", toolsHandler.AddProgrammingLang)
	tools.Post("/software-tool", toolsHandler.AddSoftwareTools)
	tools.Post("/framework", toolsHandler.AddFramework)
	tools.Post("/database", toolsHandler.AddDatabase)

	tools.Delete("/programming-lang", toolsHandler.DeleteProgrammingLang)
	tools.Delete("/software-tool", toolsHandler.DeleteSoftwareTool)
	tools.Delete("/framework", toolsHandler.DeleteFramework)
	tools.Delete("/database", toolsHandler.DeleteDatabase)
}
