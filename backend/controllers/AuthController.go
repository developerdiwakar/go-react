package controllers

import (
	"gofiber-api/database"
	"gofiber-api/managers"
	"gofiber-api/models"
	"gofiber-api/structs"
	"gofiber-api/validators"
	"log"

	"github.com/gofiber/fiber/v2"
)

func Register(c *fiber.Ctx) error {

	requestBody := &structs.UserRegister{}
	response := managers.GetBaseResponseObject()

	if err := c.BodyParser(requestBody); err != nil {
		response["errors"] = err
		return c.Status(fiber.StatusUnprocessableEntity).JSON(response)
	}

	if err := validators.ValidateStruct(requestBody); err != nil {
		log.Println(err)
		response["errors"] = err
		response["message"] = "Invalid Request Formats"
		return c.Status(fiber.StatusForbidden).JSON(response)
	} else {
		// Database operation
		// log.Println("database.Conn().DB", database.Conn())
		userCreated := database.Conn().DB.Create(&models.User{
			Name:         requestBody.Name,
			Email:        requestBody.Email,
			MobileNumber: requestBody.MobileNumber,
			Password:     requestBody.Password,
		})
		log.Println(userCreated)
		if err := userCreated.Error; err != nil {
			log.Println(err)
			response["errors"] = err
			response["message"] = "Registration Failed"
			return c.Status(fiber.StatusForbidden).JSON(response)
		}

		response["message"] = "Registration Successful"

		return c.Status(fiber.StatusOK).JSON(response)
	}
}
