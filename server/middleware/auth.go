package middleware

import (
	"blog/helper"
	"github.com/gofiber/fiber/v2"
	"strings"
)

func Authenticate(c *fiber.Ctx) error {
	token := strings.TrimSpace(c.Get("token"))
	authHeader := strings.TrimSpace(c.Get("Authorization"))
	if strings.HasPrefix(strings.ToLower(authHeader), "bearer ") {
		token = strings.TrimSpace(authHeader[7:])
	}

	if token == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Token not present."})
	}

	claims, msg := helper.ValidateToken(token)

	if msg != "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": msg})
	}

	// Setting the claims email in context locals
	c.Locals("email", claims.Email)

	// Continue to the next handler
	return c.Next()
}
