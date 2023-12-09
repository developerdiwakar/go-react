package validators

import (
	"log"

	"github.com/go-playground/validator/v10"
)

// Define a struct that would handle errors for us
type ErrorResponse struct {
	FailedField string      `json:"failedField"`
	Tag         string      `json:"tag"`
	Message     string      `json:"message"`
	Value       interface{} `json:"value"`
	Param       string      `json:"param"`
}

// Create a validator object
var validate = validator.New()

// Package Initialization
func init() {}

// Create a function that validates post body and return error if there are any
func ValidateStruct(postBody interface{}) []*ErrorResponse {
	err := validate.Struct(postBody)
	return handleErrors(err)

}

func handleErrors(err error) []*ErrorResponse {
	var errors []*ErrorResponse
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			log.Println(err)
			element := ErrorResponse{
				FailedField: err.Field(),
				Tag:         err.Tag(),
				Value:       err.Value(),
				Param:       err.Param(),
				Message:     GetErrMsg(err.Tag(), err.Field(), err.Value(), err.Param()),
			}
			errors = append(errors, &element)
		}
	}
	return errors
}
