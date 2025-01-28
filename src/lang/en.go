package lang

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

var ValidateField = MessagesMap{
	"username": "username",
}

var Messages = MessagesMap{
	"invalid_data": "Invalid data",
	"not_found":    "Not found",
}

var ValidateErrorsMap = ValidateErrors{
	"required": func(fe validator.FieldError) string {
		return fmt.Sprintf("The %s is required", GetField(fe.Field()))
	},
	"min": func(fe validator.FieldError) string {
		return fmt.Sprintf("The %s must be at least %s characters", GetField(fe.Field()), fe.Param())
	},
	"max": func(fe validator.FieldError) string {
		return fmt.Sprintf("The %s must be less than %s characters", GetField(fe.Field()), fe.Param())
	},
}
