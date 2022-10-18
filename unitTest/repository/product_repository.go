package repository

import "unitTest/entity"

type ProductRepository interface {
	FindById(id string) *entity.Product
}
