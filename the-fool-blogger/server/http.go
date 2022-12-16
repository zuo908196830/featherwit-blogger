package server

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"sync"
)

type HttpServer struct {
	port int32
	server *http.Server
}

var (
	httpServer *HttpServer
	once       sync.Once
)

func NewHttpServer(Router *gin.Engine) *HttpServer {
	if httpServer != nil {
		return httpServer
	}
	port := "8080"
	once.Do(func() {
		httpServer = &HttpServer{
			port: 8081,
			server: &http.Server{
				Addr:           fmt.Sprintf(":%d", port),
				Handler:        Router,
				MaxHeaderBytes: 1 << 20,
			},
		}
	})
	return httpServer
}

func (hs *HttpServer) Start()  {
	go func() {
		if err := hs.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server start error: %s\n", err)
		}
	}()
}

func (h *HttpServer) ShutDown(ctx context.Context) {
	if err := h.server.Shutdown(ctx); err != nil {
		log.Fatal("Merak Http Driver Shutdown:", err)
	}
}
