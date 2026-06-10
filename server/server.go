package main

import (
	"blog/database"
	"blog/router"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
	"log"
	"os"
	"strings"
)

func init() {
	if err := godotenv.Load(".env"); err != nil {
		log.Println("No .env file loaded, using environment variables.")
	}
	if err := godotenv.Load(".env.prod"); err != nil {
		log.Println("No .env.prod file loaded, using environment variables.")
	}

	database.ConnectDB()
}

func main() {

	sqlDB, err := database.DBConn.DB()

	if err != nil {
		panic("Error in sql connection.")
	}
	defer sqlDB.Close()

	app := fiber.New()

	app.Static("/static", "./static")

	allowOrigins := strings.TrimSpace(os.Getenv("cors_allow_origins"))
	if allowOrigins == "" {
		allowOrigins = "http://localhost:3000"
	}

	app.Use(cors.New(cors.Config{
		AllowOrigins: allowOrigins,
		AllowMethods: "GET,POST,PUT,DELETE,OPTIONS",
		AllowHeaders: "Origin, Content-Type, Accept, Authorization, token",
	}))

	router.SetupRoutes(app)

	app.Listen(":8000")

}
