package validators

import (
	"net/http"

	"github.com/alaa-aqeel/todo/src/helpers"
)

type UserValidate struct {
	*BaseValidator
}

func NewUserValidate(r *http.Request) *UserValidate {
	return &UserValidate{
		&BaseValidator{
			validator: helpers.NewValidatorMap(r),
		},
	}
}

func (u UserValidate) CreateValidate() bool {
	return u.Validate(map[string]interface{}{
		"name":     "required",
		"username": "required,min=3,max=20",
		"password": "required,min=6",
	})
}

func (u UserValidate) UpdateValidate() bool {
	return u.Validate(map[string]interface{}{
		"name":     "min=2",
		"username": "max=8",
	})
}
