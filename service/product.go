package service

import (
	"github.com/sametxpolat/vatansoft/dto"
	"github.com/sametxpolat/vatansoft/model"
	"github.com/sametxpolat/vatansoft/repository"
)

type ProductService struct {
	productRepository *repository.ProductRepository
}

func NewProductService(productRepository *repository.ProductRepository) *ProductService {
	return &ProductService{
		productRepository: productRepository,
	}
}

func (s *ProductService) Stocks() ([]model.Product, error) {
	return s.productRepository.Stocks()
}

func (s *ProductService) Stock(id uint) (model.Product, error) {
	return s.productRepository.Stock(id)
}

func (s *ProductService) Create(product *dto.CProduct) error {
	return s.productRepository.Create(product)
}

func (s *ProductService) Update(id uint, product *dto.UProduct) error {
	return s.productRepository.Update(id, product)
}

func (s *ProductService) Delete(id uint) error {
	return s.productRepository.Delete(id)
}

func (s *ProductService) Filter(name string, barcode string, price uint, categoryID uint) ([]model.Product, error) {
	return s.productRepository.Filter(name, barcode, price, categoryID)
}
