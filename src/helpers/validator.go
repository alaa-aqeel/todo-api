package helpers

import (
	"net/http"

	"github.com/alaa-aqeel/todo/src/lang"
	"github.com/go-playground/validator/v10"
)

func Validator(errorsMap ErrorValidation, body interface{}) {
	validate := validator.New()
	if err := validate.Struct(body); err != nil {
		for _, vErr := range err.(validator.ValidationErrors) {

			errorsMap.Add(vErr.Field(), FieldErrorMessage(vErr))
		}
	}
}

// Map error messages for validation tags
func FieldErrorMessage(err validator.FieldError) string {

	return lang.GetErrors(err.Tag(), err)
}

func MakeValidate[T any](r *http.Request) (*T, ErrorValidation) {
	errorsMap := NewErrorValidation()
	body, err := ParseBody[T](r.Body)
	if err != nil {
		errorsMap.Add("invalid_body", "invalid_body")
		return nil, errorsMap
	}
	Validator(errorsMap, body)

	return body, errorsMap
}
