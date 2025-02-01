package socialHandler

import (
	"strings"

	"github.com/2SSK/portfolio_terminal/backend/prisma/db"
	"github.com/gofiber/fiber/v2"
)

type Social struct {
	Title string `json:"title"`
	Url   string `json:"url"`
}

func GetSocial(c *fiber.Ctx) error {
	// Prisma client
	client := db.NewClient()
	if err := client.Connect(); err != nil {
		return c.Status(500).JSON(fiber.Map{"message": "Error connecting to the database"})
	}
	defer client.Disconnect()

	// Get all socials
	socials, err := client.Socials.FindMany().Exec(c.Context())
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"message": "Error fetching socials"})
	}
	if len(socials) == 0 {
		return c.Status(404).JSON(fiber.Map{"message": "No socials found"})
	}

	return c.Status(200).JSON(fiber.Map{"message": "Socials", "data": socials})
}

func GetSpecificSocial(c *fiber.Ctx) error {
	title := c.Params("title")

	// Prisma client
	client := db.NewClient()
	if err := client.Connect(); err != nil {
		return c.Status(500).JSON(fiber.Map{"message": "Error connecting to the database"})
	}
	defer client.Disconnect()

	// Get the specific social
	social, err := client.Socials.FindUnique(
		db.Socials.Title.Equals(strings.ToLower(title)),
	).Exec(c.Context())
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"message": "Error fetching the social"})
	}
	if social == nil {
		return c.Status(404).JSON(fiber.Map{"message": "No social found"})
	}

	return c.Status(200).JSON(fiber.Map{"message": "Social", "data": social})
}

func SetSocial(c *fiber.Ctx) error {
	var body Social

	// Parse the request body
	if err := c.BodyParser(&body); err != nil {
		return c.Status(400).JSON(fiber.Map{"message": "Invalid input"})
	}

	// Prisma client
	client := db.NewClient()
	if err := client.Connect(); err != nil {
		return c.Status(500).JSON(fiber.Map{"message": "Error connecting to the database"})
	}
	defer client.Disconnect()

	// Create a new social
	_, err := client.Socials.UpsertOne(
		db.Socials.Title.Equals(body.Title),
	).Create(
		db.Socials.Title.Set(strings.ToLower(body.Title)),
		db.Socials.URL.Set(body.Url),
	).Update(
		db.Socials.Title.Set(strings.ToLower(body.Title)),
		db.Socials.URL.Set(body.Url),
	).Exec(c.Context())
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"message": "Error creating a new social"})
	}

	// If everything is fine, return the response
	return c.Status(200).JSON(fiber.Map{"message": "Successfully created a socail link", "data": body})
}

func DeleteSocial(c *fiber.Ctx) error {
	title := c.Params("title")

	// Prisma client
	client := db.NewClient()
	if err := client.Connect(); err != nil {
		return c.Status(500).JSON(fiber.Map{"message": "Error connecting to the database"})
	}
	defer client.Disconnect()

	// Delete the specific social
	social, err := client.Socials.FindUnique(
		db.Socials.Title.Equals(strings.ToLower(title)),
	).Delete().Exec(c.Context())
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"message": "Error deleting the social"})
	}

	if social == nil {
		return c.Status(404).JSON(fiber.Map{"message": "No social found"})
	}

	return c.Status(200).JSON(fiber.Map{"message": "Successfully deleted the social"})
}
