package projectHandler

import (
	"fmt"
	"path/filepath"

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
}

func AddProject(c *fiber.Ctx) error {
	// Parse userId from query parameter
	userId := c.QueryInt("userId")

	// Parse multipart form data
	form, err := c.MultipartForm()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Failed to parse form data"})
	}

	// Extract and validate required fields
	project := new(ProjectRequest)
	if titles := form.Value["title"]; len(titles) == 0 || titles[0] == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Title is required"})
	} else {
		project.Title = titles[0]
	}

	// Extract optional fields with defaults
	if urls := form.Value["url"]; len(urls) > 0 {
		project.Url = urls[0]
	}
	if githubs := form.Value["github"]; len(githubs) > 0 {
		project.Github = githubs[0]
	}
	if descriptions := form.Value["description"]; len(descriptions) > 0 {
		project.Description = descriptions[0]
	}

	// Extract and validate preview file
	files, ok := form.File["preview"]
	if !ok || len(files) == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Preview file is required"})
	}
	file := files[0] // Only process the first file

	// Check for existing project
	if existing, _ := config.PrismaClient.Projects.FindUnique(
		db.Projects.Title.Equals(project.Title),
	).Exec(c.Context()); existing != nil {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{"error": "Project with this title already exists"})
	}

	// Generate unique filename
	ext := filepath.Ext(file.Filename)
	filename := fmt.Sprintf("%d_%s_preview%s", userId, project.Title, ext)
	filePath := fileHandler.GetFilePath(filename, "project")

	// Validate and save the file
	if err := fileHandler.ValidateFile(file, "project"); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	if err := c.SaveFile(file, filePath); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to save preview file"})
	}
	project.Preview = filename

	// Create project in the database
	newProject, err := config.PrismaClient.Projects.CreateOne(
		db.Projects.Title.Set(project.Title),
		db.Projects.Preview.Set(project.Preview),
		db.Projects.URL.Set(project.Url),
		db.Projects.Github.Set(project.Github),
		db.Projects.Description.Set(project.Description),
		db.Projects.UserID.Set(userId),
	).Exec(c.Context())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   "Failed to create project",
			"details": err.Error(),
		})
	}

	// Return success response
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Project created successfully",
		"project": newProject,
	})
}

func GetAllProjects(c *fiber.Ctx) error {
	userId := c.QueryInt("userId")

	projects, err := config.PrismaClient.Projects.FindMany(
		db.Projects.UserID.Equals(userId),
	).Exec(c.Context())

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to get projects"})
	}

	if len(projects) == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "No projects found"})
	}

	return c.JSON(fiber.Map{"projects": projects})
}

func GetProject(c *fiber.Ctx) error {
	userId := c.QueryInt("userId")

	projectId, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid project ID"})
	}

	project, err := config.PrismaClient.Projects.FindFirst(
		db.Projects.ID.Equals(projectId),
		db.Projects.UserID.Equals(userId),
	).Exec(c.Context())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   "Failed to get project",
			"details": err.Error(),
		})
	}
	if project == nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Project not found"})
	}

	return c.JSON(fiber.Map{"project": project})
}

func UpdateProject(c *fiber.Ctx) error {
	userId := c.QueryInt("userId")
	projectId, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid project ID"})
	}

	// Verify project exists and belongs to user
	existing, err := config.PrismaClient.Projects.FindFirst(
		db.Projects.ID.Equals(projectId),
		db.Projects.UserID.Equals(userId),
	).Exec(c.Context())
	if err != nil || existing == nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Project not found or doesn't belong to user"})
	}

	// Parse multipart form data
	form, err := c.MultipartForm()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Failed to parse form data"})
	}

	// Create project update struct with existing values as defaults
	project := &ProjectRequest{
		Title:       existing.Title,
		Preview:     existing.Preview,
		Url:         existing.URL,
		Github:      existing.Github,
		Description: existing.Description,
	}

	// Update fields if provided in the form
	if titles := form.Value["title"]; len(titles) > 0 && titles[0] != "" {
		project.Title = titles[0]
	}
	if urls := form.Value["url"]; len(urls) > 0 {
		project.Url = urls[0]
	}
	if githubs := form.Value["github"]; len(githubs) > 0 {
		project.Github = githubs[0]
	}
	if descriptions := form.Value["description"]; len(descriptions) > 0 {
		project.Description = descriptions[0]
	}

	// Handle preview file update if provided
	if files, ok := form.File["preview"]; ok && len(files) > 0 {
		file := files[0]

		// Generate new filename
		ext := filepath.Ext(file.Filename)
		filename := fmt.Sprintf("%d_%s_preview%s", userId, project.Title, ext)
		filePath := fileHandler.GetFilePath(filename, "project")

		// Validate and save new file
		if err := fileHandler.ValidateFile(file, "project"); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
		}
		if err := c.SaveFile(file, filePath); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to save preview file"})
		}

		// Update preview filename
		project.Preview = filename

		// Delete old preview file using filename and category
		if err := fileHandler.DeleteFile(existing.Preview, "project"); err != nil {
			// Log error but don't fail the request
			fmt.Printf("Failed to delete old preview file: %v\n", err)
		}
	}

	// Update project in database
	updatedProject, err := config.PrismaClient.Projects.FindUnique(
		db.Projects.ID.Equals(projectId),
	).Update(
		db.Projects.Title.Set(project.Title),
		db.Projects.Preview.Set(project.Preview),
		db.Projects.URL.Set(project.Url),
		db.Projects.Github.Set(project.Github),
		db.Projects.Description.Set(project.Description),
	).Exec(c.Context())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   "Failed to update project",
			"details": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Project updated successfully",
		"project": updatedProject,
	})
}

func DeleteProject(c *fiber.Ctx) error {
	userId := c.QueryInt("userId")

	projectId, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid project ID"})
	}

	// Verify project exists and belongs to user
	existing, err := config.PrismaClient.Projects.FindFirst(
		db.Projects.ID.Equals(projectId),
		db.Projects.UserID.Equals(userId),
	).Exec(c.Context())
	if err != nil || existing == nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Project not found or doesn't belong to user"})
	}

	// Delete the project
	deleted, err := config.PrismaClient.Projects.FindUnique(
		db.Projects.ID.Equals(projectId),
	).Delete().Exec(c.Context())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to delete project"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Project deleted successfully",
		"project": deleted,
	})
}
