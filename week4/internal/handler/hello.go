package handler

import (
	"context"
	"fmt"
	"gowork/week4/internal/domain/service"
	"gowork/week4/proto/hello"

	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(NewHelloService)

type IHelloService interface {
	SayHello(ctx context.Context, request *hello.SayHelloRequest) (*hello.SayHelloResponse, error)
}

type HelloService struct {
	hello service.IHelloUsecase
}

func NewHelloService(hello service.IHelloUsecase) IHelloService {
	return &HelloService{
		hello: hello,
	}
}

func (s *HelloService) SayHello(ctx context.Context, request *hello.SayHelloRequest) (*hello.SayHelloResponse, error) {
	fmt.Println("test...grpc服务端接收到请求")
	return &hello.SayHelloResponse{
		Message: "hello " + request.Message,
	}, nil
}
