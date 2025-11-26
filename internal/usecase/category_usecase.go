package usecase

import (
	"API-Ecommerce-Evermos/internal/entity"
	"API-Ecommerce-Evermos/internal/model"
	"API-Ecommerce-Evermos/internal/model/converter"
	"API-Ecommerce-Evermos/internal/repository"
	"errors"

	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
)

type CategoryUsecase struct {
	log          *logrus.Logger
	validate     *validator.Validate
	categoryRepo repository.ICategoryRepository
}

func NewCategoryUsecase(log *logrus.Logger, validate *validator.Validate, categoryRepo repository.ICategoryRepository) *CategoryUsecase {
	return &CategoryUsecase{log, validate, categoryRepo}
}

func (u *CategoryUsecase) Create(request model.CreateCategoryRequest) (model.CategoryResponse, error) {
	if err := u.validate.Struct(request); err != nil {
		return model.CategoryResponse{}, err
	}

	category := entity.Category{NamaCategory: request.NamaCategory}
	newCategory, err := u.categoryRepo.Save(category)
	if err != nil {
		u.log.Errorf("Failed to save category: %v", err)
		return model.CategoryResponse{}, err
	}

	return converter.CategoryToResponse(&newCategory), nil
}

func (u *CategoryUsecase) List() ([]model.CategoryResponse, error) {
	categories, err := u.categoryRepo.FindAll()
	if err != nil {
		u.log.Errorf("Failed to fetch categories: %v", err)
		return []model.CategoryResponse{}, err
	}

	var responses []model.CategoryResponse
	for _, cat := range categories {
		responses = append(responses, converter.CategoryToResponse(&cat))
	}
	return responses, nil
}

func (u *CategoryUsecase) Update(CategoryId uint, request model.UpdateCategoryRequest) (model.CategoryResponse, error) {
	category, err := u.categoryRepo.FindById(CategoryId)
	if err != nil {
		u.log.Warnf("Category not found id %d: %v", CategoryId, err)
		return model.CategoryResponse{}, errors.New("category tidak ditemukan")
	}

	if request.NamaCategory != nil {
		category.NamaCategory = *request.NamaCategory
	}

	categoryUpdated, err := u.categoryRepo.Update(category)
	if err != nil {
		u.log.Errorf("Failed to update category: %v", err)
		return model.CategoryResponse{}, err
	}

	categoryResponse := converter.CategoryToResponse(&categoryUpdated)

	return categoryResponse, nil
}

func (u *CategoryUsecase) Delete(CategoryId uint) error {
	category, err := u.categoryRepo.FindById(CategoryId)
	if err != nil {
		u.log.Warnf("Delete failed, category not found id %d", CategoryId)
		return errors.New("category tidak ditemukan")
	}

	err = u.categoryRepo.Delete(category)
	if err != nil {
		u.log.Errorf("Failed to delete category: %v", err)
		return err
	}

	return nil

}
