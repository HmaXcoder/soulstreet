package service

import (
	"fmt"
	"soulstreet/internal/model"
	"soulstreet/internal/repository"
)


type ProductService struct {
	repo repository.ProductRepository
}

func NewProductService(repo repository.ProductRepository) * ProductService {
	return &ProductService{repo: repo}
}

func (s *ProductService) CreateProduct(product *model.Product) error {
	if err := s.repo.Create(product); err != nil {
		return fmt.Errorf("Erro ao criar produto: %v", err)
	}
	return nil
}

func (s *ProductService) GetProductByID(id int) (*model.Product, error) {
	return s.repo.GetByID(id)
}

func (s *ProductService) GetAll() ([]model.Product, error) {
	return s.repo.GetAll()
}