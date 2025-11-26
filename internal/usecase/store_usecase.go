package usecase

import (
	"API-Ecommerce-Evermos/internal/model"
	"API-Ecommerce-Evermos/internal/model/converter"
	"API-Ecommerce-Evermos/internal/repository"
	"errors"

	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
)

type StoreUsecase struct {
	log       *logrus.Logger
	validate  *validator.Validate
	storeRepo repository.IStoreRepository
}

func NewStoreUsecase(log *logrus.Logger, validate *validator.Validate, storeRepo repository.IStoreRepository) *StoreUsecase {
	return &StoreUsecase{log, validate, storeRepo}
}

func (u *StoreUsecase) GetMyStore(userId uint) (model.StoreResponse, error) {
	store, err := u.storeRepo.FindByUserId(userId)
	if err != nil {
		return model.StoreResponse{}, errors.New("toko tidak ditemukan")
	}

	storeResponse := converter.StoreToResponse(store)

	return storeResponse, nil
}

func (u *StoreUsecase) Update(userId uint, request model.UpdateStoreRequest) (model.StoreResponse, error) {
	if err := u.validate.Struct(request); err != nil {
		return model.StoreResponse{}, err
	}

	store, err := u.storeRepo.FindByUserId(userId)
	if err != nil {
		return model.StoreResponse{}, errors.New("toko tidak ditemukan")
	}

	// 3. Update Data
	if request.NamaToko != "" {
		store.NamaToko = request.NamaToko
	}

	updatedStore, err := u.storeRepo.Update(store)
	if err != nil {
		return model.StoreResponse{}, err
	}

	storeResponse := converter.StoreToResponse(updatedStore)

	return storeResponse, nil
}
