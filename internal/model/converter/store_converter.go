package converter

import (
	"API-Ecommerce-Evermos/internal/entity"
	"API-Ecommerce-Evermos/internal/model"
)

func StoreToResponse(store entity.Store) model.StoreResponse {
	return model.StoreResponse{
		ID:        store.ID,
		UserID:    store.UserId,
		NamaToko:  store.NamaToko,
		UrlFoto:   store.UrlFoto,
		CreatedAt: store.CreatedAt,
		UpdatedAt: store.UpdatedAt,
	}
}
