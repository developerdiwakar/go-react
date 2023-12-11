package validators

import (
	"fmt"
	"gofiber-api/helpers"
)

var ruleMessages map[string]interface{}

func GetErrMsg(tag, field, param, ruleType string) string {
	// If params are available
	ruleMsg := getRuleMsg(tag, field, param, ruleType)

	if helpers.IsStringInSlice([]string{
		"ConfirmPassword",
		"MobileNumber",
	}, field) {
		return ruleMsg[field].(map[string]string)[ruleType]
	}
	return ruleMsg[tag].(map[string]string)[ruleType]
}

// Get rule message
func getRuleMsg(tag, field, param, ruleType string) map[string]interface{} {
	ruleMessages = map[string]interface{}{
		"ConfirmPassword": map[string]string{
			"string": fmt.Sprintf("%s must be same as %s.", field, param),
		},
		"MobileNumber": map[string]string{
			"string": fmt.Sprintf("The %s must be %s characters including (+91).", field, param),
		},
		"required": map[string]string{
			"string": fmt.Sprintf("%s must have a value.", field),
		},
		"email": map[string]string{
			"string": fmt.Sprintf("%s must be a valid e-mail address.", field),
		},
		"oneof": map[string]string{
			"string": fmt.Sprintf("%s field is not valid.", field),
		},
		"istplexists": map[string]string{
			"string": fmt.Sprintf("The given %s didn't exist.", field),
		},
		"isfileexist": map[string]string{
			"string": fmt.Sprintf("The given %s file didn't exist.", field),
		},
		"unique": map[string]string{
			"string": fmt.Sprintf("The %s has already been taken.", field),
		},
		"min": map[string]string{
			"string": fmt.Sprintf("The %s must be at least %s characters.", field, param),
		},
		"max": map[string]string{
			"string": fmt.Sprintf("The %s may not be greater than %s characters.", field, param),
		},
		"len": map[string]string{
			"string": fmt.Sprintf("The %s must be %s characters.", field, param),
		},
	}

	// return the messages
	return ruleMessages
}
