package server

import (
	"gowork/week4/global"
	"gowork/week4/initialize"
	"gowork/week4/internal/handler"

	ghttp "gowork/week4/server/http"
)

func NewHTTPServer(helloService handler.IHelloService) *ghttp.Server {
	httpServer := ghttp.NewServer(ghttp.Address(":8000"), ghttp.Router(initialize.InitRouters()))
	global.Servers = append(global.Servers, httpServer)

	return httpServer
}
