package bioHandler

import "github.com/gofiber/fiber/v2"

func GetBio(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"message": "Hello, World!"})
}
