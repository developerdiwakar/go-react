package routes

import (
	"gofiber-api/controllers"
	"gofiber-api/http/middlewares"

	"github.com/gofiber/fiber/v2"
)

func Load(app *fiber.App) {
	// Routes
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("index", fiber.Map{
			"title": "Gofiber with React App",
		})
	})
	// Auth Routes
	app.Post("/register", controllers.Register)
	app.Post("/login", controllers.Login)
	// Creating API group for Authenticated API routes
	api := app.Group("api", middlewares.JWTAuthMiddleware)
	api.Get("user", controllers.GetUserInfo)
}
