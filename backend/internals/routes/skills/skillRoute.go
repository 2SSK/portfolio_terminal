package skillRoute

import (
	databaseRoute "github.com/2SSK/portfolio_terminal/backend/internals/routes/skills/databases"
	programmingLangRoute "github.com/2SSK/portfolio_terminal/backend/internals/routes/skills/porgrammingLang"
	toolRoute "github.com/2SSK/portfolio_terminal/backend/internals/routes/skills/tools"
	"github.com/gofiber/fiber/v2"
)

func SetupSkillRoute(router fiber.Router) {
	skill := router.Group("/skill")

	programmingLangRoute.SetupProgrammingLangRoute(skill) // /api/skill/pl
	databaseRoute.SetupDatabaseRoutes(skill)              // /api/skill/db
	toolRoute.SetupToolRoute(skill)                       // /api/skill/tool
}
