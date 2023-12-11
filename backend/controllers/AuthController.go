package controllers

import (
	"gofiber-api/database"
	"gofiber-api/helpers"
	"gofiber-api/models"
	"gofiber-api/requests"
	"gofiber-api/services"
	"gofiber-api/validators"

	"github.com/gofiber/fiber/v2"
)

func Register(c *fiber.Ctx) error {

	requestBody := &requests.UserRegister{}
	response := services.GetBaseResponseObject()

	if err := c.BodyParser(requestBody); err != nil {
		response["errors"] = err
		return c.Status(fiber.StatusUnprocessableEntity).JSON(response)
	}

	if err := validators.ValidateStruct(requestBody); err != nil {
		// log.Println(err)
		response["errors"] = err
		response["message"] = "Invalid Request Formats"
		return c.Status(fiber.StatusForbidden).JSON(response)
	} else {
		// Prepare requests
		hashPassword, _ := helpers.HashPassword(requestBody.Password)
		// Database operation
		userCreated := database.Conn().DB.Create(&models.User{
			Name:         requestBody.Name,
			Email:        requestBody.Email,
			MobileNumber: requestBody.MobileNumber,
			Password:     hashPassword,
		})
		if err := userCreated.Error; err != nil {
			response["errors"] = validators.GetSQLError(err)
			response["message"] = "Registration Failed"
			return c.Status(fiber.StatusForbidden).JSON(response)
		}
		response["success"] = true
		response["message"] = "Registration Successful"

		return c.Status(fiber.StatusOK).JSON(response)
	}
}
