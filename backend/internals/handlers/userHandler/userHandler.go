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

func generateToken(userId int) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["userId"] = userId
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix() // Token expires in 72 hours

	return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}
