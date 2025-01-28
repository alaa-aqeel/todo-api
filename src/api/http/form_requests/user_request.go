package form_requests

import (
	"net/http"

	"github.com/alaa-aqeel/todo/src/helpers"
)

type UserValidated struct {
	Name     string `json:"name" validate:"required,min=3,max=50"`
	Username string `json:"username" validate:"required,min=3,max=20"`
	Password string `json:"password" validate:"required,min=8"`
}

func UserValidate(r *http.Request) (*UserValidated, helpers.ErrorValidation) {

	return helpers.MakeValidate[UserValidated](r)
}
