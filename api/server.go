package main

import (
	"fmt"
	"log"
	"net"
	"time"

	"github.com/valyala/fasthttp"
)

func getTimeOut(timeOut int) time.Duration {
	return time.Second * time.Duration(timeOut)
}

func fastHTTPServer(config *Config, hander fasthttp.RequestHandler) *fasthttp.Server {
	fastServer := &fasthttp.Server{
		Handler:     hander,
		Name:        config.Name,
		ReadTimeout: getTimeOut(config.TimeOut),
	}
	listen, err := net.Listen("tcp", fmt.Sprintf("%s:%s", config.HOST, config.PORT))
	if err != nil && config.Debug {
		log.Fatalf("Error[net-listen]: %v", err)
	}
	fmt.Printf("Server: http://%s:%s", config.HOST, config.PORT)
	if err := fastServer.Serve(listen); err != nil && config.Debug {
		log.Fatalf("Error[fasthttp-server]: %v", err)
	}
	if err := fastServer.Shutdown(); err != nil && config.Debug {
		log.Fatalf("Error[shutting-down-server]: while shutting down the server: %v", err)
	}

	return fastServer
}
