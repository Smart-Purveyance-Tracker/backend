package service

import (
	"github.com/Smart-Purveyance-Tracker/backend/entity"
	"github.com/Smart-Purveyance-Tracker/backend/repository"
)

type Product interface {
	ByID(id string) (entity.Product, error)
	Create(product entity.Product) (entity.Product, error)
	List(args repository.ProductListArgs) ([]entity.Product, error)
}

type ProductImpl struct {
	productRepo repository.Product
}

func NewProductImpl(product repository.Product) *ProductImpl {
	return &ProductImpl{
		productRepo: product,
	}
}

func (p *ProductImpl) ByID(id string) (entity.Product, error) {
	return p.productRepo.Find(id)
}

func (p *ProductImpl) Create(product entity.Product) (entity.Product, error) {
	return p.productRepo.Insert(product)
}

func (p *ProductImpl) List(args repository.ProductListArgs) ([]entity.Product, error) {
	return p.productRepo.List(args)
}
