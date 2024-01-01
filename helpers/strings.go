package helpers

import (
	"regexp"
	_const "starter-golang-new/const"
	"strings"
)

func ToSnakeCase(str string) string {
	var matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
	var matchAllCap = regexp.MustCompile("([a-z0-9])([A-Z])")
	snake := matchFirstCap.ReplaceAllString(str, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
	return strings.ToLower(snake)
}

/***
*  trim text
 */
func ClearText(text string) string {
	return strings.Trim(text, " ")
}

func GetMethodName(path string) string {
	return strings.Split(path, ".")[1]
}

func ClearToken(token string) string {
	if strings.Contains(token, "Bearer") {
		return strings.Replace(token, "Bearer ", "", -1)
	}
	return token
}

func GetClearToken() string {
	return ClearToken(_const.Services.GIN.GetHeader("Authorization"))
}