package helpers

import "strings"

func StringsNewReplacer(textString string) string {
	textString = strings.NewReplacer(
		"+", "%20",
		"%", "%25",
		"/", "%2F",
		" ", "%20",
	).Replace(textString)

	return textString
}
