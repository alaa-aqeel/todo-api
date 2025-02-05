package validators

import "github.com/alaa-aqeel/todo/src/helpers"

type BaseValidator struct {
	validator *helpers.ValidateMap
}

func (u *BaseValidator) Validate(rules map[string]interface{}) bool {

	return u.validator.Rules(rules).Valid()
}

func (u *BaseValidator) Validated() map[string]interface{} {

	return u.validator.Validated()
}

func (u *BaseValidator) Errors() helpers.ErrorValidation {

	return u.validator.Errors()
}
