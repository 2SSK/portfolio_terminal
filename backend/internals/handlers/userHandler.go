package userHandler

import "github.com/gofiber/fiber/v2"

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Login(c *fiber.Ctx) error {
	return nil
}
