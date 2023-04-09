package service

import (
	"errors"
	"jwt_token/models"
	"jwt_token/repository"
)

type ProductService struct {
	Repository repository.ProductRepository

}

type ProductServiceAll struct {
	Repository repository.ProductRepositoryAll
}

func (s ProductService) GetOneProduct(id string) (*models.Product, error) {
	product := s.Repository.FindByID(id)

	if product == nil {
		return nil, errors.New("product not found")
	}

	return product, nil
}

// func (s ProductService) GetAllProduct(models.Product) (*models.Product, error) {
// 	product1 := s.Repository.FindByALL()

// 	if product1 == nil {
// 		return nil, errors.New("product not found")
// 	}

// 	return product1, nil
// }