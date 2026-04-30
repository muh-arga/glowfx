package usecase

import (
	"errors"
	"glow-fx/domain"
	"glow-fx/repository"

	"gorm.io/gorm"
)

type ProductUseCase struct {
	repo repository.ProductRepository
}

func NewProductUseCase(r repository.ProductRepository) *ProductUseCase {
	return &ProductUseCase{repo: r}
}

func (u *ProductUseCase) GetProducts(search string, status string, page int, limit int) ([]domain.Product, int64, error) {
	if page < 1 {
		page = 1
	}

	if limit <= 0 {
		limit = 10
	}

	offset := (page - 1) * limit
	return u.repo.FindAll(search, status, limit, offset)
}

func (u *ProductUseCase) GetProductById(id uint) (domain.Product, error) {
	product, err := u.repo.FindById(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return product, errors.New("product not found")
		}

		return product, err
	}

	return product, nil
}

func (u *ProductUseCase) CreateProduct(input domain.Product) (domain.Product, error) {
	errs := ValidationErrors{}

	if input.Name == "" {
		errs["name"] = "name is required"
	}

	if input.SKU == "" {
		errs["sku"] = "sku is rerquired"
	}

	if input.Status == "" {
		input.Status = "active"
	}

	if len(errs) > 0 {
		return input, errs
	}

	existing, err := u.repo.FindBySKU(input.SKU)
	if err == nil && existing.ID != 0 {
		return input, ValidationErrors{
			"sku": "sku already exists",
		}
	}

	if len(errs) > 0 {
		return input, errs
	}

	return u.repo.Create(input)
}

func (u *ProductUseCase) UpdateProduct(id uint, input domain.Product) (domain.Product, error) {
	product, err := u.repo.FindById(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return product, errors.New("product not found")
		}

		return product, err
	}

	errs := ValidationErrors{}

	if input.Name == "" {
		errs["name"] = "name is required"
	}

	if input.SKU == "" {
		errs["sku"] = "sku is required"
	}

	if input.Status != "" {
		product.Status = input.Status
	}

	if len(errs) > 0 {
		return product, errs
	}

	if input.SKU != product.SKU {
		existing, err := u.repo.FindBySKU(input.SKU)

		if err == nil && existing.ID != 0 {
			return product, ValidationErrors{
				"sku": "sku already exists",
			}
		}
	}

	product.Name = input.Name
	product.SKU = input.SKU

	err = u.repo.Update(product)
	return product, err
}

func (u *ProductUseCase) DeleteProduct(id uint) error {
	_, err := u.repo.FindById(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("product not found")
		}

		return err
	}

	return u.repo.Delete(id)
}
