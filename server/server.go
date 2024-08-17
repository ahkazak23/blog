package main

import (
	"blog/database"
	"blog/router"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
	"log"
)

func init() {
	if err := godotenv.Load(".env.prod"); err != nil {
		log.Fatal("Error in loading .env file.")
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

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	router.SetupRoutes(app)

	app.Listen(":8000")

}
