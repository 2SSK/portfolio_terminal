package experienceHandler

import (
	"strings"

	"github.com/2SSK/portfolio_terminal/backend/config"
	"github.com/2SSK/portfolio_terminal/backend/prisma/db"
	"github.com/gofiber/fiber/v2"
)

type ExperienceRequest struct {
	Company     string `json:"company"`
	CompanyURL  string `json:"company_url"`
	Position    string `json:"position"`
	StartDate   string `json:"start_date"`
	EndDate     string `json:"end_date"`
	Description string `json:"description"`
}

func AddExperience(c *fiber.Ctx) error {
	user := c.Locals("user").(*db.UserModel)
	userId := user.ID

	experience := new(ExperienceRequest)
	if err := c.BodyParser(experience); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	// Check if the experience already exists
	isExist, err := config.PrismaClient.Experience.FindUnique(
		db.Experience.Company.Equals(strings.ToLower(experience.Company)),
	).Exec(c.Context())
	if err == nil && isExist != nil {
		return c.Status(400).JSON(fiber.Map{"error": "This experience already exists"})
	}

	newExperience, err := config.PrismaClient.Experience.CreateOne(
		db.Experience.Company.Set(strings.ToLower(experience.Company)),
		db.Experience.CompanyURL.Set(experience.CompanyURL),
		db.Experience.Position.Set(experience.Position),
		db.Experience.StartDate.Set(experience.StartDate),
		db.Experience.EndDate.Set(experience.EndDate),
		db.Experience.Description.Set(experience.Description),
		db.Experience.User.Link(db.User.ID.Equals(userId)),
	).Exec(c.Context())
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to create new experience"})
	}

	return c.Status(201).JSON(fiber.Map{"experience": newExperience})

}

func GetAllExperience(c *fiber.Ctx) error {
	user := c.Locals("user").(*db.UserModel)
	userId := user.ID

	experiences, err := config.PrismaClient.Experience.FindMany(
		db.Experience.UserID.Equals(userId),
	).Exec(c.Context())
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to fetch experiences"})
	}

	return c.Status(200).JSON(fiber.Map{"experiences": experiences})
}

func GetExperience(c *fiber.Ctx) error {
	user := c.Locals("user").(*db.UserModel)
	userId := user.ID

	experienceId, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid experience ID"})
	}

	experience, err := config.PrismaClient.Experience.FindFirst(
		db.Experience.ID.Equals(experienceId),
		db.Experience.UserID.Equals(userId),
	).Exec(c.Context())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   "Failed to get experience",
			"details": err.Error(),
		})
	}

	if experience == nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Experience not found"})
	}

	return c.JSON(fiber.Map{"experience": experience})
}

func UpdateExperience(c *fiber.Ctx) error {
	user := c.Locals("user").(*db.UserModel)
	userId := user.ID

	experienceId, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid experience ID"})
	}

	existing, err := config.PrismaClient.Experience.FindFirst(
		db.Experience.ID.Equals(experienceId),
		db.Experience.UserID.Equals(userId),
	).Exec(c.Context())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to get experience", "details": err.Error()})
	}
	if existing == nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Experience not found or doesn't belong to user"})
	}

	experience := new(ExperienceRequest)
	if err := c.BodyParser(experience); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request format"})
	}

	updatedExperience, err := config.PrismaClient.Experience.FindUnique(
		db.Experience.ID.Equals(experienceId),
	).Update(
		db.Experience.Company.Set(strings.ToLower(experience.Company)),
		db.Experience.CompanyURL.Set(experience.CompanyURL),
		db.Experience.Position.Set(experience.Position),
		db.Experience.StartDate.Set(experience.StartDate),
		db.Experience.EndDate.Set(experience.EndDate),
		db.Experience.Description.Set(experience.Description),
	).Exec(c.Context())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to update experience", "details": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message":    "Experience updated successfully",
		"experience": updatedExperience,
	})
}

func DeleteExperience(c *fiber.Ctx) error {
	user := c.Locals("user").(*db.UserModel)
	userId := user.ID

	experienceId, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid experience ID"})
	}

	existing, err := config.PrismaClient.Experience.FindFirst(
		db.Experience.ID.Equals(experienceId),
		db.Experience.UserID.Equals(userId),
	).Exec(c.Context())
	if err != nil || existing == nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Experience not found or doesn't belong to user"})
	}

	deleted, err := config.PrismaClient.Experience.FindUnique(
		db.Experience.ID.Equals(experienceId),
	).Delete().Exec(c.Context())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to delete experience"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message":    "Experience deleted successfully",
		"experience": deleted,
	})
}
