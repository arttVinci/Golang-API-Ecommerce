package converter

import (
	"API-Ecommerce-Evermos/internal/entity"
	"API-Ecommerce-Evermos/internal/model"
)

func UserToResponse(user entity.User) model.UserResponse {
	return model.UserResponse{
		ID:           user.ID,
		Nama:         user.Nama,
		Email:        user.Email,
		Notelp:       user.Notelp,
		TanggalLahir: user.TanggalLahir,
		Tentang:      user.Tentang,
		Pekerjaan:    user.Pekerjaan,
		IdProvinsi:   user.IdProvinsi,
		IdKota:       user.IdKota,
		IsAdmin:      user.IsAdmin,
	}
}
