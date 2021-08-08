package service

import (
	"gowork/week4/internal/domain/model"
	"gowork/week4/internal/domain/repository"

	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(NewHelloUsecase)

type IHelloUsecase interface {
	List() (result []*model.Hello)
}

type HelloUsecase struct {
	repo repository.IHelloRepo
}

func NewHelloUsecase(repo repository.IHelloRepo) IHelloUsecase {
	return &HelloUsecase{
		repo: repo,
	}
}

func (uc *HelloUsecase) List() (result []*model.Hello) {
	result = uc.repo.ListHello()
	return
}
