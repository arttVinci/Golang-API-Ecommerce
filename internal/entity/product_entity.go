package entity

import "time"

type Product struct {
	ID            uint   `gorm:"column:id;primaryKey"`
	CategoryId    uint   `gorm:"column:category_id"`
	TokoID        uint   `gorm:"column:toko_id"`
	NamaProduk    string `gorm:"column:nama_produk"`
	Slug          string `gorm:"column:slug"`
	HargaReseller string `gorm:"column:harga_reseller"`
	HargaKonsumen string `gorm:"column:harga_konsumen"`
	Stok          int    `gorm:"column:stok"`
	Deskripsi     string `gorm:"column:deskripsi"`
	CreatedAt     time.Time
	UpdatedAt     time.Time

	// Relasi
	FotoProduk []FotoProduct `gorm:"foreignKey:ProdukID"`
	LogProduk  []LogProduct  `gorm:"foreignKey:ProdukId"`
}

func (u *Product) TableName() string {
	return "products"
}
