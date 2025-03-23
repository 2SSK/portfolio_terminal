package projectHandler

import (
	"fmt"
	"mime/multipart"

	"github.com/2SSK/portfolio_terminal/backend/config"
	"github.com/2SSK/portfolio_terminal/backend/prisma/db"
	"github.com/2SSK/portfolio_terminal/backend/utils/fileHandler"
	"github.com/gofiber/fiber/v2"
)

type ProjectRequest struct {
	Title       string `json:"title"`
	Preview     string `json:"preview"`
	Url         string `json:"url"`
	Github      string `json:"github"`
	Description string `json:"description"`
	publicID    string // Store Cloudinary public ID internally
}

func AddProject(c *fiber.Ctx) error {
	user := c.Locals("user").(*db.UserModel)
	userID := user.ID

	form, err := c.MultipartForm()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "failed to parse form data"})
	}

	project := new(ProjectRequest)
	if titles := form.Value["title"]; len(titles) == 0 || titles[0] == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "title is required"})
	} else {
		project.Title = titles[0]
	}
	project.Url = getFormValue(form, "url", "")
	project.Github = getFormValue(form, "github", "")
	project.Description = getFormValue(form, "description", "")

	files, ok := form.File["preview"]
	if !ok || len(files) == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "preview file is required"})
	}
	file := files[0]

	if existing, _ := config.PrismaClient.Projects.FindUnique(
		db.Projects.Title.Equals(project.Title),
	).Exec(c.Context()); existing != nil {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{"error": "project with this title already exists"})
	}

	if err := fileHandler.ValidateFile(file, "project"); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	url, publicID, _, err := fileHandler.UploadFile(file, "project", userID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	project.Preview = url
	project.publicID = publicID

	newProject, err := config.PrismaClient.Projects.CreateOne(
		db.Projects.Title.Set(project.Title),
		db.Projects.Preview.Set(project.Preview),
		db.Projects.URL.Set(project.Url),
		db.Projects.Github.Set(project.Github),
		db.Projects.Description.Set(project.Description),
		db.Projects.UserID.Set(userID),
	).Exec(c.Context())
	if err != nil {
		fileHandler.DeleteFile(publicID)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "failed to create project", "details": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Project created successfully",
		"project": newProject,
	})
}

func GetAllProjects(c *fiber.Ctx) error {
	user := c.Locals("user").(*db.UserModel)
	userID := user.ID

	projects, err := config.PrismaClient.Projects.FindMany(
		db.Projects.UserID.Equals(userID),
	).Exec(c.Context())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "failed to get projects"})
	}
	if len(projects) == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "no projects found"})
	}

	return c.JSON(fiber.Map{"projects": projects})
}

func GetProject(c *fiber.Ctx) error {
	user := c.Locals("user").(*db.UserModel)
	userID := user.ID

	projectID, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid project ID"})
	}

	project, err := config.PrismaClient.Projects.FindFirst(
		db.Projects.ID.Equals(projectID),
		db.Projects.UserID.Equals(userID),
	).Exec(c.Context())
	if err != nil || project == nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "project not found"})
	}

	return c.JSON(fiber.Map{"project": project})
}

func UpdateProject(c *fiber.Ctx) error {
	user := c.Locals("user").(*db.UserModel)
	userID := user.ID

	projectID, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid project ID"})
	}

	existing, err := config.PrismaClient.Projects.FindFirst(
		db.Projects.ID.Equals(projectID),
		db.Projects.UserID.Equals(userID),
	).Exec(c.Context())
	if err != nil || existing == nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "project not found"})
	}

	form, err := c.MultipartForm()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "failed to parse form data"})
	}

	project := &ProjectRequest{
		Title:       existing.Title,
		Preview:     existing.Preview,
		Url:         existing.URL,
		Github:      existing.Github,
		Description: existing.Description,
		publicID:    "", // Will be updated if new file uploaded
	}

	if titles := form.Value["title"]; len(titles) > 0 && titles[0] != "" {
		project.Title = titles[0]
	}
	project.Url = getFormValue(form, "url", project.Url)
	project.Github = getFormValue(form, "github", project.Github)
	project.Description = getFormValue(form, "description", project.Description)

	if files, ok := form.File["preview"]; ok && len(files) > 0 {
		file := files[0]
		if err := fileHandler.ValidateFile(file, "project"); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
		}

		url, publicID, _, err := fileHandler.UploadFile(file, "project", userID)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}
		project.Preview = url
		project.publicID = publicID

		// Delete old preview if it exists (assuming previous publicID is derivable)
		// Note: You might need to store publicID in DB for accurate deletion
		oldPublicID := fmt.Sprintf("portfolio/project/%d_%s", userID, existing.Title)
		fileHandler.DeleteFile(oldPublicID)
	}

	updatedProject, err := config.PrismaClient.Projects.FindUnique(
		db.Projects.ID.Equals(projectID),
	).Update(
		db.Projects.Title.Set(project.Title),
		db.Projects.Preview.Set(project.Preview),
		db.Projects.URL.Set(project.Url),
		db.Projects.Github.Set(project.Github),
		db.Projects.Description.Set(project.Description),
	).Exec(c.Context())
	if err != nil {
		if project.publicID != "" {
			fileHandler.DeleteFile(project.publicID)
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "failed to update project", "details": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Project updated successfully",
		"project": updatedProject,
	})
}

func DeleteProject(c *fiber.Ctx) error {
	user := c.Locals("user").(*db.UserModel)
	userID := user.ID

	projectID, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid project ID"})
	}

	existing, err := config.PrismaClient.Projects.FindFirst(
		db.Projects.ID.Equals(projectID),
		db.Projects.UserID.Equals(userID),
	).Exec(c.Context())
	if err != nil || existing == nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "project not found"})
	}

	// Delete preview from Cloudinary (assuming publicID format)
	publicID := fmt.Sprintf("portfolio/project/%d_%s", userID, existing.Title)
	fileHandler.DeleteFile(publicID)

	_, err = config.PrismaClient.Projects.FindUnique(
		db.Projects.ID.Equals(projectID),
	).Delete().Exec(c.Context())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "failed to delete project"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Project deleted successfully",
	})
}

// Helper function
func getFormValue(form *multipart.Form, key, defaultValue string) string {
	if values := form.Value[key]; len(values) > 0 {
		return values[0]
	}
	return defaultValue
}
