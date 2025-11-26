package converter

import (
	"API-Ecommerce-Evermos/internal/entity"
	"API-Ecommerce-Evermos/internal/model"
)

func AddressToResponse(address entity.Address) model.AddressResponse {
	return model.AddressResponse{
		ID:           address.ID,
		JudulAlamat:  address.JudulAlamat,
		NamaPenerima: address.NamaPenerima,
		UserID:       address.UserId,
		NoTelp:       address.NoTelp,
		DetailAlamat: address.DetailAlamat,
	}
}
