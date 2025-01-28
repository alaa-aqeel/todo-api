package lang

import "github.com/go-playground/validator/v10"

type ValidateErrors map[string]func(validator.FieldError) string

type MessagesMap map[string]string
