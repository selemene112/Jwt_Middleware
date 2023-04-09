package service

import (
	"jwt_token/models"
	"jwt_token/repository"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)


var productRepo = &repository.ProductRepositoryMock{Mock: mock.Mock{}}
var productService = ProductService{Repository: productRepo}

func TestProductServiceGetOneProductNotFound(t *testing.T) {
	productRepo.Mock.On("FindByID","1").Return(nil)

	product, err := productService.GetOneProduct("1") 

	assert.Nil(t, product)
	assert.NotNil(t, err)
	assert.Equal(t, "product not found", err.Error(), "error response has to be 'product not found'")
}

func TestProductServiceGetOneProductGet(t *testing.T) {

	product := models.Product{
		UserID: 1,
		Title: "mau di peluk",
		
	}
	productRepo.Mock.On("FindByID","2").Return(product)

	result, err := productService.GetOneProduct("2") 

	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, product.UserID, result.UserID)
	assert.Equal(t, product.Title, result.Title)
}

// func TestProductServiceGetOneProductGETALL(t *testing.T) {
// 	productRepo.Mock.On("FindByID",models.Product{}).Return(nil)

// 	product, err := productService.GetOneProduct() 

// 	assert.Nil(t, product)
// 	assert.NotNil(t, err)
// 	assert.Equal(t, "product not found", err.Error(), "error response has to be 'product not found'")
// }