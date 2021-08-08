package main

import (
	"gowork/week4/core"
	"gowork/week4/server/grpc"
	ghttp "gowork/week4/server/http"

	"github.com/jinzhu/gorm"
)

type App struct {
	hs *ghttp.Server
	gs *grpc.Server
}

func main() {
	db := &gorm.DB{}
	initApp(db)

	if err := core.Run(); err != nil {
		panic(err)
	}
}

func newApp(hs *ghttp.Server, gs *grpc.Server) *App {
	return &App{
		hs: hs,
		gs: gs,
	}
}
