package contactHandler

import (
	"strings"

	"github.com/2SSK/portfolio_terminal/backend/prisma/db"
	"github.com/gofiber/fiber/v2"
)

type contactStruct struct {
	ContactType string `json:"contactType"`
	ContactData string `json:"contactData"`
}

func GetContact(c *fiber.Ctx) error {
	client := db.NewClient()
	if err := client.Connect(); err != nil {
		return c.Status(500).JSON(fiber.Map{"message": "Error connecting to the database"})
	}
	defer client.Disconnect()

	contacts, err := client.Contacts.FindMany().Exec(c.Context())
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"message": "Error fetching contacts"})
	}
	if len(contacts) == 0 {
		return c.Status(404).JSON(fiber.Map{"message": "No contacts found"})
	}

	return c.Status(200).JSON(fiber.Map{"message": "Contacts", "data": contacts})
}

func GetSpecificContact(c *fiber.Ctx) error {
	param := c.Params("contactType")

	client := db.NewClient()
	if err := client.Connect(); err != nil {
		return c.Status(500).JSON(fiber.Map{"message": "Error connecting to the database"})
	}
	defer client.Disconnect()

	contact, err := client.Contacts.FindUnique(
		db.Contacts.ContactType.Equals(strings.ToLower(param)),
	).Exec(c.Context())
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"message": "Error fetching the contact"})
	}
	if contact == nil {
		return c.Status(404).JSON(fiber.Map{"message": "No contact found"})
	}

	return c.Status(200).JSON(fiber.Map{"message": "Contact", "data": contact})
}

func SetContact(c *fiber.Ctx) error {
	var body contactStruct

	if err := c.BodyParser(&body); err != nil {
		return c.Status(400).JSON(fiber.Map{"message": "Invalid input"})
	}

	client := db.NewClient()
	if err := client.Connect(); err != nil {
		return c.Status(500).JSON(fiber.Map{"message": "Error connecting to the database"})
	}
	defer client.Disconnect()

	contactType := strings.ToLower(body.ContactType)

	// Upsert: If the contact exists, update it; otherwise, create a new one.
	_, err := client.Contacts.UpsertOne(
		db.Contacts.ContactType.Equals(contactType), // Condition for upsert
	).Create(
		db.Contacts.ContactType.Set(contactType),
		db.Contacts.ContactValue.Set(body.ContactData),
	).Update(
		db.Contacts.ContactValue.Set(body.ContactData),
	).Exec(c.Context())

	if err != nil {
		return c.Status(500).JSON(fiber.Map{"message": "Error creating or updating contact"})
	}

	return c.Status(200).JSON(fiber.Map{"message": "Successfully created or updated contact"})
}

func DeleteContact(c *fiber.Ctx) error {
	param := c.Params("contactType")

	// Prisma client
	client := db.NewClient()
	if err := client.Connect(); err != nil {
		return c.Status(500).JSON(fiber.Map{"message": "Error connecting to the database"})
	}
	defer client.Disconnect()

	// Get the specific contact
	contact, err := client.Contacts.FindUnique(
		db.Contacts.ContactType.Equals(strings.ToLower(param)),
	).Delete().Exec(c.Context())
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"message": "Error fetching the contact"})
	}
	if contact == nil {
		return c.Status(404).JSON(fiber.Map{"message": "No contact found"})
	}

	return c.Status(200).JSON(fiber.Map{"message": "Successfully deleted contact"})
}
