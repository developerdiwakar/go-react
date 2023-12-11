package validators

import (
	"encoding/json"
	"fmt"
	"gofiber-api/database"
	"log"
	"strings"
	"time"

	"github.com/go-playground/validator/v10"
)

// Define a struct that would handle errors for us
type ErrorResponse struct {
	FailedField string      `json:"failedField"`
	Tag         string      `json:"-"`
	Message     string      `json:"message"`
	Value       interface{} `json:"value"`
	Param       string      `json:"-"`
	RuleType    string      `json:"-"`
}

// Define SQL error
type SQLError struct {
	Number   int
	SQLState []int8
	Message  string
}

// Struct for storing Raw sql result
type Result struct {
	ID        uint
	CreatedAt time.Time
	UpdatedAt time.Time
}

// Create a validator object
var validate = validator.New()

// var db = database.Conn()

// Package Initialization
func init() {
	// Get Database Instance
	// db = database.Conn()
	// Register Custom Validation rules
	validate.RegisterValidation("unique", IsUniqueRule)
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
				RuleType:    err.Type().Name(),
				Message:     GetErrMsg(err.Tag(), err.Field(), err.Param(), err.Type().Name()),
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
func IsUniqueRule(data validator.FieldLevel) bool {
	var result *Result
	targetValue := data.Field().String()
	params := strings.Split(data.Param(), "@")
	sqlQuery := fmt.Sprintf("SELECT id, created_at, updated_at FROM %s WHERE %s='%s'", params[0], params[1], targetValue)
	database.Conn().DB.Raw(sqlQuery).Scan(&result)
	return result == nil
}
