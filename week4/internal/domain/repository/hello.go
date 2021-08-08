package repository

import (
	"gowork/week4/internal/domain/model"

	"github.com/google/wire"

	"github.com/jinzhu/gorm"
)

var ProviderSet = wire.NewSet(NewHelloRepo)

type IHelloRepo interface {
	ListHello() []*model.Hello
}

type helloRepo struct {
	db *gorm.DB
}

func NewHelloRepo(db *gorm.DB) IHelloRepo {
	return &helloRepo{
		db: db,
	}
}

func (h *helloRepo) ListHello() []*model.Hello {
	rv := make([]*model.Hello, 0)
	return rv
}
