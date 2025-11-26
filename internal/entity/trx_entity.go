package entity

import "time"

type Trx struct {
	ID          uint   `gorm:"column:id;primaryKey"`
	UserId      uint   `gorm:"column:user_id"`
	AlamatId    uint   `gorm:"column:alamat_id"`
	TokoId      uint   `gorm:"column:toko_id"`
	HargaTotal  int64  `gorm:"column:harga_total"`
	KodeInvoice string `gorm:"column:code_invoice"`
	MethodBayar string `gorm:"column:method_bayar"`
	CreatedAt   time.Time
	UpdatedAt   time.Time

	User      *User       `gorm:"foreignKey:UserId"`
	DetailTrx []DetailTrx `gorm:"foreignKey:TrxId"`
	Alamat    *Address    `gorm:"foreignKey:AlamatId"`
	Toko      *Store      `gorm:"foreignKey:TokoId"`
}

func (u *Trx) TableName() string {
	return "trxes"
}
