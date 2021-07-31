package http

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
)

type ServerOption func(s *Server)

type Server struct {
	*http.Server
	ctx     context.Context
	lis     net.Listener
	once    sync.Once
	network string
	addr    string
	err     error
	router  *gin.Engine
}

func NewServer(opts ...ServerOption) *Server {
	srv := &Server{}
	for _, o := range opts {
		o(srv)
	}

	srv.Server = &http.Server{
		Addr:    srv.addr,
		Handler: srv.router,
	}

	return srv
}

func Address(addr string) ServerOption {
	return func(s *Server) {
		s.addr = addr
	}
}

func Router(r *gin.Engine) ServerOption {
	return func(s *Server) {
		s.router = r
	}
}

func (s *Server) Start(ctx context.Context) error {
	fmt.Println("HTTP Server Start......")
	return s.ListenAndServe()
}

func (s *Server) Stop(ctx context.Context) error {
	fmt.Println("HTTP Server Shutdown......")
	return s.Shutdown(ctx)
}
