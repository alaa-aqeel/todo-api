package lang

import (
	"strings"

	"github.com/go-playground/validator/v10"
)

func FormatMessage(msg string, params map[string]string) string {
	for key, val := range params {
		placeholder := ":" + key
		msg = strings.ReplaceAll(msg, placeholder, val)
	}
	return msg
}

func GetErrors(rule, fielName string, field validator.FieldError) string {
	if _, ok := ValidateErrorsMap[rule]; !ok {
		return rule
	}

	return FormatMessage(ValidateErrorsMap[rule], map[string]string{
		"attr":  fielName,
		"param": field.Param(),
	})
}

func GetField(key string) string {
	key = strings.ToLower(key)
	value, ok := ValidateField[key]
	if !ok {
		return key
	}

	return value
}

func Message(key string) string {
	value, ok := Messages[key]
	if !ok {
		return key
	}

	return value
}
