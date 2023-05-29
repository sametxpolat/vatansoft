package repository

import (
	"github.com/sametxpolat/vatansoft/dto"
	"github.com/sametxpolat/vatansoft/model"
	"gorm.io/gorm"
)

type ProductRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{
		db: db,
	}
}

func (r *ProductRepository) Stocks() ([]model.Product, error) {
	var products []model.Product

	err := r.db.Preload("Category").Find(&products).Error

	return products, err
}

func (r *ProductRepository) Stock(id uint) (model.Product, error) {
	var product model.Product

	err := r.db.Where("id = ?", id).First(&product).Error

	return product, err
}

func (r *ProductRepository) Create(product *dto.CProduct) error {
	var prod model.Product

	prod.Name = product.Name
	prod.Barcode = product.Barcode
	prod.Price = product.Price
	prod.CategoryID = product.CategoryID

	err := r.db.Save(&prod).Error

	return err
}

func (r *ProductRepository) Update(id uint, product *dto.UProduct) error {
	var prod model.Product

	err := r.db.Where("id = ?", id).First(&prod).Error
	if err != nil {
		return err
	}

	prod.Name = product.Name
	prod.Price = product.Price

	err = r.db.Save(&prod).Error

	return err
}

func (r *ProductRepository) Delete(id uint) error {
	var product model.Product

	err := r.db.Where("id = ?", id).Delete(&product).Error

	return err
}

func (r *ProductRepository) Filter(name string, barcode string, price uint, categoryID uint) ([]model.Product, error) {
	var products []model.Product

	query := r.db.Model(&model.Product{})

	if name != "" {
		query = query.Where("name LIKE ?", "%"+name+"%")
	}
	if barcode != "" {
		query = query.Where("barcode LIKE ?", "%"+barcode+"%")
	}
	if price != 0 {
		query = query.Where("price = ?", price)
	}
	if categoryID != 0 {
		query = query.Where("category_id = ?", categoryID)
	}

	err := query.Preload("Category").Find(&products).Error

	return products, err
}
