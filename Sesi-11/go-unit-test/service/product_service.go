package service

import (
	"errors"
	"go-unit-test/entity"
	"go-unit-test/repository"
)

type ProductService struct {
	Repository repository.ProductRepository
}

func (service ProductService) GetOneProduct(id string) (*entity.Produk, error) {
	product := service.Repository.FindById(id)
	if product == nil {
		return nil, errors.New("Product not found")
	}

	return product, nil
}