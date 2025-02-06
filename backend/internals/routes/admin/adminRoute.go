package adminRoute

import (
	adminHandler "github.com/2SSK/portfolio_terminal/backend/internals/handlers/admin"
	"github.com/gofiber/fiber/v2"
)

func SetupAdminRoutes(router fiber.Router) {
	admin := router.Group("/admin")

	admin.Post("/", adminHandler.VerifyAdmin)
}
