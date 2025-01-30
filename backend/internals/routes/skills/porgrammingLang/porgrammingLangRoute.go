package programmingLangRoute

import (
	programminglangHandler "github.com/2SSK/portfolio_terminal/backend/internals/handlers/skills/programmingLang"
	"github.com/gofiber/fiber/v2"
)

func SetupProgrammingLangRoute(router fiber.Router) {
	pl := router.Group("/pl")

	pl.Get("/", programminglangHandler.GetProgrammingLang) // /api/skill/pl/
	pl.Post("/", programminglangHandler.SetProgrammingLang)
	pl.Delete("/:language", programminglangHandler.DeleteProgrammingLang)
}
