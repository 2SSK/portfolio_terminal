package toolsHandler

import (
	"github.com/2SSK/portfolio_terminal/backend/config"
	"github.com/2SSK/portfolio_terminal/backend/prisma/db"
	"github.com/gofiber/fiber/v2"
)

func GetAllTools(c *fiber.Ctx) error {
	user := c.Locals("user").(*db.UserModel)
	userId := user.ID

	tools, err := config.PrismaClient.Tools.FindUnique(
		db.Tools.UserID.Equals(userId),
	).With(
		db.Tools.ProgrammingLang.Fetch(),
		db.Tools.SoftwareTools.Fetch(),
		db.Tools.Frameworks.Fetch(),
		db.Tools.Databases.Fetch(),
	).Exec(c.Context())

	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Failed to fetch programming languages",
		})
	}

	if tools == nil {
		return c.Status(404).JSON(fiber.Map{"error": "No tools found for this user"})
	}

	return c.Status(200).JSON(fiber.Map{"Tools": tools})
}

func GetAllProgrammingLangs(c *fiber.Ctx) error {
	user := c.Locals("user").(*db.UserModel)
	userId := user.ID

	tools, err := config.PrismaClient.Tools.FindMany(
		db.Tools.UserID.Equals(userId),
	).With(
		db.Tools.ProgrammingLang.Fetch(),
	).Exec(c.Context())
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Failed to fetch programming languages",
		})
	}
	if len(tools) == 0 {
		return c.Status(404).JSON(fiber.Map{"error": "No tools found for this user"})
	}

	// Extract programming languages from the tools
	var programmingLangs []string
	for _, tool := range tools {
		for _, lang := range tool.ProgrammingLang() {
			programmingLangs = append(programmingLangs, lang.Lang)
		}
	}

	return c.Status(200).JSON(fiber.Map{"programmingLanguages": programmingLangs})
}

func GetAllSoftwareTools(c *fiber.Ctx) error {
	user := c.Locals("user").(*db.UserModel)
	userId := user.ID

	tools, err := config.PrismaClient.Tools.FindMany(
		db.Tools.UserID.Equals(userId),
	).With(
		db.Tools.SoftwareTools.Fetch(),
	).Exec(c.Context())
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Failed to fetch software tools",
		})
	}
	if len(tools) == 0 {
		return c.Status(404).JSON(fiber.Map{"error": "No tools found for this user"})
	}

	// Extract programming languages from the tools
	var softwareTools []string
	for _, tool := range tools {
		for _, application := range tool.SoftwareTools() {
			softwareTools = append(softwareTools, application.Name)
		}
	}

	return c.Status(200).JSON(fiber.Map{"softwareTools": softwareTools})
}

func GetAllFrameworks(c *fiber.Ctx) error {
	user := c.Locals("user").(*db.UserModel)
	userId := user.ID

	tools, err := config.PrismaClient.Tools.FindMany(
		db.Tools.UserID.Equals(userId),
	).With(
		db.Tools.Frameworks.Fetch(),
	).Exec(c.Context())
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Failed to fetch frameworks",
		})
	}
	if len(tools) == 0 {
		return c.Status(404).JSON(fiber.Map{"error": "No frameworks found for this user"})
	}

	// Extract programming languages from the tools
	var frameworks []string
	for _, tool := range tools {
		for _, framework := range tool.Frameworks() {
			frameworks = append(frameworks, framework.Name)
		}
	}

	return c.Status(200).JSON(fiber.Map{"frameworks": frameworks})
}

func GetAllDatabases(c *fiber.Ctx) error {
	user := c.Locals("user").(*db.UserModel)
	userId := user.ID

	tools, err := config.PrismaClient.Tools.FindMany(
		db.Tools.UserID.Equals(userId),
	).With(
		db.Tools.Databases.Fetch(),
	).Exec(c.Context())
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error":   "Failed to fetch databases",
			"details": err.Error(),
		})
	}
	if len(tools) == 0 {
		return c.Status(404).JSON(fiber.Map{"error": "No database found for this user"})
	}

	// Extract programming languages from the tools
	var databases []string
	for _, tool := range tools {
		for _, db := range tool.Databases() {
			databases = append(databases, db.Name)
		}
	}

	return c.Status(200).JSON(fiber.Map{"databases": databases})
}
