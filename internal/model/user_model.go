package model

import "time"

type UserResponse struct {
	ID           uint       `json:"id"`
	Nama         string     `json:"nama"`
	Email        string     `json:"email"`
	Notelp       string     `json:"notelp"`
	TanggalLahir *time.Time `json:"tanggal_lahir,omitempty"`
	JenisKelamin string     `json:"jenis_kelamin,omitempty"`
	Tentang      string     `json:"tentang,omitempty"`
	Pekerjaan    string     `json:"pekerjaan,omitempty"`
	IdProvinsi   string     `json:"id_provinsi,omitempty"`
	IdKota       string     `json:"id_kota,omitempty"`
	IsAdmin      bool       `json:"is_admin"`
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at"`
}

type RegisterUserRequest struct {
	Nama     string `json:"nama" validate:"required,max=100"`
	Email    string `json:"email" validate:"required,email,max=100"`
	Password string `json:"password" validate:"required,min=8,max=100"`
	Notelp   string `json:"notelp" validate:"required,max=20"`
}

type LoginUserRequest struct {
	Email    string `json:"email" validate:"required,email,max=100"`
	Password string `json:"password" validate:"required,max=100"`
}

type LoginUserResponse struct {
	User  UserResponse `json:"user"`
	Token string       `json:"token"`
}

type UpdateUserRequest struct {
	Nama         *string    `json:"nama" validate:"omitempty,max=100"`
	Email        *string    `json:"email" validate:"omitempty,max=20"`
	Notelp       *string    `json:"notelp" validate:"omitempty,max=20"`
	TanggalLahir *time.Time `json:"tanggal_lahir" validate:"omitempty,datetime=2006-01-02"`
	JenisKelamin *string    `json:"jenis_kelamin" validate:"omitempty,oneof=pria wanita"`
	Tentang      *string    `json:"tentang"`
	Pekerjaan    *string    `json:"pekerjaan"`
	IdProvinsi   *string    `json:"id_provinsi"`
	IdKota       *string    `json:"id_kota"`
}
