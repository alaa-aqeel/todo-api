package test

import (
	"github.com/alaa-aqeel/todo-api/core"
	"github.com/alaa-aqeel/todo-api/helpers"
	"github.com/valyala/fasthttp"
)

func CreateAppForTest(config *core.Config) *core.FastApp {
	config.BaseDir = helpers.GetBaseDir("api")
	config.LoadConfig(".env")
	return core.NewApp(config)
}

func CreateCtx(method string, path string, body []byte) *fasthttp.RequestCtx {
	ctx := &fasthttp.RequestCtx{
		Request:  *fasthttp.AcquireRequest(),
		Response: *fasthttp.AcquireResponse(),
	}
	ctx.Request.Header.SetMethod(method)
	ctx.Request.SetRequestURI(path)
	ctx.Request.SetBody(body)
	return ctx
}

// For testing
func RequestCtx(method string, path string, body []byte) *fasthttp.RequestCtx {
	app := CreateAppForTest(&core.Config{})

	ctx := CreateCtx(method, path, body)
	app.Handler(ctx)
	return ctx
}
