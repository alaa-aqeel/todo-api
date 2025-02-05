package user_handler

import (
	"net/http"

	"github.com/alaa-aqeel/todo/src/api/services"
	"github.com/alaa-aqeel/todo/src/api/validators"
	"github.com/alaa-aqeel/todo/src/helpers"
)

func (h UserHandler) Store(w http.ResponseWriter, r *http.Request) {
	validator := validators.NewUserValidate(r)
	if !validator.CreateValidate() {
		helpers.Response(w).Status(422).Error("invalid error", validator.Errors())

		return
	}
	user, er := services.UserService().Create(validator.Validated())
	if er != nil {
		helpers.Response(w).Error("invalid error", er.Error())

		return
	}
	helpers.Response(w).
		Status(http.StatusCreated).
		Success("User updated successfully", user)
}
