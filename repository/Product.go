package repository

import "jwt_token/models"

type ProductRepository interface {
	FindByID(id string) *models.Product
	
}

type ProductRepositoryAll interface {
	FindByALL(models.Product) *models.Product
}