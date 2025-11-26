package entity

import "time"

type LogProduct struct {
	ID            uint   `gorm:"column:id;primaryKey"`
	ProdukId      uint   `gorm:"column:produk_id"`
	CategoryId    uint   `gorm:"column:category_id"`
	TokoID        uint   `gorm:"column:toko_id"`
	NamaProduk    string `gorm:"column:nama_produk"`
	Slug          string `gorm:"column:slug"`
	HargaReseller string `gorm:"column:harga_reseller"`
	HargaKonsumen string `gorm:"column:harga_konsumen"`
	Deskripsi     string `gorm:"column:deskripsi"`
	CreatedAt     time.Time
	UpdatedAt     time.Time

	DetailTrx *DetailTrx `gorm:"foreignKey:LogProductId"`
}

func (u *LogProduct) TableName() string {
	return "log_products"
}
