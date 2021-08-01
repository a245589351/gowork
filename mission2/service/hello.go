package service

import (
	"context"
	"fmt"
	"time"

	"gowork/mission2/proto/hello"
)

type Hello struct{}

func (s *Hello) SayHello(ctx context.Context, request *hello.SayHelloRequest) (*hello.SayHelloResponse, error) {
	time.Sleep(10 * time.Second)
	fmt.Println("test...grpc服务端接收到请求")
	return &hello.SayHelloResponse{
		Message: "hello " + request.Message,
	}, nil
}
