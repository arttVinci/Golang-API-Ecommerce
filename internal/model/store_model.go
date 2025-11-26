package model

import "time"

type StoreResponse struct {
	ID        uint      `json:"id"`
	UserID    uint      `json:"user_id"`
	NamaToko  string    `json:"nama_toko"`
	UrlFoto   string    `json:"url_foto"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UpdateStoreRequest struct {
	NamaToko string `json:"nama_toko" validate:"required"`
}
