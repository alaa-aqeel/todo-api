package user_handler

import (
	"net/http"

	"github.com/alaa-aqeel/todo/src/api/http/form_requests"
	"github.com/alaa-aqeel/todo/src/helpers"
)

func (h UserHandler) Store(w http.ResponseWriter, r *http.Request) {
	validated, err := form_requests.UserValidate(r)
	if err.HasErrors() {
		helpers.Response(w).ValidationErrors(err)

		return
	}

	helpers.Response(w).
		Status(http.StatusCreated).
		Success("User created successfully", validated)
}
