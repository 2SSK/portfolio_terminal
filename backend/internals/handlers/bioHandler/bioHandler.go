package bioHandler

import (
	"mime/multipart"

	"github.com/2SSK/portfolio_terminal/backend/config"
	"github.com/2SSK/portfolio_terminal/backend/prisma/db"
	"github.com/2SSK/portfolio_terminal/backend/utils/fileHandler"
	"github.com/gofiber/fiber/v2"
)

type BioRequest struct {
	Image       string `json:"image"`
	Name        string `json:"name"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

func AddBio(c *fiber.Ctx) error {
	user := c.Locals("user").(*db.UserModel)
	userID := user.ID

	form, err := c.MultipartForm()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "failed to parse form data"})
	}

	bio := new(BioRequest)
	if titles := form.Value["title"]; len(titles) == 0 || titles[0] == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "title is required"})
	} else {
		bio.Title = titles[0]
	}
	bio.Name = getFormValue(form, "name", "")
	bio.Description = getFormValue(form, "description", "")

	images, ok := form.File["image"]
	if !ok || len(images) == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "image file is required"})
	}
	image := images[0]

	if err := fileHandler.ValidateFile(image, "bio"); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	url, publicID, _, err := fileHandler.UploadFile(image, "bio", userID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	bio.Image = url

	newBio, err := config.PrismaClient.Bio.UpsertOne(
		db.Bio.UserID.Equals(userID),
	).Create(
		db.Bio.Image.Set(bio.Image),
		db.Bio.Name.Set(bio.Name),
		db.Bio.Title.Set(bio.Title),
		db.Bio.Description.Set(bio.Description),
		db.Bio.User.Link(db.User.ID.Equals(userID)),
	).Update(
		db.Bio.Image.Set(bio.Image),
		db.Bio.Name.Set(bio.Name),
		db.Bio.Title.Set(bio.Title),
		db.Bio.Description.Set(bio.Description),
	).Exec(c.Context())
	if err != nil {
		fileHandler.DeleteFile(publicID) // Cleanup on DB failure
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "failed to update bio", "details": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Bio updated successfully",
		"bio":     newBio,
	})
}

func GetBio(c *fiber.Ctx) error {
	user := c.Locals("user").(*db.UserModel)
	userID := user.ID

	bio, err := config.PrismaClient.Bio.FindUnique(
		db.Bio.UserID.Equals(userID),
	).Exec(c.Context())
	if err == db.ErrNotFound {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "bio not found"})
	}
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "failed to get bio", "details": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Bio fetched successfully",
		"bio":     bio,
	})
}

// Helper to simplify form value extraction
func getFormValue(form *multipart.Form, key, defaultValue string) string {
	if values := form.Value[key]; len(values) > 0 {
		return values[0]
	}
	return defaultValue
}
