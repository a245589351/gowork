package global

import (
	"gowork/week4/proto/hello"
	"gowork/week4/server"
)

var (
	HelloSrvClient hello.HelloClient
	Servers        []server.Server
)
