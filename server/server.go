package main

import (
	"blog/database"
	"github.com/gofiber/fiber/v2"
)

func main() {
	database.ConnectDB()

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"message": "Hello World!"})
		return c.SendString("Fang Yuan")
	})
	app.Listen(":8000")

}
