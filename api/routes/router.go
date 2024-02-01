package routes

import (
	"github.com/alaa-aqeel/todo-api/handlers/auth"
	"github.com/fasthttp/router"
)

func ApiRoutes(g *router.Group) {

	g.GET("/login", auth.LoginHandler)
}
