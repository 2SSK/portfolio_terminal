package resumeHandler

import (
	"fmt"
	"path/filepath"
	"time"

	"github.com/2SSK/portfolio_terminal/backend/config"
	"github.com/2SSK/portfolio_terminal/backend/prisma/db"
	"github.com/2SSK/portfolio_terminal/backend/utils/fileHandler"
	"github.com/gofiber/fiber/v2"
)

func GetResume(c *fiber.Ctx) error {
	userId := c.QueryInt("userId")
	disposition := c.Query("disposition", "inline") // default to inline viewing

	// Get resume record
	resume, err := config.PrismaClient.Resume.FindUnique(
		db.Resume.UserID.Equals(userId),
	).Exec(c.Context())
	if err != nil || resume == nil {
		return c.Status(404).JSON(fiber.Map{"error": "Resume not found"})
	}

	// Get file path and send file
	filePath := fileHandler.GetFilePath(resume.File, "resume")
	ext := filepath.Ext(resume.File)

	// Set content type based on file extension
	contentType := "application/pdf"
	if ext == ".doc" || ext == ".docx" {
		contentType = "application/msword"
	}

	c.Set("Content-Type", contentType)
	c.Set("Content-Disposition", fmt.Sprintf("%s; filename=%s", disposition, resume.File))

	return c.SendFile(filePath, true)
}

func AddResume(c *fiber.Ctx) error {
	userId := c.QueryInt("userId")

	// Get and validate file
	file, err := c.FormFile("resume")
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "No file uploaded"})
	}

	if err := fileHandler.ValidateFile(file, "resume"); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	// Handle existing resume
	existingResume, _ := config.PrismaClient.Resume.FindUnique(
		db.Resume.UserID.Equals(userId),
	).Exec(c.Context())

	if existingResume != nil {
		if err := fileHandler.DeleteFile(existingResume.File, "resume"); err != nil {
			return c.Status(500).JSON(fiber.Map{"error": "Failed to delete existing resume"})
		}
	}

	// Save new file
	filename := fmt.Sprintf("%d_%s%s", userId, time.Now().Format("20060102_150405"),
		filepath.Ext(file.Filename))
	filePath := fileHandler.GetFilePath(filename, "resume")

	if err := c.SaveFile(file, filePath); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to save file"})
	}

	// Update database
	var resume *db.ResumeModel
	if existingResume != nil {
		resume, err = config.PrismaClient.Resume.FindUnique(
			db.Resume.UserID.Equals(userId),
		).Update(
			db.Resume.File.Set(filename),
		).Exec(c.Context())
	} else {
		resume, err = config.PrismaClient.Resume.CreateOne(
			db.Resume.File.Set(filename),
			db.Resume.User.Link(
				db.User.ID.Equals(userId),
			),
		).Exec(c.Context())
	}

	if err != nil {
		fileHandler.DeleteFile(filename, "resume") // Clean up file if database operation fails
		return c.Status(500).JSON(fiber.Map{"error": "Failed to save resume information"})
	}

	return c.Status(201).JSON(fiber.Map{
		"message": "Resume uploaded successfully",
		"resume":  resume,
	})
}

func DeleteResume(c *fiber.Ctx) error {
	userId := c.QueryInt("userId")

	resume, err := config.PrismaClient.Resume.FindUnique(
		db.Resume.UserID.Equals(userId),
	).Exec(c.Context())
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Resume not found"})
	}

	// Delete file first
	if err := fileHandler.DeleteFile(resume.File, "resume"); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to delete file"})
	}

	// Delete from database
	_, err = config.PrismaClient.Resume.FindUnique(
		db.Resume.UserID.Equals(userId),
	).Delete().Exec(c.Context())
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to delete resume record"})
	}

	return c.Status(200).JSON(fiber.Map{
		"message": "Resume deleted successfully",
	})
}
