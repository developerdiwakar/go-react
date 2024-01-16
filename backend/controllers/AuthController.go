package controllers

import (
	"errors"
	"gofiber-api/database"
	"gofiber-api/helpers"
	"gofiber-api/http/requests"
	"gofiber-api/models"
	"gofiber-api/services"
	"gofiber-api/validators"
	"log"

	"github.com/gofiber/fiber/v2"
)

// User Registration
func Register(c *fiber.Ctx) error {

	requestBody := &requests.UserRegisterRequest{}
	response := services.GetBaseResponseObject()

	if err := c.BodyParser(requestBody); err != nil {
		response["errors"] = err
		return c.Status(fiber.StatusUnprocessableEntity).JSON(response)
	}

	if err := validators.ValidateStruct(requestBody); err != nil {
		// log.Println(err)
		response["errors"] = err
		response["message"] = MsgInvalidRequestFormat
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

// User Login with 
func Login(c *fiber.Ctx) error {
	requestBody := &requests.UserLoginRequest{}
	response := services.GetBaseResponseObject()
	user := models.User{}

	if err := c.BodyParser(requestBody); err != nil {
		response["errors"] = err
		return c.Status(fiber.StatusUnprocessableEntity).JSON(response)
	}

	if err := validators.ValidateStruct(requestBody); err != nil {
		response["errors"] = err
		response["message"] = MsgInvalidRequestFormat
		return c.Status(fiber.StatusForbidden).JSON(response)
	} else {
		username := requestBody.Username
		password := requestBody.Password

		// Prepare db query
		where := ""
		if helpers.CheckValidEmail(username) {
			where += "email = ?"
		} else {
			username = "+91" + username
			where += "mobile_number = ?"
		}

		// Fetch user by username either email or mobile_number
		userObj := database.Conn().DB.Where(where, username).First(&user)
		if err := userObj.Error; err != nil {
			response["errors"] = validators.GetSQLError(err)
			response["message"] = MsgInvalidUsername
			return c.Status(fiber.StatusForbidden).JSON(response)
		}
		// Check whether Password is correct or not
		if ok := helpers.CheckPasswordHash(password, user.Password); !ok {
			response["errors"] = validators.GetGeneralError(errors.New("login failed"))
			response["message"] = MsgIncorrectPassword
			return c.Status(fiber.StatusUnauthorized).JSON(response)
		}

		// Create Paseto Token
		token, err := CreateJwtToken(int(user.ID))
		if err != nil {
			log.Panicln(err)
		}

		// Success Response
		response["success"] = true
		response["message"] = "Login Successful"
		response["data"] = map[string]string{"token": token}
		return c.Status(fiber.StatusOK).JSON(response)
	}
}
