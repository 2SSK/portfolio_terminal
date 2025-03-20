package userHandler

import (
	"os"
	"time"

	"github.com/2SSK/portfolio_terminal/backend/config"
	"github.com/2SSK/portfolio_terminal/backend/prisma/db"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type input struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// UserResponse is what we send back to the client
type UserResponse struct {
	ID    int    `json:"id"`
	Email string `json:"email"`
	Token string `json:"token"`
}

func CreateUser(c *fiber.Ctx) error {
	var body input

	if err := c.BodyParser(&body); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request"})
	}

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(body.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to hash password"})
	}

	user, err := config.PrismaClient.User.CreateOne(
		db.User.Email.Set(body.Email),
		db.User.Password.Set(string(hashedPassword)),
	).Exec(c.Context())

	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to create user"})
	}

	// Create an initial tool instance for the user
	_, err = config.PrismaClient.Tools.CreateOne(
		db.Tools.User.Link(db.User.ID.Equals(user.ID)),
	).Exec(c.Context())
	if err != nil {
		config.PrismaClient.User.FindUnique(db.User.ID.Equals(user.ID)).Delete().Exec(c.Context())
		return c.Status(500).JSON(fiber.Map{"error": "Failed to create user's tools"})
	}

	// Generate JWT token
	token, err := generateToken(user.ID)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to generate token"})
	}

	return c.Status(201).JSON(UserResponse{
		ID:    user.ID,
		Email: user.Email,
		Token: token,
	})
}

func GetUser(c *fiber.Ctx) error {
	var body input

	if err := c.BodyParser(&body); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid inputs"})
	}

	user, err := config.PrismaClient.User.FindUnique(
		db.User.Email.Equals(body.Email),
	).Exec(c.Context())
	if err != nil {
		return c.Status(401).JSON(fiber.Map{"error": "User not found"})
	}

	// Compare the password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))
	if err != nil {
		return c.Status(401).JSON(fiber.Map{"error": "Invalid password"})
	}

	// Generate JWT token
	token, err := generateToken(user.ID)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to generate token"})
	}

	return c.Status(200).JSON(UserResponse{
		ID:    user.ID,
		Email: user.Email,
		Token: token,
	})
}

func DeleteUser(c *fiber.Ctx) error {
	userId := c.QueryInt("userId")
	if userId == 0 {
		return c.Status(400).JSON(fiber.Map{"error": "UserId is required"})
	}

	// First find the user to make sure they exist and store their email for the response
	existingUser, err := config.PrismaClient.User.FindUnique(
		db.User.ID.Equals(userId),
	).Exec(c.Context())
	if err != nil || existingUser == nil {
		return c.Status(404).JSON(fiber.Map{"error": "User not found"})
	}
	userEmail := existingUser.Email // Store email before deletion

	// Delete user's links
	links, err := config.PrismaClient.Links.FindMany(
		db.Links.UserID.Equals(userId),
	).Exec(c.Context())
	if err == nil {
		for _, link := range links {
			_, err := config.PrismaClient.Links.FindUnique(
				db.Links.ID.Equals(link.ID),
			).Delete().Exec(c.Context())
			if err != nil {
				return c.Status(500).JSON(fiber.Map{"error": "Failed to delete user's links"})
			}
		}
	}

	// Delete user's resume
	resume, err := config.PrismaClient.Resume.FindUnique(
		db.Resume.UserID.Equals(userId),
	).Exec(c.Context())
	if err == nil && resume != nil {
		_, err := config.PrismaClient.Resume.FindUnique(
			db.Resume.UserID.Equals(userId),
		).Delete().Exec(c.Context())
		if err != nil {
			return c.Status(500).JSON(fiber.Map{"error": "Failed to delete user's resume"})
		}
	}

	// Delete user's projects
	projects, err := config.PrismaClient.Projects.FindMany(
		db.Projects.UserID.Equals(userId),
	).Exec(c.Context())
	if err == nil {
		for _, project := range projects {
			_, err := config.PrismaClient.Projects.FindUnique(
				db.Projects.ID.Equals(project.ID),
			).Delete().Exec(c.Context())
			if err != nil {
				return c.Status(500).JSON(fiber.Map{"error": "Failed to delete user's projects"})
			}
		}
	}

	// Delete user's experiences
	experiences, err := config.PrismaClient.Experience.FindMany(
		db.Experience.UserID.Equals(userId),
	).Exec(c.Context())
	if err == nil {
		for _, exp := range experiences {
			_, err := config.PrismaClient.Experience.FindUnique(
				db.Experience.ID.Equals(exp.ID),
			).Delete().Exec(c.Context())
			if err != nil {
				return c.Status(500).JSON(fiber.Map{"error": "Failed to delete user's experiences"})
			}
		}
	}

	// Finally delete the user
	_, err = config.PrismaClient.User.FindUnique(
		db.User.ID.Equals(userId),
	).Delete().Exec(c.Context())
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to delete user"})
	}

	return c.Status(200).JSON(fiber.Map{
		"message": "User and all associated data deleted successfully",
		"email":   userEmail,
	})
}

func generateToken(userId int) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["userId"] = userId
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix() // Token expires in 72 hours

	return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}
