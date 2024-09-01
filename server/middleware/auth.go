package middleware

import (
	"blog/helper"
	"github.com/gofiber/fiber/v2"
	"log"
)

func Authenticate(c *fiber.Ctx) error {
	token := c.Get("token")

	if token == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Token not present."})
	}

	claims, msg := helper.ValidateToken(token)

	log.Println(claims)

	if msg != "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": msg})
	}

	// Setting the claims email in context locals
	c.Locals("email", claims.Email)

	// Continue to the next handler
	return c.Next()
}
