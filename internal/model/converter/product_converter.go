package converter

import (
	"API-Ecommerce-Evermos/internal/entity"
	"API-Ecommerce-Evermos/internal/model"
)

func ProductToResponse(product entity.Product, store model.StoreResponse, category model.CategoryResponse) model.ProductResponse {
	return model.ProductResponse{
		ID:            product.ID,
		CategoryID:    product.CategoryId,
		TokoID:        product.TokoID,
		NamaProduk:    product.NamaProduk,
		Slug:          product.Slug,
		HargaKonsumen: product.HargaKonsumen,
		Stok:          product.Stok,
		Deskripsi:     product.Deskripsi,
		CreatedAt:     product.CreatedAt,
		Toko:          store,
		Category:      category,
	}
}
