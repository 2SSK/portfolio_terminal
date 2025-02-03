package projectHandler

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/2SSK/portfolio_terminal/backend/prisma/db"
	"github.com/gofiber/fiber/v2"
)

type ProjectStruct struct {
	Image       string `json:"Image"`
	Title       string `json:"Title"`
	Description string `json:"Description"`
	RepoURL     string `json:"RepoURL"`
	LiveURL     string `json:"LiveURL"`
	SkillsUsed  string `json:"SkillsUsed"`
}

func GetProject(c *fiber.Ctx) error {
	client := db.NewClient()
	if err := client.Connect(); err != nil {
		return c.Status(500).JSON(fiber.Map{"message": "Error connecting to the database"})
	}
	defer client.Disconnect()

	projects, err := client.Project.FindMany().Exec(c.Context())
	if len(projects) == 0 {
		return c.Status(404).JSON(fiber.Map{"message": "No projects found", "data": projects})
	}
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"message": "Error fetching projects"})
	}

	return c.Status(200).JSON(fiber.Map{"message": "Projects fetched successfully", "data": projects})
}

func SetProject(c *fiber.Ctx) error {
	// Parse form data
	project := new(ProjectStruct)
	project.Title = strings.ReplaceAll(strings.TrimSpace(c.FormValue("Title")), " ", "-")
	project.Description = c.FormValue("Description")
	project.RepoURL = c.FormValue("RepoURL")
	project.LiveURL = c.FormValue("LiveURL")
	project.SkillsUsed = c.FormValue("SkillsUsed")

	// Define the upload directory and Image subdirectory
	uploadDir := "./uploads"
	projectDir := filepath.Join(uploadDir, "projects")

	// Ensure the `uploads/projects/` directory exists
	if _, err := os.Stat(projectDir); os.IsNotExist(err) {
		os.MkdirAll(projectDir, 0755)
	}

	// Handle file upload (Image)
	file, err := c.FormFile("Image")
	if err == nil {
		filename := file.Filename
		savePath := filepath.Join(projectDir, filename)

		if err := c.SaveFile(file, savePath); err != nil {
			return c.Status(500).JSON(fiber.Map{"message": "Could not save the image"})
		}
		project.Image = filename
	}

	// Prisma client
	client := db.NewClient()
	if err := client.Connect(); err != nil {
		return c.Status(500).JSON(fiber.Map{"message": "Error connecting to the database"})
	}
	defer client.Disconnect()

	// Insert the project into the database
	_, err = client.Project.UpsertOne(
		db.Project.Title.Equals(project.Title),
	).Create(
		db.Project.Image.Set(project.Image),
		db.Project.Title.Set(strings.ToLower(project.Title)),
		db.Project.Description.Set(project.Description),
		db.Project.RepoURL.Set(project.RepoURL),
		db.Project.LiveURL.Set(project.LiveURL),
		db.Project.SkillsUsed.Set(project.SkillsUsed),
	).Update(
		db.Project.Image.Set(project.Image),
		db.Project.Title.Set(strings.ToLower(project.Title)),
		db.Project.Description.Set(project.Description),
		db.Project.RepoURL.Set(project.RepoURL),
		db.Project.LiveURL.Set(project.LiveURL),
		db.Project.SkillsUsed.Set(project.SkillsUsed),
	).Exec(c.Context())

	// If everything is fine, we can proceed to upload the image
	return c.Status(200).JSON(fiber.Map{"message": "Project created successfully"})
}

func DeleteProject(c *fiber.Ctx) error {
	projectName := c.Params("projectName")

	client := db.NewClient()
	if err := client.Connect(); err != nil {
		return c.Status(500).JSON(fiber.Map{"message": "Error connecting to the database"})
	}
	defer client.Disconnect()

	// Delete the project from the database
	project, err := client.Project.FindUnique(
		db.Project.Title.Equals(projectName),
	).Exec(c.Context())
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"message": "Error deleting project", "error": err.Error()})
	}

	// Delete the project from the database
	_, err = client.Project.FindUnique(
		db.Project.ID.Equals(project.ID),
	).Delete().Exec(c.Context())

	// Delete the associated image file if it exists
	if project.Image != "" {
		imagePath := filepath.Join("./uploads/projects", project.Image)
		if _, err := os.Stat(imagePath); err == nil {
			if err := os.Remove(imagePath); err != nil {
				return c.Status(500).JSON(fiber.Map{"message": "Error deleting project image", "error": err.Error()})
			}
		}
	}

	return c.Status(200).JSON(fiber.Map{"message": "Project deleted successfully"})
}
