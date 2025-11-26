package converter

import (
	"API-Ecommerce-Evermos/internal/entity"
	"API-Ecommerce-Evermos/internal/model"
)

func CategoryToResponse(category *entity.Category) model.CategoryResponse {
	return model.CategoryResponse{
		ID:           category.ID,
		NamaCategory: category.NamaCategory,
	}
}
