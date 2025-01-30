package toolHandler

import (
	"strings"

	"github.com/2SSK/portfolio_terminal/backend/prisma/db"
	"github.com/gofiber/fiber/v2"
)

type ToolStruct struct {
	ToolName string `json:"toolName"`
}

func GetTools(c *fiber.Ctx) error {
	// prisma client
	client := db.NewClient()
	if err := client.Connect(); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Could not connect to the database"})
	}
	defer client.Disconnect()

	// Query all the tools from the database
	tools, err := client.Tools.FindMany().Exec(c.Context())
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Could not fetch the tools"})
	}

	if len(tools) == 0 {
		return c.Status(200).JSON(fiber.Map{"error": "No tools found", "tools": tools})
	}

	// If everything is fine, return the tools
	return c.Status(200).JSON(fiber.Map{"message": "Tools fetched successfully", "tools": tools})
}

func SetTools(c *fiber.Ctx) error {
	var body ToolStruct

	// Parse the request body
	if err := c.BodyParser(&body); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid Input"})
	}

	// Convert the tool name to lowercase
	input := strings.ToLower(body.ToolName)

	// Prisma client
	client := db.NewClient()
	if err := client.Connect(); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Could not connect to the database"})
	}
	defer client.Disconnect()

	// Check if the tool already exists
	existingTool, _ := client.Tools.FindUnique(
		db.Tools.ToolName.Equals(input),
	).Exec(c.Context())
	if existingTool != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Tool already exists"})
	}

	// Create the tool
	_, err := client.Tools.CreateOne(
		db.Tools.ToolName.Set(input),
	).Exec(c.Context())
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Error while creating the tool"})
	}

	// If everthing is fine, return a success message
	return c.Status(200).JSON(fiber.Map{"message": "Tool Added successfully"})
}

func DeleteTools(c *fiber.Ctx) error {
	toolName := c.Params("tool")

	// Normalize tool name
	input := strings.ToLower(toolName)

	client := db.NewClient()
	if err := client.Connect(); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Could not connect to the database"})
	}
	defer client.Disconnect()

	// Delete the tool
	_, err := client.Tools.FindUnique(
		db.Tools.ToolName.Equals(input),
	).Delete().Exec(c.Context())

	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Error while deleting the tool", "err": err.Error()})
	}

	return c.Status(200).JSON(fiber.Map{"message": "Tool deleted successfully"})
}
