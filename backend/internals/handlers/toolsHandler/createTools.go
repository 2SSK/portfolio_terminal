package toolsHandler

import (
	"strings"

	"github.com/2SSK/portfolio_terminal/backend/config"
	"github.com/2SSK/portfolio_terminal/backend/prisma/db"
	"github.com/gofiber/fiber/v2"
)

type input struct {
	Name string `json:"name"`
}

func AddProgrammingLang(c *fiber.Ctx) error {
	var body input
	if err := c.BodyParser(&body); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request"})
	}

	user := c.Locals("user").(*db.UserModel)
	userId := user.ID

	inputTool := strings.ToLower(body.Name)

	// Check if the input lang already exists
	isExist, err := config.PrismaClient.ProgrammingLang.FindUnique(
		db.ProgrammingLang.Lang.Equals(inputTool),
	).Exec(c.Context())
	if err == nil && isExist != nil {
		return c.Status(400).JSON(fiber.Map{"error": "This language already exists"})
	}

	// Find the Tool for this user
	tool, err := config.PrismaClient.Tools.FindFirst(
		db.Tools.UserID.Equals(userId),
	).Exec(c.Context())

	// Create new programming language
	newLang, err := config.PrismaClient.ProgrammingLang.CreateOne(
		db.ProgrammingLang.Lang.Set(inputTool),
		db.ProgrammingLang.Tools.Link(db.Tools.ID.Equals(tool.ID)),
	).Exec(c.Context())
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to create new language"})
	}

	return c.Status(201).JSON(fiber.Map{
		"lang": map[string]any{
			"id":   newLang.ID,
			"name": newLang.Lang,
		},
	})
}

func AddSoftwareTools(c *fiber.Ctx) error {
	var body input
	if err := c.BodyParser(&body); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request"})
	}

	user := c.Locals("user").(*db.UserModel)
	userId := user.ID

	inputTool := strings.ToLower(body.Name)

	// Check if the input lang already exists
	isExist, err := config.PrismaClient.SoftwareTools.FindUnique(
		db.SoftwareTools.Name.Equals(inputTool),
	).Exec(c.Context())
	if err == nil && isExist != nil {
		return c.Status(400).JSON(fiber.Map{"error": "This software tool already exists"})
	}

	// Find the Tool for this user
	tool, err := config.PrismaClient.Tools.FindFirst(
		db.Tools.UserID.Equals(userId),
	).Exec(c.Context())

	// Create new programming language
	newTool, err := config.PrismaClient.SoftwareTools.CreateOne(
		db.SoftwareTools.Name.Set(strings.ToLower(inputTool)),
		db.SoftwareTools.Tools.Link(db.Tools.ID.Equals(tool.ID)),
	).Exec(c.Context())
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to create new software tool"})
	}

	return c.Status(201).JSON(fiber.Map{
		"software": map[string]any{
			"id":   newTool.ID,
			"name": newTool.Name,
		},
	})
}

func AddFramework(c *fiber.Ctx) error {
	var body input
	if err := c.BodyParser(&body); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request"})
	}

	user := c.Locals("user").(*db.UserModel)
	userId := user.ID

	inputTool := strings.ToLower(body.Name)

	// Check if the input lang already exists
	isExist, err := config.PrismaClient.Frameworks.FindUnique(
		db.Frameworks.Name.Equals(inputTool),
	).Exec(c.Context())
	if err == nil && isExist != nil {
		return c.Status(400).JSON(fiber.Map{"error": "This framework already exists"})
	}

	// Find the Tool for this user
	tool, err := config.PrismaClient.Tools.FindFirst(
		db.Tools.UserID.Equals(userId),
	).Exec(c.Context())

	// Create new programming language
	newFramework, err := config.PrismaClient.Frameworks.CreateOne(
		db.Frameworks.Name.Set(inputTool),
		db.Frameworks.Tools.Link(db.Tools.ID.Equals(tool.ID)),
	).Exec(c.Context())
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to create new framework"})
	}

	return c.Status(201).JSON(fiber.Map{
		"lang": map[string]any{
			"id":   newFramework.ID,
			"name": newFramework.Name,
		},
	})
}

func AddDatabase(c *fiber.Ctx) error {
	var body input
	if err := c.BodyParser(&body); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request"})
	}

	user := c.Locals("user").(*db.UserModel)
	userId := user.ID

	inputTool := strings.ToLower(body.Name)

	// Check if the input lang already exists
	isExist, err := config.PrismaClient.Databases.FindUnique(
		db.Databases.Name.Equals(inputTool),
	).Exec(c.Context())
	if err == nil && isExist != nil {
		return c.Status(400).JSON(fiber.Map{"error": "This database already exists"})
	}

	// Find the Tool for this user
	tool, err := config.PrismaClient.Tools.FindFirst(
		db.Tools.UserID.Equals(userId),
	).Exec(c.Context())

	// Create new programming language
	newDb, err := config.PrismaClient.Databases.CreateOne(
		db.Databases.Name.Set(inputTool),
		db.Databases.Tools.Link(db.Tools.ID.Equals(tool.ID)),
	).Exec(c.Context())
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to create new database"})
	}

	return c.Status(201).JSON(fiber.Map{
		"lang": map[string]any{
			"id":   newDb.ID,
			"name": newDb.Name,
		},
	})
}
