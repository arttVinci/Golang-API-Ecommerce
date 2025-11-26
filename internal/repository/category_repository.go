package repository

import (
	"API-Ecommerce-Evermos/internal/entity"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type ICategoryRepository interface {
	Save(category entity.Category) (entity.Category, error)
	FindAll() ([]entity.Category, error)
	FindById(id uint) (entity.Category, error)
	Update(category entity.Category) (entity.Category, error)
	Delete(category entity.Category) error
}
type CategoryRepository struct {
	db  *gorm.DB
	log *logrus.Logger
}

func NewCategoryRepository(db *gorm.DB, log *logrus.Logger) *CategoryRepository {
	return &CategoryRepository{db, log}
}

func (r *CategoryRepository) Save(category entity.Category) (entity.Category, error) {
	err := r.db.Create(&category).Error
	return category, err
}

func (r *CategoryRepository) FindAll() ([]entity.Category, error) {
	var categories []entity.Category
	err := r.db.Find(&categories).Error
	return categories, err
}

func (r *CategoryRepository) FindById(id uint) (entity.Category, error) {
	var category entity.Category
	err := r.db.Where("id = ?", id).First(&category).Error
	return category, err
}

func (r *CategoryRepository) Update(category entity.Category) (entity.Category, error) {
	err := r.db.Save(&category).Error
	return category, err
}

func (r *CategoryRepository) Delete(category entity.Category) error {
	return r.db.Delete(&category).Error
}
