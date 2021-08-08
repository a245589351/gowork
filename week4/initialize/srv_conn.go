package initialize

import (
	"gowork/week4/global"
	"gowork/week4/proto/hello"

	"google.golang.org/grpc"
)

func InitSrvConn() error {
	// 连接服务端
	helloConn, err := grpc.Dial(
		"127.0.0.1:9000",
		grpc.WithInsecure(),
	)

	if err != nil {
		return err
	}

	global.HelloSrvClient = hello.NewHelloClient(helloConn)

	return nil
}
