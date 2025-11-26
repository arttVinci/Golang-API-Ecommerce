package entity

import "time"

type DetailTrx struct {
	ID           uint `gorm:"column:id;primaryKey"`
	TrxId        uint `gorm:"column:trx_id"`
	TokoId       uint `gorm:"column:toko_id;"`
	LogProductId uint `gorm:"column:log_product_id"`
	Kuantitas    int  `gorm:"column:quantity"`
	HargaTotal   int  `gorm:"column:harga_total"`
	CreatedAt    time.Time
	UpdatedAt    time.Time

	Trx        *Trx        `gorm:"foreignKey:TrxId"`
	LogProduct *LogProduct `gorm:"foreignKey:LogProductId"`
	Toko       *Store      `gorm:"foreignKey:TokoId"`
}

func (u *DetailTrx) TableName() string {
	return "detail_trxes"
}
