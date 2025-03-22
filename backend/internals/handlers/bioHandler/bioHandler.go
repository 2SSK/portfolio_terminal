package bioHandler

import (
	"fmt"
	"path/filepath"

	"github.com/2SSK/portfolio_terminal/backend/config"
	"github.com/2SSK/portfolio_terminal/backend/prisma/db"
	"github.com/2SSK/portfolio_terminal/backend/utils/fileHandler"
	"github.com/gofiber/fiber/v2"
)

type BioRequest struct {
	Image       string `json:"image"`
	Name        string `json:"name"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

func AddBio(c *fiber.Ctx) error {
	// Parse userId from query parameter
	userId := c.QueryInt("userId")

	// Parse multipart form data
	form, err := c.MultipartForm()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Failed to parse form data"})
	}

	// Extract and validate required fields
	bio := new(BioRequest)
	if titles := form.Value["title"]; len(titles) == 0 || titles[0] == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Title is required"})
	} else {
		bio.Title = titles[0]
	}

	// Extract optional fields with defaults
	if names := form.Value["name"]; len(names) > 0 {
		bio.Name = names[0]
	}
	if titles := form.Value["title"]; len(titles) > 0 {
		bio.Title = titles[0]
	}
	if descriptions := form.Value["description"]; len(descriptions) > 0 {
		bio.Description = descriptions[0]
	}

	// Extract and validate preview file
	images, ok := form.File["image"]
	if !ok || len(images) == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Image file is required"})
	}
	image := images[0] // Only process the first file

	// Generate unique filename
	ext := filepath.Ext(image.Filename)
	filename := fmt.Sprintf("%d_%s_dp%s", userId, bio.Name, ext)
	filePath := fileHandler.GetFilePath(filename, "bio")

	// Validate and save the file
	if err := fileHandler.ValidateFile(image, "bio"); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	if err := c.SaveFile(image, filePath); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to save dp image"})
	}
	bio.Image = filename

	// Create project in the database
	newBio, err := config.PrismaClient.Bio.UpsertOne(
		db.Bio.UserID.Equals(userId),
	).Create(
		db.Bio.Image.Set(bio.Image),
		db.Bio.Name.Set(bio.Name),
		db.Bio.Title.Set(bio.Title),
		db.Bio.Description.Set(bio.Description),
		db.Bio.User.Link(
			db.User.ID.Equals(userId),
		),
	).Update(
		db.Bio.Image.Set(bio.Image),
		db.Bio.Name.Set(bio.Name),
		db.Bio.Title.Set(bio.Title),
		db.Bio.Description.Set(bio.Description),
	).Exec(c.Context())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   "Failed to Update bio",
			"details": err.Error(),
		})
	}

	// Return success response
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Bio updated successfully",
		"project": newBio,
	})
}

func GetBio(c *fiber.Ctx) error {
	userId := c.QueryInt("userId")

	bio, err := config.PrismaClient.Bio.FindUnique(
		db.Bio.UserID.Equals(userId),
	).Exec(c.Context())

	if err != nil {
		if err == db.ErrNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Bio not found"})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   "Failed to get bio",
			"details": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Bio fetched successfully",
		"bio":     bio,
	})
}
