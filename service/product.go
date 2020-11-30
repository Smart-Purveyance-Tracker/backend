package service

import (
	"time"

	"github.com/Smart-Purveyance-Tracker/backend/entity"
	"github.com/Smart-Purveyance-Tracker/backend/repository"
)

type Product interface {
	ByID(id string) (entity.Product, error)
	Create(product entity.Product) (entity.Product, error)
	List(args repository.ProductListArgs) ([]entity.Product, error)
	Update(product entity.Product) (entity.Product, error)
	ScanProducts(args ScanProductsArgs) ([]ProductCount, error)
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

func (p *ProductImpl) Update(product entity.Product) (entity.Product, error) {
	return p.productRepo.Update(product)
}

func (p *ProductImpl) List(args repository.ProductListArgs) ([]entity.Product, error) {
	return p.productRepo.List(args)
}

type ScanProductsArgs struct {
	BoughtAt time.Time
	Type     string
}

type ProductCount struct {
	Product entity.Product
	Count   uint64
}

func (p *ProductImpl) ScanProducts(args ScanProductsArgs) ([]ProductCount, error) {
	return []ProductCount{
		{
			Count: 1,
			Product: entity.Product{
				ID:       "1",
				Name:     "ОВОЩ",
				BoughtAt: args.BoughtAt,
			},
		},
	}, nil
}
