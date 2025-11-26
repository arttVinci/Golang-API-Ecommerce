package entity

import "time"

type User struct {
	ID           uint       `gorm:"column:id;primaryKey"`
	Nama         string     `gorm:"column:name"`
	Notelp       string     `gorm:"column:no_telp;unique"`
	Email        string     `gorm:"column:email;unique"`
	Password     string     `gorm:"column:password"`
	TanggalLahir *time.Time `gorm:"column:tanggal_lahir"`
	JenisKelamin string     `gorm:"column:jenis_kelamin"`
	Tentang      string     `gorm:"column:tentang"`
	Pekerjaan    string     `gorm:"column:pekerjaan"`
	IdProvinsi   string     `gorm:"column:id_provinsi"`
	IdKota       string     `gorm:"column:id_kota"`
	IsAdmin      bool       `gorm:"column:is_admin;default:false"`
	CreatedAt    *time.Time
	UpdatedAt    *time.Time

	// Relasi
	Toko   *Store    `gorm:"foreignKey:UserId"`
	Alamat []Address `gorm:"foreignKey:UserId"`
	Trx    []Trx     `gorm:"foreignKey:UserId"`
}

func (u *User) TableName() string {
	return "users"
}
