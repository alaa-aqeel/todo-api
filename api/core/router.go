package core

import (
	"github.com/alaa-aqeel/todo-api/routes"
	"github.com/fasthttp/router"
)

func InitRoutes(r *router.Router) {
	apiV1 := r.Group("/api/v1")
	routes.ApiRoutes(apiV1)
}
