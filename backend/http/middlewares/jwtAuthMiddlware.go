package middlewares

import (
	"gofiber-api/database"
	"gofiber-api/models"
	"gofiber-api/services/tokens"
	"log"
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
)

var user models.User

func JWTAuthMiddleware(c *fiber.Ctx) error {
	// Get Authorization header fronm request header
	authorization := c.Get("Authorization")
	if authorization == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Unauthorized Access!",
		})
	}
	bearerToken := strings.Split(authorization, " ")
	tokePrefix := bearerToken[0]
	token := bearerToken[1]

	maker, err := tokens.NewJWTMaker(os.Getenv("AUTH_KEY"))
	if err != nil {
		log.Panicln(err)
	}
	if token == "" || tokePrefix != "Bearer" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Authorization Error: Bearer token missing or mismatched",
		})
	}
	payload, err := maker.VerifyToken(token)

	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Authorization Error: " + err.Error(),
		})
	}
	// Get user from db
	userObj := database.Conn().DB.Where("id = ?", payload.Subject).First(&user)
	if err := userObj.Error; err != nil {
		return c.Status(fiber.StatusServiceUnavailable).JSON(fiber.Map{
			"message": "Server Error: " + err.Error(),
		})
	}

	c.Locals("user", user)
	return c.Next()
}
