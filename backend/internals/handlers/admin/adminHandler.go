package adminHandler

import (
	"context"

	"github.com/2SSK/portfolio_terminal/backend/prisma/db"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func VerifyAdmin(c *fiber.Ctx) error {
	// Parse the request body
	var body LoginRequest
	if err := c.BodyParser(&body); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid Username or Password"})
	}

	// check if username or password is empy
	if body.Username == "" || body.Password == "" {
		return c.Status(400).JSON(fiber.Map{"error": "username and password are required"})
	}

	// Prisma client
	prisma := db.NewClient()
	if err := prisma.Connect(); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Could not connect to the database"})
	}
	defer prisma.Disconnect()

	// Query the  database
	admin, err := prisma.Admin.FindUnique(db.Admin.Username.Equals(body.Username)).Exec(context.Background())
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "User not found"})
	}

	// Verify the password
	err = bcrypt.CompareHashAndPassword([]byte(admin.Password), []byte(body.Password))
	if err != nil {
		return c.Status(401).JSON(fiber.Map{"error": "Invalid Username or Password"})
	}

	// If everything is fine, return a success message
	return c.Status(200).JSON(fiber.Map{"message": "Admin verified"})
}
