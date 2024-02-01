package auth

import (
	"github.com/alaa-aqeel/todo-api/helpers"
	"github.com/valyala/fasthttp"
)

func LoginHandler(ctx *fasthttp.RequestCtx) {

	helpers.Response(ctx).Text("Hi").Send()
}
