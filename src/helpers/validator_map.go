package helpers

import (
	"net/http"

	"github.com/go-playground/validator/v10"
)

type ValidateMap struct {
	r        *http.Request
	errors   ErrorValidation
	body     map[string]interface{}
	rules    map[string]interface{}
	validate *validator.Validate
}

func NewValidatorMap(r *http.Request) *ValidateMap {

	return &ValidateMap{
		r:        r,
		errors:   NewErrorValidation(),
		validate: validator.New(),
	}
}

func (vm *ValidateMap) parseBody() bool {
	vm.body = Request(vm.r).Body()

	return vm.body != nil
}

func (v *ValidateMap) Rules(rules map[string]interface{}) *ValidateMap {
	v.rules = rules

	return v
}

func (v *ValidateMap) makeValidateMap() map[string]interface{} {
	return v.validate.ValidateMap(v.body, v.rules)
}

func (v *ValidateMap) parseErrorMessages(errs map[string]interface{}) {
	for f, err := range errs {
		for _, vErr := range err.(validator.ValidationErrors) {
			v.errors.Add(f, FieldErrorMessage(f, vErr))
		}
	}
}

func (v *ValidateMap) Errors() ErrorValidation {

	return v.errors
}

func (v *ValidateMap) Valid() bool {
	if !v.parseBody() {
		v.errors.Add("invalid_body", "invalid_body")

		return false
	}
	errs := v.makeValidateMap()
	v.parseErrorMessages(errs)

	return !v.errors.HasErrors()
}

func (v *ValidateMap) Validated() map[string]interface{} {
	if v.errors.HasErrors() {
		return nil
	}
	validated := make(map[string]interface{})
	for key := range v.rules {
		if value, exists := v.body[key]; exists {
			validated[key] = value
		}
	}

	return validated
}
