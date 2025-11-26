package model

import "time"

type FotoProductResponse struct {
	ID        uint      `json:"id"`
	ProdukID  uint      `json:"produk_id"`
	Url       string    `json:"url"`
	CreatedAt time.Time `json:"created_at"`
}

type ProductResponse struct {
	ID            uint                  `json:"id"`
	CategoryID    uint                  `json:"category_id"`
	TokoID        uint                  `json:"toko_id"`
	NamaProduk    string                `json:"nama_produk"`
	Slug          string                `json:"slug"`
	HargaKonsumen string                `json:"harga_konsumen"`
	Stok          int                   `json:"stok"`
	Deskripsi     string                `json:"deskripsi"`
	CreatedAt     time.Time             `json:"created_at"`
	UpdatedAt     time.Time             `json:"updated_at"`
	Toko          StoreResponse         `json:"toko"`
	Category      CategoryResponse      `json:"category"`
	Fotos         []FotoProductResponse `json:"fotos,omitempty"`
}

type CreateProductRequest struct {
	CategoryID    uint   `json:"category_id" validate:"required"`
	NamaProduk    string `json:"nama_produk" validate:"required"`
	Slug          string `json:"slug" validate:"required"`
	HargaReseller string `json:"harga_reseller" validate:"required"`
	HargaKonsumen string `json:"harga_konsumen" validate:"required"`
	Stok          int    `json:"stok" validate:"required,gte=0"`
	Deskripsi     string `json:"deskripsi" validate:"required"`
}

type UpdateProductRequest struct {
	CategoryID    uint   `json:"category_id" validate:"omitempty,required"`
	NamaProduk    string `json:"nama_produk" validate:"omitempty,required"`
	Slug          string `json:"slug" validate:"omitempty,required"`
	HargaReseller string `json:"harga_reseller" validate:"omitempty,required"`
	HargaKonsumen string `json:"harga_konsumen" validate:"omitempty,required"`
	Stok          int    `json:"stok" validate:"omitempty,gte=0"`
	Deskripsi     string `json:"deskripsi" validate:"omitempty,required"`
}
