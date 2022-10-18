package service

import (
	"testing"
	"unitTest/entity"
	"unitTest/repository"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var productRepository = &repository.ProductRepositoryMock{Mock: mock.Mock{}}
var productService = ProductService{Repository: productRepository}

func TestProductServiceGetOneProductNotFound(t *testing.T) {
	product := entity.Product{
		Id:   "2",
		Name: "Kaca Mata",
	}

	productRepository.Mock.On("FindById", "2").Return(product)

	result, err := productService.GetOneProduct("2")

	assert.Nil(t, err)

	assert.NotNil(t, result)

	assert.Equal(t, product.Id, result.Id, "result has to be '2'")
	assert.Equal(t, product.Name, result.Name, "result has to be 'Name'")
	assert.Equal(t, &product, result, "result has to be Product with id '2'")
}

// func TestProductServiceGetOneProductNotFound(t *testing.T) {
// 	productRepository.Mock.On("FindById", "1").Return(nil)

// 	product, err := productService.GetOneProduct("1")

// 	assert.Nil(t, product)
// 	assert.NotNil(t, err)
// 	assert.Equal(t, "product not found", err.Error(), "error responce has to be 'roduct not found'")
// }
