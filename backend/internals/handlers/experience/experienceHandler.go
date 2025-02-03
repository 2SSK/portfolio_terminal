package experienceHandler

import (
	"strings"

	"github.com/2SSK/portfolio_terminal/backend/prisma/db"
	"github.com/gofiber/fiber/v2"
)

type experienceStruct struct {
	Title            string `json:"title"`
	Company          string `json:"company"`
	Responsibilities string `json:"responsibilities"`
	From             string `json:"from"`
	To               string `json:"to"`
}

func GetExperience(c *fiber.Ctx) error {
	// Prisma client
	client := db.NewClient()
	if err := client.Connect(); err != nil {
		return c.Status(500).JSON(fiber.Map{"message": "Error connecting to the database"})
	}
	defer client.Disconnect()

	experiences, err := client.Experience.FindMany().Exec(c.Context())
	if len(experiences) == 0 {
		return c.Status(404).JSON(fiber.Map{"message": "No experiences found", "data": experiences})
	}
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"message": "Error fetching experiences"})
	}

	// If everything is fine, then return
	return c.Status(200).JSON(fiber.Map{"message": "Experiences fetched successfully", "data": experiences})
}

func SetExperience(c *fiber.Ctx) error {
	// Parse form data
	var experience experienceStruct
	if err := c.BodyParser(&experience); err != nil {
		return c.Status(400).JSON(fiber.Map{"message": "Error parsing input data", "error": err.Error()})
	}

	// Prisma client
	client := db.NewClient()
	if err := client.Connect(); err != nil {
		return c.Status(500).JSON(fiber.Map{"message": "Error connecting to the database"})
	}
	defer client.Disconnect()

	// Create or Update a new experience
	_, err := client.Experience.UpsertOne(
		db.Experience.Company.Equals(experience.Company),
	).Create(
		db.Experience.Title.Set(experience.Title),
		db.Experience.Company.Set(strings.ReplaceAll(strings.ToLower(experience.Company), " ", "_")),
		db.Experience.Responsibilities.Set(experience.Responsibilities),
		db.Experience.From.Set(experience.From),
		db.Experience.To.Set(experience.To),
	).Update(
		db.Experience.Title.Set(experience.Title),
		db.Experience.Company.Set(experience.Company),
		db.Experience.Responsibilities.Set(experience.Responsibilities),
		db.Experience.From.Set(experience.From),
		db.Experience.To.Set(experience.To),
	).Exec(c.Context())
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"message": "Error setting experience"})
	}

	// If everything is fine, then return
	return c.Status(200).JSON(fiber.Map{"message": "Experience set successfully"})
}

func DeleteExperience(c *fiber.Ctx) error {
	// Get company name from the URL
	company := c.Params("company")

	// Prisma client
	client := db.NewClient()
	if err := client.Connect(); err != nil {
		return c.Status(500).JSON(fiber.Map{"message": "Error connecting to the database"})
	}
	defer client.Disconnect()

	// Delete the experience from the database
	_, err := client.Experience.FindUnique(
		db.Experience.Company.Equals(company),
	).Delete().Exec(c.Context())

	if err != nil {
		return c.Status(500).JSON(fiber.Map{"message": "Error deleting experience", "error": err.Error()})
	}

	// If everything is fine, then return
	return c.Status(200).JSON(fiber.Map{"message": "Experience deleted successfully"})
}
