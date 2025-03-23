package toolsHandler

import (
	"strings"

	"github.com/2SSK/portfolio_terminal/backend/config"
	"github.com/2SSK/portfolio_terminal/backend/prisma/db"
	"github.com/gofiber/fiber/v2"
)

type DeleteRequest struct {
	Name string `json:"name"`
}

func DeleteProgrammingLang(c *fiber.Ctx) error {
	user := c.Locals("user").(*db.UserModel)
	userId := user.ID

	var req DeleteRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error":   "Failed to parse request body",
			"details": err.Error(),
		})
	}

	req.Name = strings.TrimSpace(req.Name)
	if req.Name == "" {
		return c.Status(400).JSON(fiber.Map{"error": "Name parameter is required"})
	}

	// Convert the name to lowercase
	name := strings.ToLower(req.Name)

	// Find the programming language to delete
	lang, err := config.PrismaClient.ProgrammingLang.FindUnique(
		db.ProgrammingLang.Lang.Equals(name),
	).With(
		db.ProgrammingLang.Tools.Fetch().Take(1),
	).Exec(c.Context())

	// Check if the language exists
	if err != nil {
		if err == db.ErrNotFound {
			return c.Status(404).JSON(fiber.Map{"error": "Programming language not found"})
		}
		return c.Status(500).JSON(fiber.Map{
			"error":   "Failed to check programming language",
			"details": err.Error(),
		})
	}

	// Check if the user is authorized to delete the language
	if len(lang.Tools()) > 0 && lang.Tools()[0].UserID != userId {
		return c.Status(403).JSON(fiber.Map{"error": "You are not authorized to delete this language"})
	}

	// Delete the programming language
	_, err = config.PrismaClient.ProgrammingLang.FindUnique(
		db.ProgrammingLang.Lang.Equals(name),
	).Delete().Exec(c.Context())

	// Check if the language was deleted
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error":   "Failed to delete programming language",
			"details": err.Error(),
		})
	}

	// Return a success message
	return c.Status(200).JSON(fiber.Map{"message": "successfully deleted lang"})
}

func DeleteSoftwareTool(c *fiber.Ctx) error {
	user := c.Locals("user").(*db.UserModel)
	userId := user.ID

	var req DeleteRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error":   "Failed to parse request body",
			"details": err.Error(),
		})
	}

	req.Name = strings.TrimSpace(req.Name)
	if req.Name == "" {
		return c.Status(400).JSON(fiber.Map{"error": "Name parameter is required"})
	}

	// Convert the name to lowercase
	name := strings.ToLower(req.Name)

	// Find the programming language to delete
	lang, err := config.PrismaClient.SoftwareTools.FindUnique(
		db.SoftwareTools.Name.Equals(name),
	).With(
		db.SoftwareTools.Tools.Fetch().Take(1),
	).Exec(c.Context())

	// Check if the language exists
	if err != nil {
		if err == db.ErrNotFound {
			return c.Status(404).JSON(fiber.Map{"error": "software tool not found"})
		}
		return c.Status(500).JSON(fiber.Map{
			"error":   "Failed to check software tool",
			"details": err.Error(),
		})
	}

	// Check if the user is authorized to delete the language
	if len(lang.Tools()) > 0 && lang.Tools()[0].UserID != userId {
		return c.Status(403).JSON(fiber.Map{"error": "You are not authorized to delete this software tool"})
	}

	// Delete the programming language
	_, err = config.PrismaClient.SoftwareTools.FindUnique(
		db.SoftwareTools.Name.Equals(name),
	).Delete().Exec(c.Context())

	// Check if the language was deleted
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error":   "Failed to delete software tool",
			"details": err.Error(),
		})
	}

	// Return a success message
	return c.Status(200).JSON(fiber.Map{"message": "successfully deleted software tool"})
}

func DeleteFramework(c *fiber.Ctx) error {
	user := c.Locals("user").(*db.UserModel)
	userId := user.ID

	var req DeleteRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error":   "Failed to parse request body",
			"details": err.Error(),
		})
	}

	req.Name = strings.TrimSpace(req.Name)
	if req.Name == "" {
		return c.Status(400).JSON(fiber.Map{"error": "Name parameter is required"})
	}

	// Convert the name to lowercase
	name := strings.ToLower(req.Name)

	// Find the programming language to delete
	lang, err := config.PrismaClient.Frameworks.FindUnique(
		db.Frameworks.Name.Equals(name),
	).With(
		db.Frameworks.Tools.Fetch().Take(1),
	).Exec(c.Context())

	// Check if the language exists
	if err != nil {
		if err == db.ErrNotFound {
			return c.Status(404).JSON(fiber.Map{"error": "framework not found"})
		}
		return c.Status(500).JSON(fiber.Map{
			"error":   "Failed to check framwork",
			"details": err.Error(),
		})
	}

	// Check if the user is authorized to delete the language
	if len(lang.Tools()) > 0 && lang.Tools()[0].UserID != userId {
		return c.Status(403).JSON(fiber.Map{"error": "You are not authorized to delete this framework"})
	}

	// Delete the programming language
	_, err = config.PrismaClient.Frameworks.FindUnique(
		db.Frameworks.Name.Equals(name),
	).Delete().Exec(c.Context())

	// Check if the language was deleted
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error":   "Failed to delete framework",
			"details": err.Error(),
		})
	}

	// Return a success message
	return c.Status(200).JSON(fiber.Map{"message": "successfully deleted framework"})
}

func DeleteDatabase(c *fiber.Ctx) error {
	user := c.Locals("user").(*db.UserModel)
	userId := user.ID

	var req DeleteRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error":   "Failed to parse request body",
			"details": err.Error(),
		})
	}

	req.Name = strings.TrimSpace(req.Name)
	if req.Name == "" {
		return c.Status(400).JSON(fiber.Map{"error": "Name parameter is required"})
	}

	// Convert the name to lowercase
	name := strings.ToLower(req.Name)

	// Find the programming language to delete
	lang, err := config.PrismaClient.Databases.FindUnique(
		db.Databases.Name.Equals(name),
	).With(
		db.Databases.Tools.Fetch().Take(1),
	).Exec(c.Context())

	// Check if the language exists
	if err != nil {
		if err == db.ErrNotFound {
			return c.Status(404).JSON(fiber.Map{"error": "database not found"})
		}
		return c.Status(500).JSON(fiber.Map{
			"error":   "Failed to check database",
			"details": err.Error(),
		})
	}

	// Check if the user is authorized to delete the language
	if len(lang.Tools()) > 0 && lang.Tools()[0].UserID != userId {
		return c.Status(403).JSON(fiber.Map{"error": "You are not authorized to delete this database"})
	}

	// Delete the programming language
	_, err = config.PrismaClient.Databases.FindUnique(
		db.Databases.Name.Equals(name),
	).Delete().Exec(c.Context())

	// Check if the language was deleted
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error":   "Failed to delete database",
			"details": err.Error(),
		})
	}

	// Return a success message
	return c.Status(200).JSON(fiber.Map{"message": "successfully deleted database"})
}
