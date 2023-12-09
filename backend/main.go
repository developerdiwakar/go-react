package main

import (
	"gofiber-api/database"
	"gofiber-api/migrations"
	"gofiber-api/routes"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/template/handlebars/v2"
	"github.com/joho/godotenv"
)

func main() {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Fatalln(err)
	}
	// Create template engine instance
	engine := handlebars.New("./views", ".hbs")

	// Create new fiber instance
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	// Load the Database

	db := database.Conn()
	defer database.Close()
	// Called Migrations
	migrations.Migrate(db.DB)
	// Called Middlewares
	app.Use(cors.New())
	// Load routes
	routes.Load(app)

	app.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(404) // => 404 "Not Found"
	})

	// Start the app on port 8000
	log.Fatal(app.Listen(":8000"))
}
