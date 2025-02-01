package socialRoute

import (
	socialHandler "github.com/2SSK/portfolio_terminal/backend/internals/handlers/socials"
	"github.com/gofiber/fiber/v2"
)

func SetupSocailRoute(router fiber.Router) {
	social := router.Group("/social")

	social.Get("/", socialHandler.GetSocial)
	social.Get("/:title", socialHandler.GetSpecificSocial)
	social.Post("/", socialHandler.SetSocial)
	social.Delete("/:title", socialHandler.DeleteSocial)
}
