package lang

import (
	"strings"

	"github.com/go-playground/validator/v10"
)

func GetErrors(key string, field validator.FieldError) string {
	if _, ok := ValidateErrorsMap[key]; !ok {
		return key
	}

	return ValidateErrorsMap[key](field)
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
