package entity

import "time"

type Address struct {
	ID           uint   `gorm:"column:id;primaryKey"`
	UserId       uint   `gorm:"column:id_user"`
	JudulAlamat  string `gorm:"column:judul_alamat"`
	NamaPenerima string `gorm:"column:nama_penerima"`
	NoTelp       string `gorm:"column:no_telp"`
	DetailAlamat string `gorm:"column:detail_alamat"`
	CreatedAt    time.Time
	UpdatedAt    time.Time

	Trx  []Trx `gorm:"foreignKey:Alamat"`
	User *User `gorm:"foreignKey:UserId"`
}

func (u *Address) TableName() string {
	return "address"
}
