// +build wireinject

package main

import (
	"gowork/week4/internal/domain/repository"
	"gowork/week4/internal/domain/service"
	"gowork/week4/internal/handler"
	"gowork/week4/internal/server"

	"github.com/jinzhu/gorm"

	"github.com/google/wire"
)

func initApp(*gorm.DB) *App {
	panic(wire.Build(server.ProviderSet, repository.ProviderSet, service.ProviderSet, handler.ProviderSet, newApp))
}
