package middleware

import (
	"github.com/2SSK/portfolio_terminal/backend/config"
	"github.com/2SSK/portfolio_terminal/backend/prisma/db"
	"github.com/gofiber/fiber/v2"
)

func VerifyUser() fiber.Handler {
	return func(c *fiber.Ctx) error {
		userId := c.QueryInt("userId")
		if userId == 0 {
			return c.Status(400).JSON(fiber.Map{
				"error": "UserId is required",
			})
		}

		existingUser, err := config.PrismaClient.User.FindUnique(
			db.User.ID.Equals(userId),
		).Exec(c.Context())
		if err != nil || existingUser == nil {
			return c.Status(404).JSON(fiber.Map{"error": "User not found"})
		}

		// Store uesr in locals for later if needed
		c.Locals("user", existingUser)
		return c.Next()
	}
}
