package entity

import (
	"time"
)

type Store struct {
	ID        uint   `gorm:"column:id;primaryKey"`
	UserId    uint   `gorm:"column:user_id;unique"`
	NamaToko  string `gorm:"column:nama_toko"`
	UrlFoto   string `gorm:"column:url_foto"`
	CreatedAt time.Time
	UpdatedAt time.Time

	// Relasi
	User      *User       `gorm:"foreignKey:UserId"`
	Produk    []Product   `gorm:"foreignKey:TokoID"`
	DetailTrx []DetailTrx `gorm:"foreignKey:TokoId"`
	Trx       []Trx       `gorm:"foreignKey:TokoId"`
}

func (u *Store) TableName() string {
	return "stores"
}
