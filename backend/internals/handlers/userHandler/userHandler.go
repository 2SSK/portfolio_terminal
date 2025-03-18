package userHandler

import (
	"github.com/2SSK/portfolio_terminal/backend/config"
	"github.com/2SSK/portfolio_terminal/backend/prisma/db"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

type input struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func CreateUser(c *fiber.Ctx) error {
	var body input

	if err := c.BodyParser(&body); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request"})
	}

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(body.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to hash password"})
	}

	user, err := config.PrismaClient.User.CreateOne(
		db.User.Email.Set(body.Email),
		db.User.Password.Set(string(hashedPassword)),
	).Exec(c.Context())

	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to create user"})
	}

	return c.Status(201).JSON(user)
}

func GetUser(c *fiber.Ctx) error {
	var body input

	if err := c.BodyParser(&body); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid inputs"})
	}

	user, err := config.PrismaClient.User.FindUnique(
		db.User.Email.Equals(body.Email),
	).Exec(c.Context())
	if err != nil {
		return c.Status(401).JSON(fiber.Map{"error": "User not found"})
	}

	// Compare the password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))
	if err != nil {
		return c.Status(401).JSON(fiber.Map{"error": "Invalid password"})
	}

	return c.Status(200).JSON(user)
}
