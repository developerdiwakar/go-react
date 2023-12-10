package validators

import (
	"encoding/json"
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

// Define SQL error
type SQLError struct {
	Number   int
	SQLState []int8
	Message  string
}

// Create a validator object
var validate = validator.New()

// Package Initialization
func init() {
	// Register Custom Validation rules
	validate.RegisterValidation("unique", IsUnique)
}

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

func GetSQLError(err error) *SQLError {
	var errByte []byte
	errByte, err1 := json.Marshal(err)
	if err1 != nil {
		log.Fatalln(err)
	}
	log.Println("SQLError: ", string(errByte))
	var sqlErr *SQLError
	json.Unmarshal(errByte, &sqlErr)
	sqlErr.Message = GetSqlErrMsg(sqlErr.Number, sqlErr.Message)
	return sqlErr
}

// TODO: Validation rule to check whether the target value is unique or not
func IsUnique(data validator.FieldLevel) bool {

	// targetValue := data.Field().String()
}
