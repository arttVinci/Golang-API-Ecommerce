package repository

import (
	"API-Ecommerce-Evermos/internal/entity"
	"API-Ecommerce-Evermos/internal/model"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type IProductRepository interface {
	Save(product entity.Product) (entity.Product, error)
	SaveImage(foto entity.FotoProduct) (entity.FotoProduct, error)
	FindById(id uint) (entity.Product, error)
	UpdateStok(tx *gorm.DB, id uint, quantity int) error
	Search(filter model.ProductFilter) ([]entity.Product, int64, error)
}

type ProductRepository struct {
	db  *gorm.DB
	log *logrus.Logger
}

func NewProductRepository(db *gorm.DB, log *logrus.Logger) *ProductRepository {
	return &ProductRepository{db, log}
}

func (r *ProductRepository) Save(product entity.Product) (entity.Product, error) {
	err := r.db.Create(&product).Error
	if err != nil {
		r.log.Errorf("Failed to save product: %v", err)
		return product, err
	}
	return product, nil
}

func (r *ProductRepository) SaveImage(foto entity.FotoProduct) (entity.FotoProduct, error) {
	err := r.db.Create(&foto).Error
	if err != nil {
		r.log.Errorf("Failed to save product image: %v", err)
		return foto, err
	}
	return foto, nil
}

func (r *ProductRepository) FindById(id uint) (entity.Product, error) {
	var product entity.Product
	err := r.db.Where("id = ?", id).First(&product).Error
	return product, err
}

func (r *ProductRepository) UpdateStok(tx *gorm.DB, id uint, quantity int) error {
	return tx.Model(&entity.Product{}).Where("id = ?", id).
		Update("stok", gorm.Expr("stok - ?", quantity)).Error
}

func (r *ProductRepository) Search(filter model.ProductFilter) ([]entity.Product, int64, error) {
	var products []entity.Product
	var total int64
	
	query := r.db.Model(&entity.Product{}).Preload("FotoProduk")
	if filter.Search != "" {
		query = query.Where("nama_produk LIKE ?", "%"+filter.Search+"%")
	}

	if filter.CategoryID > 0 {
		query = query.Where("category_id = ?", filter.CategoryID)
	}

	err := query.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	offset := (filter.Page - 1) * filter.Limit
	err = query.Offset(offset).Limit(filter.Limit).Find(&products).Error

	return products, total, err
}
