package userHandler

import (
	"context"
	"os"
	"strings"
	"time"

	"github.com/2SSK/portfolio_terminal/backend/config"
	"github.com/2SSK/portfolio_terminal/backend/prisma/db"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

// input is the expected request body structure
type input struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// UserResponse is what we send back to the client
type UserResponse struct {
	ID           int    `json:"id"`
	Email        string `json:"email"`
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

func CreateUser(c *fiber.Ctx) error {
	var body input

	if err := c.BodyParser(&body); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request"})
	}

	if existing, _ := config.PrismaClient.User.FindUnique(
		db.User.Email.Equals(body.Email),
	).Exec(c.Context()); existing != nil {
		return c.Status(400).JSON(fiber.Map{"error": "User already exists"})
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

	// Generate JWT tokens
	accessToken, err := generateAccessToken(user.ID, c.Context())
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to generate access token"})
	}

	refreshToken, err := generateRefreshToken(user.ID, c.Context())
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to generate refresh token"})
	}

	return c.Status(201).JSON(UserResponse{
		ID:           user.ID,
		Email:        user.Email,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
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
	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password)); err != nil {
		return c.Status(401).JSON(fiber.Map{"error": "Invalid password"})
	}

	// Generate JWT tokens
	accessToken, err := generateAccessToken(user.ID, c.Context()) // Fixed by adding c.Context()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to generate access token"})
	}

	refreshToken, err := generateRefreshToken(user.ID, c.Context())
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to generate refresh token"})
	}

	return c.Status(200).JSON(UserResponse{
		ID:           user.ID,
		Email:        user.Email,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	})
}

func GetRefreshToken(c *fiber.Ctx) error {
	type RefreshInput struct {
		RefreshToken string `json:"refreshToken"`
	}

	var input RefreshInput
	if err := c.BodyParser(&input); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid inputs"})
	}

	// Verify the refresh token
	token, err := jwt.Parse(input.RefreshToken, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fiber.NewError(fiber.StatusUnauthorized, "Unexpected signing method")
		}
		return []byte(os.Getenv("JWT_REFRESH_SECRET")), nil
	})
	if err != nil || !token.Valid {
		return c.Status(401).JSON(fiber.Map{"error": "Invalid or expired refresh token"})
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return c.Status(401).JSON(fiber.Map{"error": "Invalid refresh token claims"})
	}

	userID, ok := claims["userId"].(float64)
	if !ok {
		return c.Status(401).JSON(fiber.Map{"error": "User ID not found in refresh token"})
	}

	// Check if the refresh token exists and not expired
	refreshTokenDB, err := config.PrismaClient.RefreshToken.FindUnique(
		db.RefreshToken.Token.Equals(input.RefreshToken),
	).With(
		db.RefreshToken.User.Fetch(),
	).Exec(c.Context())
	if err != nil || refreshTokenDB == nil || refreshTokenDB.ExpiresAt.Before(time.Now()) {
		return c.Status(401).JSON(fiber.Map{"error": "Invalid or expired refresh token"})
	}

	// Generate new JWT access token
	accessToken, err := generateAccessToken(int(userID), c.Context())
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to generate access token"})
	}

	return c.JSON(fiber.Map{
		"accessToken": accessToken,
	})
}

func generateAccessToken(userId int, ctx context.Context) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId": userId,
		"exp":    time.Now().Add(time.Hour * 24 * 7).Unix(), // 7 days
		"iat":    time.Now().Unix(),
	})

	signedToken, err := token.SignedString([]byte(os.Getenv("JWT_REFRESH_SECRET")))
	if err != nil {
		return "", err
	}

	// Store refresh token in database (assuming this was intended for access token storage)
	_, err = config.PrismaClient.RefreshToken.CreateOne(
		db.RefreshToken.Token.Set(signedToken),
		db.RefreshToken.User.Link(db.User.ID.Equals(userId)),
		db.RefreshToken.ExpiresAt.Set(time.Now().Add(time.Hour*24*7)),
	).Exec(ctx)
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func generateRefreshToken(userId int, ctx context.Context) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId": userId,
		"exp":    time.Now().Add(time.Hour * 24 * 30).Unix(), // 30 days
		"iat":    time.Now().Unix(),
	})

	signedToken, err := token.SignedString([]byte(os.Getenv("JWT_REFRESH_SECRET")))
	if err != nil {
		return "", err
	}

	// Store refresh token in database
	_, err = config.PrismaClient.RefreshToken.CreateOne(
		db.RefreshToken.Token.Set(signedToken),
		db.RefreshToken.User.Link(db.User.ID.Equals(userId)),
		db.RefreshToken.ExpiresAt.Set(time.Now().Add(time.Hour*24*30)),
	).Exec(ctx)
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func VerifyToken(c *fiber.Ctx) error {
	tokenString := c.Get("Authorization")
	if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer ") {
		return c.Status(401).JSON(fiber.Map{"error": "No token provided"})
	}

	tokenString = strings.TrimPrefix(tokenString, "Bearer ")

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fiber.NewError(fiber.StatusUnauthorized, "Unexpected signing method")
		}
		return []byte(os.Getenv("JWT_REFRESH_SECRET")), nil
	})

	if err != nil || !token.Valid {
		return c.Status(401).JSON(fiber.Map{"error": "Invalid token"})
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return c.Status(401).JSON(fiber.Map{"error": "Invalid token claims"})
	}

	userID, ok := claims["userId"].(float64)
	if !ok {
		return c.Status(401).JSON(fiber.Map{"error": "User ID not found"})
	}

	user, err := config.PrismaClient.User.FindUnique(
		db.User.ID.Equals(int(userID)),
	).Exec(c.Context())
	if err != nil {
		return c.Status(401).JSON(fiber.Map{"error": "User not found"})
	}

	return c.JSON(UserResponse{
		ID:          user.ID,
		Email:       user.Email,
		AccessToken: tokenString,
	})
}
