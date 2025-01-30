package skillHandler

import "github.com/gofiber/fiber/v2"

func GetSkills(c *fiber.Ctx) error {
	return c.Status(200).JSON(fiber.Map{"message": "Get Skills"})
}
