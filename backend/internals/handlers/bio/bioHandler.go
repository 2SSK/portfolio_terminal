package bioHandler

import (
	"os"
	"path/filepath"

	"github.com/2SSK/portfolio_terminal/backend/prisma/db"
	"github.com/gofiber/fiber/v2"
)

type BioStruct struct {
	Dp          string `json:"dp"`
	Name        string `json:"name"`
	Intro       string `json:"intro"`
	Description string `json:"description"`
}

func GetBio(c *fiber.Ctx) error {
	// Prisma client
	client := db.NewClient()
	if err := client.Connect(); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Could not connect to the database"})
	}
	defer client.Disconnect()

	// Query the Database
	bio, err := client.Bio.FindFirst().Exec(c.Context())
	if err != nil || bio == nil {
		return c.Status(404).JSON(fiber.Map{"error": "Bio records not found"})
	}

	// If everything is fine, return the bio
	return c.Status(200).JSON(fiber.Map{"bio": bio})
}

func SetBio(c *fiber.Ctx) error {
	// Parse the forma data
	body := new(BioStruct)
	body.Name = c.FormValue("Name")
	body.Intro = c.FormValue("Intro")
	body.Description = c.FormValue("Description")

	// Define the upload directory and Dp subdirectory
	uploadDir := "./uploads"
	dpDir := filepath.Join(uploadDir, "dp")

	// Ensure the `uploads/Dp/` directory exists
	if _, err := os.Stat(dpDir); os.IsNotExist(err) {
		os.MkdirAll(dpDir, 0755)
	} else {
		// Removes all files inside of "uploads/Dp/" before uploading a new one
		files, errr := os.ReadDir(dpDir)
		if errr == nil {
			for _, file := range files {
				os.Remove(filepath.Join(dpDir, file.Name()))
			}
		}
	}

	// Handle file upload (Dp - Display Picture)
	file, err := c.FormFile("Dp")
	if err == nil {
		filename := file.Filename
		savePath := filepath.Join(dpDir, filename)

		if err := c.SaveFile(file, savePath); err != nil {
			return c.Status(500).JSON(fiber.Map{"error": "Could not save the image"})
		}
		body.Dp = filename
	}

	// Prisma client
	client := db.NewClient()
	if err := client.Connect(); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Could not connect to the database"})
	}
	defer client.Disconnect()

	// Update the Bio records
	_, err = client.Bio.UpsertOne(
		// query
		db.Bio.ID.Equals(1),
	).Create(
		// set these fields if bio doesn't exist already
		db.Bio.Dp.Set(body.Dp),
		db.Bio.Name.Set(body.Name),
		db.Bio.Intro.Set(body.Intro),
		db.Bio.Description.Set(body.Description),
	).Update(
		db.Bio.Dp.Set(body.Dp),
		db.Bio.Name.Set(body.Name),
		db.Bio.Intro.Set(body.Intro),
		db.Bio.Description.Set(body.Description),
	).Exec(c.Context())

	if err != nil {
		return c.Status(500).JSON(fiber.Map{"message": "Could not update Bio"})
	}

	return c.Status(200).JSON(fiber.Map{"message": "Bio updated successfully"})
}
