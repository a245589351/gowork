package server

import (
	"gowork/week4/global"
	"gowork/week4/internal/handler"
	"gowork/week4/proto/hello"
	"gowork/week4/server/grpc"
)

func NewGRPCServer(helloService handler.IHelloService) *grpc.Server {
	grpcServer := grpc.NewServer(grpc.Address(":9000"))
	hello.RegisterHelloServer(grpcServer.Server, helloService)
	global.Servers = append(global.Servers, grpcServer)
	return grpcServer
}
