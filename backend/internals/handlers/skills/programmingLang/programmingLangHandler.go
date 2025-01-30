package programminglangHandler

import (
	"strings"

	"github.com/2SSK/portfolio_terminal/backend/prisma/db"
	"github.com/gofiber/fiber/v2"
)

type ProgrammingLang struct {
	Language string `json:"language"`
}

func SetProgrammingLang(c *fiber.Ctx) error {
	// Parse the body data
	var body ProgrammingLang
	if err := c.BodyParser(&body); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid Input"})
	}

	// Convert the language to lowercase
	input := strings.ToLower(body.Language)

	// Prisma client
	client := db.NewClient()
	if err := client.Connect(); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Could not connect to the database"})
	}
	defer client.Disconnect()

	// Check if the programming language already exists
	existingLang, _ := client.ProgrammingLang.FindUnique(
		db.ProgrammingLang.LanguageName.Equals(input),
	).Exec(c.Context())

	// If the programming language already exists, return an error message
	if existingLang != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Programming Language already exists"})
	}

	// Create the programming language
	_, err := client.ProgrammingLang.CreateOne(
		db.ProgrammingLang.LanguageName.Set(input),
	).Exec(c.Context())
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Error while creating the programming language"})
	}

	// if everything is fine, return a success message
	return c.Status(200).JSON(fiber.Map{"message": "Programming Language Added successfully"})
}

func GetProgrammingLang(c *fiber.Ctx) error {
	// Prisma client
	client := db.NewClient()
	if err := client.Connect(); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Could not connect to the database"})
	}
	defer client.Disconnect()

	// Query the Database
	programmingLangs, err := client.ProgrammingLang.FindMany().Exec(c.Context())
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Error while fetching the programming languages"})
	}
	if programmingLangs == nil {
		return c.Status(204).JSON(fiber.Map{"error": "Programming Languages not found"})
	}

	// If everything is fine, return the programming languages
	return c.Status(200).JSON(fiber.Map{"message": "Programming Languages fetched successfully", "programmingLangs": programmingLangs})
}

func DeleteProgrammingLang(c *fiber.Ctx) error {
	// Get the programming language from the URL
	lang := c.Params("language")

	// Convert the language to lowercase
	lang = strings.ToLower(lang)

	// Prisma client
	client := db.NewClient()
	if err := client.Connect(); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Could not connect to the database"})
	}
	defer client.Disconnect()

	// Delete the programming language
	_, err := client.ProgrammingLang.FindUnique(
		db.ProgrammingLang.LanguageName.Equals(lang),
	).Delete().Exec(c.Context())
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"message": "Error while deleting the programming language", "error": err.Error()})
	}

	// if everything is fine, return a success message
	return c.Status(200).JSON(fiber.Map{"message": "Programming Language deleted successfully"})
}
