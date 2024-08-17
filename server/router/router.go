package router

import (
	"blog/controller"
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
}
