package entity

import "time"

type Category struct {
	ID           uint   `gorm:"column:id;primaryKey"`
	NamaCategory string `gorm:"column:nama_category"`
	CreatedAt    time.Time
	UpdatedAt    time.Time

	Produk    []Product    `gorm:"foreignKey:CategoryId"`
	LogProduk []LogProduct `gorm:"foreignKey:CategoryId"`
}

func (u *Category) TableName() string {
	return "category"
}
