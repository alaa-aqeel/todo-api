package core

import (
	"time"

	"github.com/fasthttp/router"
)

func getTimeOut(timeOut int) time.Duration {
	return time.Second * time.Duration(timeOut)
}

func newRouter() *router.Router {
	router := router.New()
	InitRoutes(router)
	return router
}

func NewApp(config *Config) *FastApp {
	r := newRouter()
	return &FastApp{
		Config:  config,
		Router:  r,
		Handler: r.Handler,
	}
}
