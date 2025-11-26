package entity

import "time"

type FotoProduct struct {
	ID        uint   `gorm:"column:id;primaryKey"`
	ProdukID  uint   `gorm:"column:produk_id"`
	Url       string `gorm:"column:url"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (u *FotoProduct) TableName() string {
	return "foto_products"
}
