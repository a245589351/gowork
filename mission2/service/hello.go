package service

import (
	"context"
	"gowork/mission2/proto/hello"
	"time"
)

type Hello struct{}

func (s *Hello) SayHello(ctx context.Context, request *hello.SayHelloRequest) (*hello.SayHelloResponse, error) {
	time.Sleep(time.Second)
	return &hello.SayHelloResponse{
		Message: "hello " + request.Message,
	}, nil
}
