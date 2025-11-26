package repository

import (
	"API-Ecommerce-Evermos/internal/entity"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type IAddressRepository interface {
	Save(address entity.Address) (entity.Address, error)
	FindAllByUserId(userId uint) ([]entity.Address, error)
	FindById(id uint) (entity.Address, error)
	Update(address entity.Address) (entity.Address, error)
	Delete(address entity.Address) error
}

type AddressRepository struct {
	db  *gorm.DB
	log *logrus.Logger
}

func NewAddressRepository(db *gorm.DB, log *logrus.Logger) *AddressRepository {
	return &AddressRepository{db, log}
}

func (r *AddressRepository) Save(address entity.Address) (entity.Address, error) {
	err := r.db.Create(&address).Error
	if err != nil {
		r.log.Errorf("Failed to save address: %v", err)
		return address, err
	}
	r.log.Tracef("Address saved with id: %d", address.ID)
	return address, nil
}

func (r *AddressRepository) FindAllByUserId(userId uint) ([]entity.Address, error) {
	var addresses []entity.Address
	err := r.db.Where("id_user = ?", userId).Find(&addresses).Error
	if err != nil {
		r.log.Errorf("Failed to find addresses for user %d: %v", userId, err)
		return nil, err
	}
	return addresses, nil
}

func (r *AddressRepository) FindById(id uint) (entity.Address, error) {
	var address entity.Address
	err := r.db.Where("id = ?", id).First(&address).Error
	if err != nil {
		return address, err
	}
	return address, nil
}

func (r *AddressRepository) Update(address entity.Address) (entity.Address, error) {
	err := r.db.Save(&address).Error
	if err != nil {
		r.log.Errorf("Failed to update address %d: %v", address.ID, err)
		return address, err
	}
	r.log.Tracef("Address updated with id: %d", address.ID)
	return address, nil
}

func (r *AddressRepository) Delete(address entity.Address) error {
	err := r.db.Delete(&address).Error
	if err != nil {
		r.log.Errorf("Failed to delete address %d: %v", address.ID, err)
		return err
	}
	r.log.Tracef("Address deleted with id: %d", address.ID)
	return nil
}
