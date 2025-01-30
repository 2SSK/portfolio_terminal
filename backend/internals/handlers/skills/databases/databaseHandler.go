package databaseHandler

import (
	"strings"

	"github.com/2SSK/portfolio_terminal/backend/prisma/db"
	"github.com/gofiber/fiber/v2"
)

type dbStruct struct {
	DatabaseName string `json:"databaseName"`
	SkillId      int    `json:"skillId"`
}

func GetDb(c *fiber.Ctx) error {
	// prisma client
	client := db.NewClient()
	if err := client.Connect(); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Could not connect to the database"})
	}
	defer client.Disconnect()

	// Query all the tools from the database
	database, err := client.Database.FindMany().Exec(c.Context())
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Could not fetch the database"})
	}

	if len(database) == 0 {
		return c.Status(200).JSON(fiber.Map{"error": "No database found", "db": database})
	}

	// If everything is fine, return the tools
	return c.Status(200).JSON(fiber.Map{"message": "DBs fetched successfully", "db": database})
}

func SetDb(c *fiber.Ctx) error {
	var body dbStruct

	// Parse the request body
	if err := c.BodyParser(&body); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid Input"})
	}

	// Prisma client
	client := db.NewClient()
	if err := client.Connect(); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Could not connect to the database"})
	}
	defer client.Disconnect()

	// Check if the database already exists
	existingDb, _ := client.Database.FindUnique(
		db.Database.DatabaseName.Equals(strings.ToLower(body.DatabaseName)),
	).Exec(c.Context())
	if existingDb != nil {
		return c.Status(400).JSON(fiber.Map{"error": "DB already exists"})
	}

	// Create the database
	_, err := client.Database.CreateOne(
		db.Database.DatabaseName.Set(strings.ToLower(body.DatabaseName)),
		// db.Database.Skill.Link(
		// 	db.Skill.ID.Equals(body.SkillId),
		// ),
	).Exec(c.Context())
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Error while creating the DB"})
	}

	// If everthing is fine, return a success message
	return c.Status(200).JSON(fiber.Map{"message": "DB Added successfully"})
}

func DeleteDb(c *fiber.Ctx) error {
	Db := c.Params("db")

	// Normalize tool name
	input := strings.ToLower(Db)

	client := db.NewClient()
	if err := client.Connect(); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Could not connect to the database"})
	}
	defer client.Disconnect()

	// Delete the tool
	_, err := client.Database.FindUnique(
		db.Database.DatabaseName.Equals(input),
	).Delete().Exec(c.Context())

	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Error while deleting the DB", "err": err.Error()})
	}

	return c.Status(200).JSON(fiber.Map{"message": "Db deleted successfully"})
}
