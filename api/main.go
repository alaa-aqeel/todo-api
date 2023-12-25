package main

import (
	"github.com/fasthttp/router"
)

func main() {
	config := &Config{}
	config.LoadConfig()
	fastRouter := router.New()
	apiRoutes(fastRouter.Group("/api"))
	fastHTTPServer(config, fastRouter.Handler)
}
