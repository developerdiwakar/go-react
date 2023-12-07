package routes

import "github.com/gofiber/fiber/v2"

func Load(app *fiber.App) {
	// Routes
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("index", fiber.Map{
			"title": "Gofiber with React App",
		})
	})
}
