package grpc

import (
	"context"
	"fmt"
	"net"
	"sync"

	"google.golang.org/grpc"
)

type ServerOption func(s *Server)

type Server struct {
	*grpc.Server
	ctx     context.Context
	lis     net.Listener
	once    sync.Once
	network string
	addr    string
	err     error
}

func Address(addr string) ServerOption {
	return func(s *Server) {
		s.addr = addr
	}
}

func (s *Server) Listen() error {
	s.once.Do(func() {
		lis, err := net.Listen(s.network, s.addr)
		if err != nil {
			s.err = err
			return
		}
		s.lis = lis
	})
	if s.err != nil {
		return s.err
	}

	return nil
}

func (s *Server) Start(ctx context.Context) error {
	err := s.Listen()
	if err != nil {
		return err
	}

	fmt.Println("GRPC Server Start......")
	return s.Serve(s.lis)
}

func (s *Server) Stop(ctx context.Context) error {
	s.GracefulStop()
	fmt.Println("GRPC Server Shutdown......")
	return nil
}

func NewServer(opts ...ServerOption) *Server {
	srv := &Server{
		network: "tcp",
	}
	srv.Server = grpc.NewServer()

	for _, s := range opts {
		s(srv)
	}
	return srv
}
