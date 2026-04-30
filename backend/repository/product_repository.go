package repository

import (
	"glow-fx/domain"

	"gorm.io/gorm"
)

type ProductRepository interface {
	FindAll(string, string, int, int) ([]domain.Product, int64, error)
	FindById(uint) (domain.Product, error)
	FindBySKU(string) (domain.Product, error)
	Create(domain.Product) (domain.Product, error)
	Update(domain.Product) error
	Delete(uint) error
}

type productRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &productRepository{db: db}
}

func (r *productRepository) FindAll(search string, status string, limit int, offset int) ([]domain.Product, int64, error) {
	var products []domain.Product
	var total int64

	query := r.db.Model(&domain.Product{})

	if search != "" {
		query = query.Where("name ILIKE ? OR sku ILIKE ?", "%"+search+"%", "%"+search+"%")
	}

	if status != "" {
		query = query.Where("status = ?", status)
	}

	query.Count(&total)
	err := query.Limit(limit).Offset(offset).Find(&products).Error

	return products, total, err
}

func (r *productRepository) FindById(id uint) (domain.Product, error) {
	var product domain.Product

	err := r.db.Where("id = ?", id).First(&product).Error

	return product, err
}

func (r *productRepository) FindBySKU(sku string) (domain.Product, error) {
	var product domain.Product
	err := r.db.Where("sku = ?", sku).First(&product).Error
	return product, err
}

func (r *productRepository) Create(product domain.Product) (domain.Product, error) {
	err := r.db.Create(&product).Error
	return product, err

}

func (r *productRepository) Update(product domain.Product) error {
	return r.db.Model(&domain.Product{}).Where("id = ?", product.ID).Updates(product).Error
}

func (r *productRepository) Delete(id uint) error {
	return r.db.Delete(&domain.Product{}, id).Error
}
