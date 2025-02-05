package user_handler

import (
	"net/http"

	"github.com/alaa-aqeel/todo/src/api/services"
	"github.com/alaa-aqeel/todo/src/api/validators"
	"github.com/alaa-aqeel/todo/src/helpers"
	"github.com/go-chi/chi/v5"
)

func (h UserHandler) Update(w http.ResponseWriter, r *http.Request) {
	validator := validators.NewUserValidate(r)
	if !validator.UpdateValidate() {
		helpers.Response(w).Status(422).Error("invalid error", validator.Errors())

		return
	}
	user, err := services.UserService().Update(chi.URLParam(r, "id"), validator.Validated())
	if err != nil {
		helpers.Response(w).
			Status(400).
			Error("invalid error", err)
		return
	}
	helpers.Response(w).
		Status(200).
		Success("User created successfully", user[0])
}
