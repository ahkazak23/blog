package router

import (
	"blog/controller"
	"blog/middleware"
	"github.com/gofiber/fiber/v2"
)

// SetupRoutes Setup routing information
func SetupRoutes(app *fiber.App) {
	/*
		list    ->  GET
		read blog => GET (id)
		add		->	POST
		update  ->	PUT
		delete  ->	DELETE
	*/

	app.Get("/", controller.BlogList)
	app.Get("/:id", controller.BlogDetail)
	app.Post("/", controller.BlogCreate)
	app.Put("/:id", controller.BlogUpdate)
	app.Delete("/:id", controller.BlogDelete)

	app.Post("/login", controller.Login)
	app.Post("/register", controller.Register)
	// r.GET("/logout", controller.Logout)

	private := app.Group("/private")

	private.Use(middleware.Authenticate)

	private.Get("/refreshtoken", controller.RefreshToken)
	private.Get("/profile", controller.Profile)

	// r.GET("/refreshtoken", controller.RefreshToken)
}
