package routes

import (
	"gofiber-api/controllers"

	"github.com/gofiber/fiber/v2"
)

func Load(app *fiber.App) {
	// Routes
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("index", fiber.Map{
			"title": "Gofiber with React App",
		})
	})
	// Creating API group for API routes
	api := app.Group("api")
	// Auth Routes
	api.Post("register", controllers.Register)
}
