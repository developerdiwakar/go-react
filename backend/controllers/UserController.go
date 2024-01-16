package controllers

import (
	"github.com/gofiber/fiber/v2"
)

func GetUserInfo(c *fiber.Ctx) error {
	user := c.Locals("user")
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"user":    user,
	})
}
