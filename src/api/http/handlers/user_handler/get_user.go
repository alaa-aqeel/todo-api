package user_handler

import (
	"net/http"

	"github.com/alaa-aqeel/todo/src/api/services"
	"github.com/alaa-aqeel/todo/src/helpers"
	"github.com/go-chi/chi/v5"
)

func (h UserHandler) Index(w http.ResponseWriter, r *http.Request) {

	users := services.
		UserService().
		Paginate(
			helpers.NewPagination(r),
			helpers.RequestQueryArags(r),
		)
	helpers.Response(w).Json(200, users)
}

func (h UserHandler) Show(w http.ResponseWriter, r *http.Request) {

	user := services.
		UserService().
		Get(chi.URLParam(r, "id"))

	helpers.Response(w).Json(200, user)
}
