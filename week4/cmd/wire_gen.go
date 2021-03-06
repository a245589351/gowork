// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//+build !wireinject

package main

import (
	"github.com/jinzhu/gorm"
	"gowork/week4/internal/domain/repository"
	"gowork/week4/internal/domain/service"
	"gowork/week4/internal/handler"
	"gowork/week4/internal/server"
)

// Injectors from wire.go:

func initApp(db *gorm.DB) *App {
	iHelloRepo := repository.NewHelloRepo(db)
	iHelloUsecase := service.NewHelloUsecase(iHelloRepo)
	iHelloService := handler.NewHelloService(iHelloUsecase)
	httpServer := server.NewHTTPServer(iHelloService)
	grpcServer := server.NewGRPCServer(iHelloService)
	app := newApp(httpServer, grpcServer)
	return app
}
