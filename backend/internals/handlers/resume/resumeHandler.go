package resumeHandler

import (
	"os"
	"path/filepath"

	"github.com/2SSK/portfolio_terminal/backend/prisma/db"
	"github.com/gofiber/fiber/v2"
)

type ResumeStruct struct {
	Resume string `json:"resume"`
}

func GetResume(c *fiber.Ctx) error {
	client := db.NewClient()
	if err := client.Connect(); err != nil {
		return c.Status(500).JSON(fiber.Map{"message": "Error connecting to the database"})
	}
	defer client.Disconnect()

	resume, err := client.Resume.FindFirst().Exec(c.Context())
	if resume == nil {
		return c.Status(404).JSON(fiber.Map{"message": "No resume found"})
	}
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"message": "Error fetching resume"})
	}

	// If everything is fine, return the resume
	return c.Status(200).JSON(fiber.Map{"message": "Resume", "data": resume})
}

func SetResume(c *fiber.Ctx) error {
	// Parse the forma data
	body := new(ResumeStruct)
	body.Resume = c.FormValue("Resume")

	// Prisma client
	client := db.NewClient()
	if err := client.Connect(); err != nil {
		return c.Status(500).JSON(fiber.Map{"message": "Error connecting to the database"})
	}
	defer client.Disconnect()

	// Define the upload directory and Resume subdirectory
	uploadDir := "./uploads"
	resumeDir := filepath.Join(uploadDir, "resume")

	// Ensure the `uploads/Resume/` directory exists
	if _, err := os.Stat(resumeDir); os.IsNotExist(err) {
		os.MkdirAll(resumeDir, 0755)
	} else {
		// Removes all files inside of "uploads/Resume/" before uploading a new one
		files, errr := os.ReadDir(resumeDir)
		if errr == nil {
			for _, file := range files {
				os.Remove(filepath.Join(resumeDir, file.Name()))
			}
		}
	}

	// Handle file upload (Resume)
	file, err := c.FormFile("Resume")
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"message": "Error uploading resume"})
	} else {
		filename := file.Filename
		savePath := filepath.Join(resumeDir, filename)

		if err := c.SaveFile(file, savePath); err != nil {
			return c.Status(500).JSON(fiber.Map{"message": "Error saving the resume"})
		}
		// Create a new SetResume
		body.Resume = filename
	}

	_, err = client.Resume.UpsertOne(
		db.Resume.ID.Equals(1),
	).Create(
		db.Resume.ResumeFile.Set(body.Resume),
	).Update(
		db.Resume.ResumeFile.Set(body.Resume),
	).Exec(c.Context())

	if err != nil {
		return c.Status(500).JSON(fiber.Map{"message": "Could not update Resume"})
	}

	// If everything is fine, return the resume
	return c.Status(200).JSON(fiber.Map{"message": "Resume set"})
}

func DeleteResume(c *fiber.Ctx) error {
	// Prisma client
	client := db.NewClient()
	if err := client.Connect(); err != nil {
		return c.Status(500).JSON(fiber.Map{"message": "Error connecting to the database"})
	}
	defer client.Disconnect()

	// Query the Database
	resume, err := client.Resume.FindFirst().Exec(c.Context())
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"message": "Error fetching resume"})
	}

	// Delete the resume
	_, err = client.Resume.FindUnique(
		db.Resume.ID.Equals(resume.ID),
	).Delete().Exec(c.Context())
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"message": "Error deleting resume"})
	}

	return c.Status(200).JSON(fiber.Map{"message": "Resume deleted"})
}
