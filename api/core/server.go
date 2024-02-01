package core

import (
	"fmt"
	"log"
	"net"

	"github.com/valyala/fasthttp"
)

func (f *FastApp) NewServer() *fasthttp.Server {
	return &fasthttp.Server{
		ReadTimeout: getTimeOut(f.Config.TimeOut),
		Handler:     f.Handler,
		Name:        f.Config.Name,
	}
}

func netListen(config *Config) net.Listener {
	listen, err := net.Listen("tcp", fmt.Sprintf("%s:%s", config.HOST, config.PORT))
	if err != nil && config.Debug {
		log.Fatalf("Error[net-listen]: %v", err)
	}
	return listen
}

func (f *FastApp) FasthttpRunServer() {
	server := f.NewServer()
	listen := netListen(f.Config)

	fmt.Printf("Server: http://%s:%s", f.Config.HOST, f.Config.PORT)
	if err := server.Serve(listen); err != nil && f.Config.Debug {
		log.Fatalf("Error[fasthttp-server]: %v", err)
	}
	if err := server.Shutdown(); err != nil && f.Config.Debug {
		log.Fatalf("Error[shutting-down-server]: while shutting down the server: %v", err)
	}
}
