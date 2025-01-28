package validators

import (
	"net/http"

	"github.com/alaa-aqeel/todo/src/helpers"
)

type UserValidator struct {
	Name     string `json:"name" validate:"required,min=3,max=50"`
	Username string `json:"username" validate:"required,min=3,max=20"`
	Password string `json:"password" validate:"required,min=8"`
}

func UserValidate(r *http.Request) (*UserValidator, helpers.ErrorValidation) {

	return helpers.MakeValidate[UserValidator](r)
}
