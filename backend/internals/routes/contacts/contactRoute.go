package contactRoute

import (
	contactHandler "github.com/2SSK/portfolio_terminal/backend/internals/handlers/contact"
	"github.com/gofiber/fiber/v2"
)

func SetupContactRoute(router fiber.Router) {
	contact := router.Group("/contact")

	contact.Get("/", contactHandler.GetContact)
	contact.Get("/:contactType", contactHandler.GetSpecificContact)
	contact.Post("/", contactHandler.SetContact)
	contact.Delete("/:contactType", contactHandler.DeleteContact)
}
