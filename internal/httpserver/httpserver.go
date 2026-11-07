package httpserver

import (
	"context"
	"fmt"
	"net/http"
	"time"
	"net"
	"errors"
	"high-performance-server/internal/config"
)




type HttpServer struct{

	Server *http.Server

	
}


func NewHttpServer(httpCfg config.HttpConfig) (*HttpServer,error){
	rootmux := NewMux()


	rootmux.HandleFunc("GET /about",func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello from about endpoint\n"))
	})


	return &HttpServer{

		Server : &http.Server{
			Addr: ":"+fmt.Sprintf("%d",httpCfg.Port),
			ReadTimeout: time.Duration(httpCfg.ReadTimeout*int32(time.Second)),
			ReadHeaderTimeout: time.Duration(httpCfg.ReadHeaderTimeout*int32(time.Second)),
			WriteTimeout: time.Duration(httpCfg.WriteTimeout*int32(time.Second)) ,
			IdleTimeout: time.Duration(httpCfg.IdleTimeout*int32(time.Second)),
			MaxHeaderBytes: 1<<20,
			Handler:rootmux,

		},

	},nil
}



func (h *HttpServer) Run(ctx context.Context)error{
	h.Server.BaseContext = func(net.Listener) context.Context {
		return ctx
	}


	err := h.Server.ListenAndServe()
	if errors.Is(err, http.ErrServerClosed) {
		return nil
	}
	if err != nil {
		return fmt.Errorf("http server failed: %w", err)
	}
	return nil
}

func (h *HttpServer) Stop(ctx context.Context)error{



	return h.Server.Shutdown(ctx)
}