package user_handler

import (
	"net/http"

	"github.com/alaa-aqeel/todo/src/api/services"
	"github.com/alaa-aqeel/todo/src/helpers"
	"github.com/go-chi/chi/v5"
)

type UserHandler struct{}

func New() UserHandler {
	return UserHandler{}
}

func (h UserHandler) Delete(w http.ResponseWriter, r *http.Request) {
	isDeleted := services.UserService().Delete(chi.URLParam(r, "id"))
	if isDeleted {
		helpers.Response(w).Json(200, "User deleted successfully")

		return
	}
	helpers.Response(w).Json(404, "User not found")
}
