package repository

import (
	"jwt_token/models"

	"github.com/stretchr/testify/mock"
)
type ProductRepositoryMock struct {
	Mock mock.Mock
}

type ProductRepositoryALLMock struct {
	Mock mock.Mock
}

func (repository *ProductRepositoryMock) FindByID(id string) *models.Product {
	arguments := repository.Mock.Called(id)

	if arguments.Get(0) == nil {
		return nil
	}

	product := arguments.Get(0).(models.Product)
	return &product
}

func (repository *ProductRepositoryALLMock) FindByALL(models.Product) *models.Product {
	arguments := repository.Mock.Called()
	if arguments.Get(0) == nil {
		return nil
	}

	product := arguments.Get(0).(models.Product)
	return &product
}