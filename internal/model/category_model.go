package model

import "time"

type CategoryResponse struct {
	ID           uint      `json:"id"`
	NamaCategory string    `json:"nama_category"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type CreateCategoryRequest struct {
	NamaCategory string `json:"nama_category" validate:"required"`
}

type UpdateCategoryRequest struct {
	NamaCategory *string `json:"nama_category" validate:"required"`
}
