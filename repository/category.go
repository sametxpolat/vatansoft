package repository

import (
	"github.com/sametxpolat/vatansoft/dto"
	"github.com/sametxpolat/vatansoft/model"
	"gorm.io/gorm"
)

type CategoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) *CategoryRepository {
	return &CategoryRepository{
		db: db,
	}
}

func (r *CategoryRepository) Categories() ([]model.Category, error) {
	var categories []model.Category

	err := r.db.Find(&categories).Error

	return categories, err
}

func (r *CategoryRepository) Category(id uint) (model.Category, error) {
	var category model.Category

	err := r.db.Where("id = ?", id).First(&category).Error

	return category, err
}

func (r *CategoryRepository) Create(category *dto.CCategory) error {
	var cat model.Category

	cat.Name = category.Name

	err := r.db.Save(&cat).Error

	return err
}

func (r *CategoryRepository) Update(id uint, category *dto.UCategory) error {
	var cat model.Category

	err := r.db.Where("id = ?", id).First(&cat).Error
	if err != nil {
		return err
	}

	cat.Name = category.Name

	err = r.db.Save(&cat).Error

	return err
}

func (r *CategoryRepository) Delete(id uint) error {
	var category model.Category

	err := r.db.Where("id = ?", id).Delete(&category).Error

	return err
}
