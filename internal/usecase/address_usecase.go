package usecase

import (
	"API-Ecommerce-Evermos/internal/entity"
	"API-Ecommerce-Evermos/internal/model"
	"API-Ecommerce-Evermos/internal/model/converter"
	"API-Ecommerce-Evermos/internal/repository"
	"errors"

	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
)

type AddressUsecase struct {
	log         *logrus.Logger
	validate    *validator.Validate
	addressRepo repository.IAddressRepository
}

func NewAddressUsecase(log *logrus.Logger, validate *validator.Validate, addressRepo repository.IAddressRepository) *AddressUsecase {
	return &AddressUsecase{log, validate, addressRepo}
}

func (u *AddressUsecase) Create(userId uint, request model.CreateAddressRequest) (model.AddressResponse, error) {
	if err := u.validate.Struct(request); err != nil {
		return model.AddressResponse{}, err
	}

	address := entity.Address{
		UserId:       userId,
		JudulAlamat:  request.JudulAlamat,
		NamaPenerima: request.NamaPenerima,
		NoTelp:       request.NoTelp,
		DetailAlamat: request.DetailAlamat,
	}

	newAddress, err := u.addressRepo.Save(address)
	if err != nil {
		return model.AddressResponse{}, err
	}

	return converter.AddressToResponse(newAddress), nil
}

func (u *AddressUsecase) List(userId uint) ([]model.AddressResponse, error) {
	addresses, err := u.addressRepo.FindAllByUserId(userId)
	if err != nil {
		return []model.AddressResponse{}, err
	}

	var responses []model.AddressResponse
	for _, addr := range addresses {
		responses = append(responses, converter.AddressToResponse(addr))
	}
	return responses, nil
}

func (u *AddressUsecase) Delete(userId uint, addressId uint) error {
	address, err := u.addressRepo.FindById(addressId)
	if err != nil {
		return errors.New("alamat tidak ditemukan")
	}

	if address.UserId != userId {
		return errors.New("forbidden: anda tidak berhak menghapus alamat ini")
	}

	return u.addressRepo.Delete(address)
}
