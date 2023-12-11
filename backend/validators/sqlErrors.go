package validators

import (
	"fmt"
	"gofiber-api/helpers"
)

var ErrorNumbers = []int{
	1062,
}

var errMessages = map[int]string{
	1062: "SQL Error number: %v Duplicate entry",
}

func GetSqlErrMsg(errCode int, message string) string {
	if helpers.IsIntInSlice(ErrorNumbers, errCode) {
		return fmt.Sprintf(errMessages[errCode], errCode)
	}
	return message
}
