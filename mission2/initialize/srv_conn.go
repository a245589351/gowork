package initialize

import (
	"gowork/mission2/global"
	"gowork/mission2/proto/hello"

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
