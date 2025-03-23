package linkHandler

import (
	"github.com/2SSK/portfolio_terminal/backend/config"
	"github.com/2SSK/portfolio_terminal/backend/prisma/db"
	"github.com/gofiber/fiber/v2"
)

type input struct {
	Title  string `json:"title" validate:"required"`
	Url    string `json:"url" validate:"required,url"`
	UserId int    `json:"userId" validate:"required"`
}

// GetAllLinks gets all links for a specific user
func GetAllLinks(c *fiber.Ctx) error {
	user := c.Locals("user").(*db.UserModel)
	userId := user.ID

	links, err := config.PrismaClient.Links.FindMany(
		db.Links.UserID.Equals(userId),
	).Exec(c.Context())
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Unable to fetch links"})
	}

	return c.Status(200).JSON(links)
}

// GetLink gets a specific link and verifies it belongs to the user
func GetLink(c *fiber.Ctx) error {
	user := c.Locals("user").(*db.UserModel)
	userId := user.ID

	linkId, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid link ID"})
	}

	link, err := config.PrismaClient.Links.FindFirst(
		db.Links.ID.Equals(linkId),
		db.Links.UserID.Equals(userId),
	).Exec(c.Context())
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Link not found"})
	}

	return c.Status(200).JSON(link)
}

// CreateLink creates a new link for a user
func CreateLink(c *fiber.Ctx) error {
	var body input

	if err := c.BodyParser(&body); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request"})
	}

	// Check if title already exists for this user
	existingLink, _ := config.PrismaClient.Links.FindFirst(
		db.Links.Title.Equals(body.Title),
		db.Links.UserID.Equals(body.UserId),
	).Exec(c.Context())
	if existingLink != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Link title already exists for this user"})
	}

	link, err := config.PrismaClient.Links.CreateOne(
		db.Links.Title.Set(body.Title),
		db.Links.URL.Set(body.Url),
		db.Links.User.Link(
			db.User.ID.Equals(body.UserId),
		),
	).Exec(c.Context())

	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Failed to create link"})
	}

	return c.Status(201).JSON(link)
}

func UpdateLink(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid ID"})
	}

	var body struct {
		Title string `json:"title"`
		Url   string `json:"url"`
	}

	if err := c.BodyParser(&body); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request body"})
	}

	// First check if link exists
	existingLink, err := config.PrismaClient.Links.FindUnique(
		db.Links.ID.Equals(id),
	).Exec(c.Context())
	if err != nil || existingLink == nil {
		return c.Status(404).JSON(fiber.Map{"error": "Link not found"})
	}

	// Create update params
	var updateParams []db.LinksSetParam

	// Only add fields that are provided
	if body.Title != "" {
		updateParams = append(updateParams, db.Links.Title.Set(body.Title))
	}
	if body.Url != "" {
		updateParams = append(updateParams, db.Links.URL.Set(body.Url))
	}

	// Update the link
	link, err := config.PrismaClient.Links.FindUnique(
		db.Links.ID.Equals(id),
	).Update(
		updateParams...,
	).Exec(c.Context())

	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to update Link"})
	}

	return c.Status(200).JSON(link)
}

// DeleteLink deletes a specific link if it belongs to the user
func DeleteLink(c *fiber.Ctx) error {
	user := c.Locals("user").(*db.UserModel)
	userId := user.ID

	linkId, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid link ID"})
	}

	// Verify link exists and belongs to user
	existingLink, err := config.PrismaClient.Links.FindFirst(
		db.Links.ID.Equals(linkId),
		db.Links.UserID.Equals(userId),
	).Exec(c.Context())
	if err != nil || existingLink == nil {
		return c.Status(404).JSON(fiber.Map{"error": "Link not found or doesn't belong to user"})
	}

	// Delete the link
	link, err := config.PrismaClient.Links.FindUnique(
		db.Links.ID.Equals(linkId),
	).Delete().Exec(c.Context())

	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to delete link"})
	}

	return c.Status(200).JSON(fiber.Map{
		"message": "Link deleted successfully",
		"link":    link,
	})
}
