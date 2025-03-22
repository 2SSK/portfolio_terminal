package resumeHandler

import (
	"fmt"
	"time"

	"github.com/2SSK/portfolio_terminal/backend/config"
	"github.com/2SSK/portfolio_terminal/backend/prisma/db"
	"github.com/2SSK/portfolio_terminal/backend/utils/fileHandler"
	"github.com/gofiber/fiber/v2"
)

type Resume struct {
	File     string `json:"file"`
	PublicID string `json:"public_id"` // Optional: store for deletion
}

func GetResume(c *fiber.Ctx) error {
	userID := c.QueryInt("userId")
	if userID == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "userId is required"})
	}
	disposition := c.Query("disposition", "inline")

	resume, err := config.PrismaClient.Resume.FindUnique(
		db.Resume.UserID.Equals(userID),
	).Exec(c.Context())
	if err != nil || resume == nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "resume not found"})
	}

	// Return URL instead of serving file directly
	c.Set("Content-Disposition", fmt.Sprintf("%s; filename=%s", disposition, "resume"))
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"url": resume.File,
	})
}

func AddResume(c *fiber.Ctx) error {
	userID := c.QueryInt("userId")
	if userID == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "userId is required"})
	}

	file, err := c.FormFile("resume")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "no file uploaded"})
	}

	if err := fileHandler.ValidateFile(file, "resume"); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	url, publicID, err := fileHandler.UploadFile(file, "resume", userID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	existingResume, _ := config.PrismaClient.Resume.FindUnique(
		db.Resume.UserID.Equals(userID),
	).Exec(c.Context())

	var resume *db.ResumeModel
	if existingResume != nil {
		fileHandler.DeleteFile(fmt.Sprintf("portfolio/resume/%d_%s", userID, existingResume.File))
		resume, err = config.PrismaClient.Resume.FindUnique(
			db.Resume.UserID.Equals(userID),
		).Update(
			db.Resume.File.Set(url),
		).Exec(c.Context())
	} else {
		resume, err = config.PrismaClient.Resume.CreateOne(
			db.Resume.File.Set(url),
			db.Resume.User.Link(db.User.ID.Equals(userID)),
		).Exec(c.Context())
	}

	if err != nil {
		fileHandler.DeleteFile(publicID)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "failed to save resume", "details": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Resume uploaded successfully",
		"resume":  resume,
	})
}

func DeleteResume(c *fiber.Ctx) error {
	userID := c.QueryInt("userId")
	if userID == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "userId is required"})
	}

	resume, err := config.PrismaClient.Resume.FindUnique(
		db.Resume.UserID.Equals(userID),
	).Exec(c.Context())
	if err != nil || resume == nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "resume not found"})
	}

	publicID := fmt.Sprintf("portfolio/resume/%d_%s", userID, time.Now().Format("20060102_150405"))
	fileHandler.DeleteFile(publicID)

	_, err = config.PrismaClient.Resume.FindUnique(
		db.Resume.UserID.Equals(userID),
	).Delete().Exec(c.Context())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "failed to delete resume"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Resume deleted successfully",
	})
}
