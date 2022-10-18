package repository

import (
	"unitTest/entity"

	"github.com/stretchr/testify/mock"
)

type ProductRepositoryMock struct {
	Mock mock.Mock
}

func (repository *ProductRepositoryMock) FindById(id string) *entity.Product {
	argument := repository.Mock.Called(id)

	if argument.Get(0) == nil {
		return nil
	}

	product := argument.Get(0).(entity.Product)

	return &product
}
