package databaseRoute

import (
	databaseHandler "github.com/2SSK/portfolio_terminal/backend/internals/handlers/skills/databases"
	"github.com/gofiber/fiber/v2"
)

func SetupDatabaseRoutes(router fiber.Router) {
	db := router.Group("/db")

	db.Get("/", databaseHandler.GetDb) // /api/skill/db/
	db.Post("/", databaseHandler.SetDb)
	db.Delete("/:db", databaseHandler.DeleteDb)
}
