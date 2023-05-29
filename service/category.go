package service

import (
	"github.com/sametxpolat/vatansoft/dto"
	"github.com/sametxpolat/vatansoft/model"
	"github.com/sametxpolat/vatansoft/repository"
)

type CategoryService struct {
	categoryRepository *repository.CategoryRepository
}

func NewCategoryService(categoryRepository *repository.CategoryRepository) *CategoryService {
	return &CategoryService{
		categoryRepository: categoryRepository,
	}
}

func (s *CategoryService) Categories() ([]model.Category, error) {
	return s.categoryRepository.Categories()
}

func (s *CategoryService) Category(id uint) (model.Category, error) {
	return s.categoryRepository.Category(id)
}

func (s *CategoryService) Create(category *dto.CCategory) error {
	return s.categoryRepository.Create(category)
}

func (s *CategoryService) Update(id uint, category *dto.UCategory) error {
	return s.categoryRepository.Update(id, category)
}

func (s *CategoryService) Delete(id uint) error {
	return s.categoryRepository.Delete(id)
}
