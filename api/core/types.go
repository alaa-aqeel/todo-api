package core

import (
	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
)

type Map map[string]interface{}

type FastApp struct {
	Server  *fasthttp.Server
	Handler fasthttp.RequestHandler
	Router  *router.Router
	Config  *Config
}

type Route func(g *router.Group)
