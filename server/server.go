package main

import (
	"blog/database"
	"github.com/gofiber/fiber/v2"
)

func init() {
	database.ConnectDB()
}

func main() {

	sqlDB, err := database.DBConn.DB()

	if err != nil {
		panic("Error in sql connection.")
	}
	defer sqlDB.Close()

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"message": "Hello World!"})
		return c.SendString("Fang Yuan")
	})
	app.Listen(":8000")

}
