package user_handler

import (
	"net/http"

	"github.com/alaa-aqeel/todo/src/api/services"
	"github.com/alaa-aqeel/todo/src/api/validators"
	"github.com/alaa-aqeel/todo/src/helpers"
)

func (h UserHandler) Store(w http.ResponseWriter, r *http.Request) {
	validated, err := validators.UserValidate(r)
	if err.HasErrors() {
		helpers.Response(w).ValidationErrors(err)

		return
	}
	user, er := services.UserService().Create(validated)
	if er != nil {
		helpers.Response(w).Error("invalid error", er.Error())

		return
	}
	helpers.Response(w).
		Status(http.StatusCreated).
		Success("User created successfully", user)
}
