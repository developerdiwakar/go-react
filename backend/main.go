package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/handlebars/v2"
)

func main() {
	// Create template engine instance
	engine := handlebars.New("./views", ".hbs")

	// Create new fiber instance
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	// Routes
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("index", fiber.Map{
			"title": "Gofiber with React App",
		})
	})

	// Start the app on port 8000
	log.Fatal(app.Listen(":8000"))
}
