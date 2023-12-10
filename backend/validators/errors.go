package validators

import (
	"fmt"
)

var ruleMessages = map[string]string{
	"required":        "%s must have a value",
	"email":           "%s must be a valid e-mail address",
	"oneof":           "%s field is not valid",
	"istplexists":     "The given %s didn't exist",
	"isfileexist":     "The given %s file didn't exist ",
	"ConfirmPassword": "%s must be same as %s",
}

func GetErrMsg(tag, field string, value interface{}, param string) string {
	if param != "" {
		if field == "ConfirmPassword" {
			return fmt.Sprintf(ruleMessages[field], field, param)
		}

		return fmt.Sprintf(ruleMessages[tag], field, param)
	}
	return fmt.Sprintf(ruleMessages[tag], field)
}
