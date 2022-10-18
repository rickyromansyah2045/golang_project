package service

import (
	"errors"
	"unitTest/entity"
	"unitTest/repository"
)

type ProductService struct {
	Repository repository.ProductRepository
}

func (service ProductService) GetOneProduct(id string) (*entity.Product, error) {
	product := service.Repository.FindById(id)

	if product == nil {
		return nil, errors.New("Product not Found")
	}

	return product, nil
}
